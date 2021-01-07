package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/book"
)

// BookController defines the struct for the book controller
type BookController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBook handles POST requests for adding book entities
// @Summary Create book
// @Description Create book
// @ID create-book
// @Accept   json
// @Produce  json
// @Param book body ent.Book true "Book entity"
// @Success 200 {object} ent.Book
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /books [post]
func (ctl *BookController) CreateBook(c *gin.Context) {
	obj := ent.Book{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "book binding failed",
		})
		return
	}

	u, err := ctl.client.Book.
		Create().
		SetBookName(obj.BookName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetBook handles GET requests to retrieve a book entity
// @Summary Get a book entity by ID
// @Description get book by ID
// @ID get-book
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} ent.Book
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /books/{id} [get]
func (ctl *BookController) GetBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Book.
		Query().
		Where(book.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListBook handles request to get a list of book entities
// @Summary List book entities
// @Description list book entities
// @ID list-book
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Book
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /books [get]
func (ctl *BookController) ListBook(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	books, err := ctl.client.Book.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, books)
}

// DeleteBook handles DELETE requests to delete a book entity
// @Summary Delete a book entity by ID
// @Description get book by ID
// @ID delete-book
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /books/{id} [delete]
func (ctl *BookController) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Book.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateBook handles PUT requests to update a book entity
// @Summary Update a book entity by ID
// @Description update book by ID
// @ID update-book
// @Accept   json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body ent.Book true "Book entity"
// @Success 200 {object} ent.Book
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /books/{id} [put]
func (ctl *BookController) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Book{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "book binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Book.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewBookController creates and registers handles for the book controller
func NewBookController(router gin.IRouter, client *ent.Client) *BookController {
	uc := &BookController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitBookController registers routes to the main engine
func (ctl *BookController) register() {
	books := ctl.router.Group("/books")

	books.GET("", ctl.ListBook)

	// CRUD
	books.POST("", ctl.CreateBook)
	books.GET(":id", ctl.GetBook)
	books.PUT(":id", ctl.UpdateBook)
	books.DELETE(":id", ctl.DeleteBook)
}

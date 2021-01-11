package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/book"
	"github.com/team11/app/ent/bookborrow"
	"github.com/team11/app/ent/servicepoint"
	"github.com/team11/app/ent/user"
)

// BookborrowController defines the struct for the bookborrow controller
type BookborrowController struct {
	client *ent.Client
	router gin.IRouter
}

//Bookborrow struct
type Bookborrow struct {
	UserID         int
	BookID         int
	ServicePointID int
}

// CreateBookborrow handles POST requests for adding bookborrow entities
// @Summary Create bookborrow
// @Description Create bookborrow
// @ID create-bookborrow
// @Accept   json
// @Produce  json
// @Param bookborrow body Bookborrow true "Bookborrow entity"
// @Success 200 {object} ent.Bookborrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows [post]
func (ctl *BookborrowController) CreateBookborrow(c *gin.Context) {
	obj := Bookborrow{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookborrow binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.UserID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	bk, err := ctl.client.Book.
		Query().
		Where(book.IDEQ(int(obj.BookID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Book not found",
		})
		return
	}

	sp, err := ctl.client.ServicePoint.
		Query().
		Where(servicepoint.IDEQ(int(obj.ServicePointID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "servicepoint not found",
		})
		return
	}

	times := time.Now().Local()

	now := time.Now()
	then := time.Date(2021, 1, 8, 15, 37, 0, 0, time.UTC)
	after := time.Date(2021, 1, 15, 15, 37, 0, 0, time.UTC)
	diff := after.Sub(then)
	times2 := now.Add(diff)

	bb, err := ctl.client.Bookborrow.
		Create().
		SetBOOK(bk).
		SetUSER(u).
		SetSERVICEPOINT(sp).
		SetBORROWDATE(times).
		SetRETURNDATE(times2).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	book, err := ctl.client.Book.
		UpdateOne(bk).
		SetStatusID(4).
		Save(context.Background())

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Update status Book error",
			})
			return
		}
	fmt.Print(book)

	c.JSON(200, bb)
}

// GetBookborrow handles GET requests to retrieve a bookborrow entity
// @Summary Get a bookborrow entity by ID
// @Description get bookborrow by ID
// @ID get-bookborrow
// @Produce  json
// @Param id path int true "Bookborrow ID"
// @Success 200 {object} ent.Bookborrow
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows/{id} [get]
func (ctl *BookborrowController) GetBookborrow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bb, err := ctl.client.Bookborrow.
		Query().
		Where(bookborrow.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bb)
}

// ListBookborrow handles request to get a list of bookborrow entities
// @Summary List bookborrow entities
// @Description list bookborrow entities
// @ID list-bookborrow
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bookborrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows [get]
func (ctl *BookborrowController) ListBookborrow(c *gin.Context) {
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

	bookborrows, err := ctl.client.Bookborrow.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bookborrows)
}

// DeleteBookborrow handles DELETE requests to delete a bookborrow entity
// @Summary Delete a bookborrow entity by ID
// @Description get bookborrow by ID
// @ID delete-bookborrow
// @Produce  json
// @Param id path int true "Bookborrow ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows/{id} [delete]
func (ctl *BookborrowController) DeleteBookborrow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bookborrow.
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

// UpdateBookborrow handles PUT requests to update a bookborrow entity
// @Summary Update a bookborrow entity by ID
// @Description update bookborrow by ID
// @ID update-bookborrow
// @Accept   json
// @Produce  json
// @Param id path int true "Bookborrow ID"
// @Param bookborrow body ent.Bookborrow true "Bookborrow entity"
// @Success 200 {object} ent.Bookborrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookborrows/{id} [put]
func (ctl *BookborrowController) UpdateBookborrow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Bookborrow{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookborrow binding failed",
		})
		return
	}
	obj.ID = int(id)
	bb, err := ctl.client.Bookborrow.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, bb)
}

// NewBookborrowController creates and registers handles for the bookborrow controller
func NewBookborrowController(router gin.IRouter, client *ent.Client) *BookborrowController {
	bb := &BookborrowController{
		client: client,
		router: router,
	}
	bb.register()
	return bb
}

// InitBookborrowController registers routes to the main engine
func (ctl *BookborrowController) register() {
	bookborrows := ctl.router.Group("/bookborrows")

	bookborrows.GET("", ctl.ListBookborrow)

	// CRUD
	bookborrows.POST("", ctl.CreateBookborrow)
	bookborrows.GET(":id", ctl.GetBookborrow)
	bookborrows.PUT(":id", ctl.UpdateBookborrow)
	bookborrows.DELETE(":id", ctl.DeleteBookborrow)
}

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
	"github.com/team11/app/ent/location"
	"github.com/team11/app/ent/user"
	"github.com/team11/app/ent/bookreturn"
)

// BookreturnController defines the struct for the bookreturn controller
type BookreturnController struct {
	client *ent.Client
	router gin.IRouter
}

// Bookreturn defines the struct for the bookreturn
type Bookreturn struct {
	UserID           int
	LocationID       int
	BookborrowID     int
	DamagedPoint     int
	DamagedPointName string
	Lost             string
	ReturnTime       string
}

// CreateBookreturn handles POST requests for adding bookreturn entities
// @Summary Create bookreturn
// @Description Create bookreturn
// @ID create-bookreturn
// @Accept   json
// @Produce  json
// @Param bookreturn body Bookreturn true "bookreturn entity"
// @Success 200 {object} ent.Bookreturn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookreturns [post]
func (ctl *BookreturnController) CreateBookreturn(c *gin.Context) {
	obj := Bookreturn{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bookreturn binding failed",
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

	l, err := ctl.client.Location.
		Query().
		Where(location.IDEQ(int(obj.LocationID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "location not found",
		})
		return
	}

	b, err := ctl.client.Bookborrow.
		Query().
		Where(bookborrow.IDEQ(int(obj.BookborrowID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bookborrow not found",
		})
		return
	}

	settime := time.Now().Format("2006-01-02T15:04:05Z07:00")
	time, err := time.Parse(time.RFC3339, settime)

	br, err := ctl.client.Bookreturn.
		Create().
		SetUser(u).
		SetLocation(l).
		SetMustreturn(b).
		SetDAMAGEDPOINT(obj.DamagedPoint).
		SetDAMAGEDPOINTNAME(obj.DamagedPointName).
		SetLOST(obj.Lost).
		SetRETURNTIME(time).
		Save(context.Background())

	if err != nil {
		fmt.Print(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   br,
	})

	bb, err := ctl.client.Book.
		Query().
		Where(book.IDEQ(int(obj.BookborrowID))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "book not found",
		})
		return
	}

	bk, err := ctl.client.Book.
		UpdateOne(bb).
		SetStatusID(1).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Update status book error",
		})
		return
	}
	fmt.Print(bk)

	borrows, err := ctl.client.Bookborrow.
		UpdateOne(b).
		SetSTATUSID(5).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Update status bookborow error",
		})
		return
	}
	fmt.Print(borrows)

}

// ListBookreturn handles request to get a list of bookreturn entities
// @Summary List bookreturn entities
// @Description list bookreturn entities
// @ID list-bookreturn
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bookreturn
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookreturns [get]
func (ctl *BookreturnController) ListBookreturn(c *gin.Context) {
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

	bookreturns, err := ctl.client.Bookreturn.
		Query().
		WithUser().
		WithLocation().
		WithMustreturn().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bookreturns)
}

// GetBookreturn handles GET requests to retrieve a bookreturn entity
// @Summary Get a bookreturn entity by ID
// @Description get bookreturn by ID
// @ID get-bookreturn
// @Produce  json
// @Param id path int true "Bookreturn ID"
// @Success 200 {array} ent.Bookreturn
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookreturns/{id} [get]
func (ctl *BookreturnController) GetBookreturn(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	br, err := ctl.client.Bookreturn.
		Query().
		WithUser().
		WithMustreturn().
		WithLocation().
		Where(bookreturn.HasUserWith(user.IDEQ(int(id)))).
		All(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, br)

}

// DeleteBookreturn handles DELETE requests to delete a bookreturn entity
// @Summary Delete a bookreturn entity by ID
// @Description get bookreturn by ID
// @ID delete-bookreturn
// @Produce  json
// @Param id path int true "Bookreturn ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookreturns/{id} [delete]
func (ctl *BookreturnController) DeleteBookreturn(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Bookreturn.
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

// NewBookreturnController creates and registers handles for the bookreturn controller
func NewBookreturnController(router gin.IRouter, client *ent.Client) *BookreturnController {
	bc := &BookreturnController{
		client: client,
		router: router,
	}
	bc.register()
	return bc
}

// InitBookreturnController registers routes to the main engine
func (ctl *BookreturnController) register() {
	bookreturns := ctl.router.Group("/bookreturns")

	bookreturns.GET(":id", ctl.GetBookreturn)
	bookreturns.GET("", ctl.ListBookreturn)
	bookreturns.POST("", ctl.CreateBookreturn)
	bookreturns.DELETE(":id", ctl.DeleteBookreturn)
}

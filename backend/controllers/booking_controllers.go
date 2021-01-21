package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/booking"
	"github.com/team11/app/ent/cliententity"
	"github.com/team11/app/ent/servicepoint"
	"github.com/team11/app/ent/user"
)

// BookingController defines the struct for the booking controller
type BookingController struct {
	client *ent.Client
	router gin.IRouter
}

//Booking struct
type Booking struct {
	User         int
	Client       int
	ServicePoint int
	UserNumber   int
	BorrowItem   int
	PhoneNumber  string
}

// CreateBooking handles POST requests for adding booking entities
// @Summary Create booking
// @Description Create booking
// @ID create-booking
// @Accept   json
// @Produce  json
// @Param booking body Booking true "Booking entity"
// @Success 200 {object} ent.Booking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookings [post]
func (ctl *BookingController) CreateBooking(c *gin.Context) {
	obj := Booking{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "booking binding failed",
		})
		return
	}
	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	sp, err := ctl.client.ServicePoint.
		Query().
		Where(servicepoint.IDEQ(int(obj.ServicePoint))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Bookingtype not found",
		})
		return
	}
	cl, err := ctl.client.ClientEntity.
		Query().
		Where(cliententity.IDEQ(int(obj.Client))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Client not found",
		})
		return
	}
	times := time.Now().Local()
	now := time.Now()
	then := time.Date(2021, 1, 8, 15, 37, 0, 0, time.UTC)
	after := time.Date(2021, 1, 8, 18, 37, 0, 0, time.UTC)
	diff := after.Sub(then)
	times2 := now.Add(diff)
	b, err := ctl.client.Booking.
		Create().
		SetBOOKINGDATE(times).
		SetUsedby(u).
		SetGetservice(sp).
		SetUSERNUMBER(obj.UserNumber).
		SetBORROWITEM(obj.BorrowItem).
		SetPHONENUMBER(obj.PhoneNumber).
		SetUsing(cl).
		SetTIMELEFT(times2).
		Save(context.Background())
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status": false,
			"error":  err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   b,
	})
}

// GetBooking handles GET requests to retrieve a booking entity
// @Summary Get a booking entity by ID
// @Description get booking by ID
// @ID get-booking
// @Produce  json
// @Param id path int true "Booking ID"
// @Success 200 {object} ent.Booking
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookings/{id} [get]
func (ctl *BookingController) GetBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := ctl.client.Booking.
		Query().
		Where(booking.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListBooking handles request to get a list of booking entities
// @Summary List booking entities
// @Description list booking entities
// @ID list-booking
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Booking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookings [get]
func (ctl *BookingController) ListBooking(c *gin.Context) {
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

	bookings, err := ctl.client.Booking.
		Query().
		WithGetservice().
		WithUsedby().
		WithUsing().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, bookings)
}

// DeleteBooking handles DELETE requests to delete a booking entity
// @Summary Delete a booking entity by ID
// @Description get booking by ID
// @ID delete-booking
// @Produce  json
// @Param id path int true "Booking ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookings/{id} [delete]
func (ctl *BookingController) DeleteBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Booking.
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

// UpdateBooking handles PUT requests to update a Booking entity
// @Summary Update a booking entity by ID
// @Description update booking by ID
// @ID update-booking
// @Accept   json
// @Produce  json
// @Param id path int true "Booking ID"
// @Param booking body ent.Booking true "Booking entity"
// @Success 200 {object} ent.Booking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bookings/{id} [put]
func (ctl *BookingController) UpdateBooking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Booking{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "booking binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.Booking.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, b)
}

// NewBookingController creates and registers handles for the booking controller
func NewBookingController(router gin.IRouter, client *ent.Client) *BookingController {
	uc := &BookingController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitBookingController registers routes to the main engine
func (ctl *BookingController) register() {
	bookings := ctl.router.Group("/bookings")

	bookings.GET("", ctl.ListBooking)

	// CRUD
	bookings.POST("", ctl.CreateBooking)
	bookings.GET(":id", ctl.GetBooking)
	bookings.PUT(":id", ctl.UpdateBooking)
	bookings.DELETE(":id", ctl.DeleteBooking)
}

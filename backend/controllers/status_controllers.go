package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team11/app/ent"
	"github.com/team11/app/ent/status"
	"github.com/gin-gonic/gin"
)

// StatusController defines the struct for the status controller
type StatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStatus handles POST requests for adding status entities
// @Summary Create status
// @Description Create status
// @ID create-status
// @Accept   json
// @Produce  json
// @Param status body ent.Status true "Status entity"
// @Success 200 {object} ent.Status
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuss [post]
func (ctl *StatusController) CreateStatus(c *gin.Context) {
	obj := ent.Status{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Status binding failed",
		})
		return
	}

	st, err := ctl.client.Status.
		Create().
		SetSTATUSNAME(obj.STATUSNAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, st)
}

// GetStatus handles GET requests to retrieve a status entity
// @Summary Get a status entity by ID
// @Description get status by ID
// @ID get-status
// @Produce  json
// @Param id path int true "Status ID"
// @Success 200 {object} ent.Status
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuss/{id} [get]
func (ctl *StatusController) GetStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	st, err := ctl.client.Status.
		Query().
		Where(status.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, st)
}

// ListStatus handles request to get a list of status entities
// @Summary List status entities
// @Description list status entities
// @ID list-status
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Status
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuss [get]
func (ctl *StatusController) ListStatus(c *gin.Context) {
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

	status, err := ctl.client.Status.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, status)
}

// DeleteStatus handles DELETE requests to delete a status entity
// @Summary Delete a status entity by ID
// @Description get status by ID
// @ID delete-status
// @Produce  json
// @Param id path int true "Status ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuss/{id} [delete]
func (ctl *StatusController) DeleteStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Status.
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

// UpdateStatus handles PUT requests to update a status entity
// @Summary Update a status entity by ID
// @Description update status by ID
// @ID update-status
// @Accept   json
// @Produce  json
// @Param id path int true "Status ID"
// @Param status body ent.Status true "Status entity"
// @Success 200 {object} ent.Status
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statuss/{id} [put]
func (ctl *StatusController) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Status{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Status binding failed",
		})
		return
	}
	obj.ID = int(id)
	st, err := ctl.client.Status.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, st)
}

// NewStatusController creates and registers handles for the status controller
func NewStatusController(router gin.IRouter, client *ent.Client) *StatusController {
	st := &StatusController{
		client: client,
		router: router,
	}
	st.register()
	return st
}

// InitStatusController registers routes to the main engine
func (ctl *StatusController) register() {
	statuss := ctl.router.Group("/statuss")

	statuss.GET("", ctl.ListStatus)

	// CRUD
	statuss.POST("", ctl.CreateStatus)
	statuss.GET(":id", ctl.GetStatus)
	statuss.PUT(":id", ctl.UpdateStatus)
	statuss.DELETE(":id", ctl.DeleteStatus)
}

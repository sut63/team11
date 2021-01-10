package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/servicepoint"
)

// ServicePointController defines the struct for the servicepoint controller
type ServicePointController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateServicePoint handles POST requests for adding servicepoint entities
// @Summary Create servicepoint
// @Description Create servicepoint
// @ID create-servicepoint
// @Accept   json
// @Produce  json
// @Param ent.servicepoint body ent.ServicePoint true "ServicePoint entity"
// @Success 200 {object} ent.ServicePoint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /servicepoints [post]
func (ctl *ServicePointController) CreateServicePoint(c *gin.Context) {
	obj := ent.ServicePoint{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "servicepoint binding failed",
		})
		return
	}

	sp, err := ctl.client.ServicePoint.
		Create().
		SetCOUNTERNUMBER(obj.COUNTERNUMBER).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sp)
}

// GetServicePoint handles GET requests to retrieve a servicepoint entity
// @Summary Get a servicepoint entity by ID
// @Description get servicepoint by ID
// @ID get-servicepoint
// @Produce  json
// @Param id path int true "ServicePoint ID"
// @Success 200 {object} ent.ServicePoint
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /servicepoints/{id} [get]
func (ctl *ServicePointController) GetServicePoint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	bb, err := ctl.client.ServicePoint.
		Query().
		Where(servicepoint.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bb)
}

// ListServicePoint handles request to get a list of servicepoint entities
// @Summary List servicepoint entities
// @Description list servicepoint entities
// @ID list-servicepoint
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ServicePoint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /servicepoints [get]
func (ctl *ServicePointController) ListServicePoint(c *gin.Context) {
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

	servicepoints, err := ctl.client.ServicePoint.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, servicepoints)
}

// DeleteServicePoint handles DELETE requests to delete a servicepoint entity
// @Summary Delete a servicepoint entity by ID
// @Description get servicepoint by ID
// @ID delete-servicepoint
// @Produce  json
// @Param id path int true "ServicePoint ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /servicepoints/{id} [delete]
func (ctl *ServicePointController) DeleteServicePoint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ServicePoint.
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

// UpdateServicePoint handles PUT requests to update a servicepoint entity
// @Summary Update a servicepoint entity by ID
// @Description update servicepoint by ID
// @ID update-servicepoint
// @Accept   json
// @Produce  json
// @Param id path int true "ServicePoint ID"
// @Param servicepoint body ent.ServicePoint true "ServicePoint entity"
// @Success 200 {object} ent.ServicePoint
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /servicepoints/{id} [put]
func (ctl *ServicePointController) UpdateServicePoint(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ServicePoint{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "servicepoint binding failed",
		})
		return
	}
	obj.ID = int(id)
	sp, err := ctl.client.ServicePoint.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, sp)
}

// NewServicePointController creates and registers handles for the servicepoint controller
func NewServicePointController(router gin.IRouter, client *ent.Client) *ServicePointController {
	bb := &ServicePointController{
		client: client,
		router: router,
	}
	bb.register()
	return bb
}

// InitServicePointController registers routes to the main engine
func (ctl *ServicePointController) register() {
	servicepoints := ctl.router.Group("/servicepoints")

	servicepoints.GET("", ctl.ListServicePoint)

	// CRUD
	servicepoints.POST("", ctl.CreateServicePoint)
	servicepoints.GET(":id", ctl.GetServicePoint)
	servicepoints.PUT(":id", ctl.UpdateServicePoint)
	servicepoints.DELETE(":id", ctl.DeleteServicePoint)
}

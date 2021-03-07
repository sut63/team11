package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team11/app/ent/status"

	"github.com/team11/app/ent"
	"github.com/team11/app/ent/cliententity"
	"github.com/gin-gonic/gin"
)

// ClientEntityController defines the struct for the cliententity controller
type ClientEntityController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateClientEntity handles POST requests for adding cliententity entities
// @Summary Create cliententity
// @Description Create cliententity
// @ID create-cliententity
// @Accept   json
// @Produce  json
// @Param cliententity body ent.ClientEntity true "ClientEntity entity"
// @Success 200 {object} ent.ClientEntity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cliententitys [post]
func (ctl *ClientEntityController) CreateClientEntity(c *gin.Context) {
	obj := ent.ClientEntity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "cliententity binding failed",
		})
		return
	}

	cl, err := ctl.client.ClientEntity.
		Create().
		SetCLIENTNAME(obj.CLIENTNAME).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, cl)
}

// GetClientEntity handles GET requests to retrieve a cliententity entity
// @Summary Get a cliententity entity by ID
// @Description get cliententity by ID
// @ID get-cliententity
// @Produce  json
// @Param id path int true "ClientEntity ID"
// @Success 200 {object} ent.ClientEntity
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cliententitys/{id} [get]
func (ctl *ClientEntityController) GetClientEntity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cl, err := ctl.client.ClientEntity.
		Query().
		Where(cliententity.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, cl)
}

// ListClientEntity handles request to get a list of cliententity entities
// @Summary List cliententity entities
// @Description list cliententity entities
// @ID list-cliententity
// @Produce json
// @Success 200 {array} ent.ClientEntity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cliententitys [get]
func (ctl *ClientEntityController) ListClientEntity(c *gin.Context) {
	
	cliententitys, err := ctl.client.ClientEntity.
		Query().
		WithState().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, cliententitys)
}

// DeleteClientEntity handles DELETE requests to delete a cliententity entity
// @Summary Delete a cliententity entity by ID
// @Description get cliententity by ID
// @ID delete-cliententity
// @Produce  json
// @Param id path int true "Cliententity ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cliententitys/{id} [delete]
func (ctl *ClientEntityController) DeleteClientEntity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ClientEntity.
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
//ClientEntity Struct
type ClientEntity struct {
	Sid int
}

// UpdateClientEntity handles PUT requests to update a cliententity entity
// @Summary Update a cliententity entity by ID
// @Description update cliententity by ID
// @ID update-cliententity
// @Accept   json
// @Produce  json
// @Param id path int true "ClientEntity ID"
// @Param cliententity body ClientEntity true "ClientEntity entity"
// @Success 200 {object} ent.ClientEntity
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /cliententitys/{id} [put]
func (ctl *ClientEntityController) UpdateClientEntity(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ClientEntity{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "cliententity binding failed",
		})
		return
	}
	s, err := ctl.client.Status.
		Query().
		Where(status.IDEQ(int(obj.Sid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "status not found",
		})
		return
	}

	cl, err := ctl.client.ClientEntity.
		UpdateOneID(int(id)).
		SetState(s).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cl)
}

// NewClientEntityController creates and registers handles for the cliententity controller
func NewClientEntityController(router gin.IRouter, client *ent.Client) *ClientEntityController {
	uc := &ClientEntityController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitClientEntityController registers routes to the main engine
func (ctl *ClientEntityController) register() {
	cliententitys := ctl.router.Group("/cliententitys")

	cliententitys.GET("", ctl.ListClientEntity)
	// CRUD
	cliententitys.POST("", ctl.CreateClientEntity)
	cliententitys.GET(":id", ctl.GetClientEntity)
	cliententitys.PUT(":id", ctl.UpdateClientEntity)
	cliententitys.DELETE(":id", ctl.DeleteClientEntity)
}

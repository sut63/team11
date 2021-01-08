package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nonsansap/app/ent"
	"github.com/nonsansap/app/ent/purpose"
)
// PurposeController defines the struct for the purpose controller
type PurposeController struct {
	client *ent.Client
	router gin.IRouter
}



// GetPurpose handles GET requests to retrieve a purpose entity
// @Summary Get a purpose entity by ID
// @Description get purpose by ID
// @ID get-purpose
// @Produce  json
// @Param id path int true "Purpose ID"
// @Success 200 {object} ent.Purpose
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /purposes/{id} [get]
func (ctl *PurposeController) GetPurpose(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Purpose.
		Query().
		Where(purpose.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

 // ListPurpose handles request to get a list of purpose entities
// @Summary List purpose entities
// @Description list purpose entities
// @ID list-purpose
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Purpose
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /purposes [get]
func (ctl *PurposeController) ListPurpose(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	purposes, err := ctl.client.Purpose.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, purposes)
 }

// NewPurposeController creates and registers handles for the purpose controller
func NewPurposeController(router gin.IRouter, client *ent.Client) *PurposeController {
	ppc := &PurposeController{
		client: client,
		router: router,
	}

	ppc.register()

	return ppc

}

func (ctl *PurposeController) register() {
	purposes := ctl.router.Group("/purposes")

	purposes.GET("", ctl.ListPurpose)
	purposes.GET(":id", ctl.GetPurpose)
	

}
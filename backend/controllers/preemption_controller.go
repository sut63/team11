package controllers

import (
	"fmt"
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/preemption"
	"github.com/team11/app/ent/purpose"
	"github.com/team11/app/ent/roominfo"
	"github.com/team11/app/ent/user"
)

// PreemptionController defines the struct for the preemption controller
type PreemptionController struct {
	client *ent.Client
	router gin.IRouter
}

// Preemption defines the struct for the preemption controller
type Preemption struct {
	User     int
	Roominfo int
	Purpose  int
	Otherpeopleid		string
	Otherpeoplephone	string
	Phoneuser			string
	Username	string
}

// CreatePreemption handles POST requests for adding preemption entities
// @Summary Create preemption
// @Description Create preemption
// @ID create-preemption
// @Accept   json
// @Produce  json
// @Param preemption body Preemption true "Preemption entity"
// @Success 200 {object} ent.Preemption
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /preemptions [post]
func (ctl *PreemptionController) CreatePreemption(c *gin.Context) {
	obj := Preemption{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "preemption binding failed",
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

	r, err := ctl.client.Roominfo.
		Query().
		Where(roominfo.IDEQ(int(obj.Roominfo))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "roominfo not found",
		})
		return
	}

	p, err := ctl.client.Purpose.
		Query().
		Where(purpose.IDEQ(int(obj.Purpose))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "purpose not found",
		})
		return
	}

	time := time.Now().Local()
	prm, err := ctl.client.Preemption.
		Create().
		SetPreemptTime(time).
		SetUserID(u).
		SetRoomID(r).
		SetPurposeID(p).
		SetPhonenumber(obj.Phoneuser).
		SetSurrogateid(obj.Otherpeopleid).
		SetSurrogatephone(obj.Otherpeoplephone).
		Save(context.Background())

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"status":	false,
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   prm,
	})
}

// ListPreemption handles request to get a list of preemption entities
// @Summary List preemption entities
// @Description list Preemption entities
// @ID list-preemption
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Preemption
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /preemptions [get]
func (ctl *PreemptionController) ListPreemption(c *gin.Context) {
	
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

	preemptions, err := ctl.client.Preemption.
		Query().
		Where(preemption.HasUserIDWith(user.ID(offset))).
		WithUserID().
		WithRoomID().
		WithPurposeID().
		Limit(limit).
		Offset(0).
		All(context.Background())
		
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, preemptions)
}

// GetPreemption handles GET requests to retrieve a preemption entity
// @Summary Get a preemption entity by ID
// @Description get preemption by ID
// @ID get-preemption
// @Produce  json
// @Param id path int true "Preemption ID"
// @Success 200 {object} ent.Preemption
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /preemptions/{id} [get]
func (ctl *PreemptionController) GetPreemption(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Preemption.
		Query().
		Where(preemption.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// NewPreemptionController creates and registers handles for the preemption controller
func NewPreemptionController(router gin.IRouter, client *ent.Client) *PreemptionController {
	prmc := &PreemptionController{
		client: client,
		router: router,
	}

	prmc.register()

	return prmc

}

func (ctl *PreemptionController) register() {
	preemptions := ctl.router.Group("/preemptions")

	preemptions.POST("", ctl.CreatePreemption)
	preemptions.GET("", ctl.ListPreemption)

}

package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/team11/app/ent"
	"github.com/team11/app/ent/author"
	"github.com/team11/app/ent/research"
	"github.com/team11/app/ent/researchtype"
	"github.com/team11/app/ent/user"
)

// ResearchController defines the struct for the research controller
type ResearchController struct {
	client *ent.Client
	router gin.IRouter
}

// Research struct
type Research struct {
	Register   int
	MyDoc      int
	DocType    int
	DOCNAME    string
	PAGENUMBER int
	YEARNUMBER int
	DATE       string
}

// CreateResearch handles POST requests for adding research entities
// @Summary Create research
// @Description Create research
// @ID create-research
// @Accept   json
// @Produce  json
// @Param research body Research true "Research entity"
// @Success 200 {object} ent.Research
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /researches [post]
func (ctl *ResearchController) CreateResearch(c *gin.Context) {
	obj := Research{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Research binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.Register))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	a, err := ctl.client.Author.
		Query().
		Where(author.IDEQ(int(obj.MyDoc))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	rt, err := ctl.client.Researchtype.
		Query().
		Where(researchtype.IDEQ(int(obj.DocType))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.DATE)
	r, err := ctl.client.Research.
		Create().
		SetRegister(u).
		SetMyDoc(a).
		SetDocType(rt).
		SetDOCNAME(obj.DOCNAME).
		SetPAGENUMBER(obj.PAGENUMBER).
		SetYEARNUMBER(obj.YEARNUMBER).
		SetDATE(time).
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
		"data":   r,
	})
}

// GetResearch handles GET requests to retrieve a research entity
// @Summary Get a research entity by ID
// @Description get research by ID
// @ID get-rearch
// @Produce  json
// @Param id path int true "Research ID"
// @Success 200 {object} ent.Research
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /researches/{id} [get]
func (ctl *ResearchController) GetResearch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	r, err := ctl.client.Research.
		Query().
        WithMyDoc().
		WithDocType().
		WithRegister().
		Where(research.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// GetSearchResearch handles GET requests to retrieve a research entity
// @Summary Get a research entity by Search
// @Description get research by Search
// @ID get-research-by-search
// @Produce  json
// @Param research query string false "Search Research"
// @Success 200 {object} ent.Research
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /searchresearchs [get]
func (ctl *ResearchController) GetSearchResearch(c *gin.Context) {
	researchsearch := c.Query("research")

	cr, err := ctl.client.Research.
		Query().
		WithMyDoc().
		WithDocType().
		WithRegister().
		Where(research.DOCNAMEContains(researchsearch)).
		All(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": cr,
	})
}


// ListResearch handles request to get a list of research entities
// @Summary List research entities
// @Description list research entities
// @ID list-research
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Research
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /researches [get]
func (ctl *ResearchController) ListResearch(c *gin.Context) {
	researchs, err := ctl.client.Research.
		Query().
		WithMyDoc().
		WithDocType().
		WithRegister().
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, researchs)
}

// DeleteResearch handles DELETE requests to delete a Research entity
// @Summary Delete a research entity by ID
// @Description get research by ID
// @ID delete-research
// @Produce  json
// @Param id path int true "Research ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /researches/{id} [delete]
func (ctl *ResearchController) DeleteResearch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Research.
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

// UpdateResearch handles PUT requests to update a research entity
// @Summary Update a research entity by ID
// @Description update research by ID
// @ID update-research
// @Accept   json
// @Produce  json
// @Param id path int true "Research ID"
// @Param research body ent.Research true "Research entity"
// @Success 200 {object} ent.Research
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /researches/{id} [put]
func (ctl *ResearchController) UpdateResearch(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Research{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Author binding failed",
		})
		return
	}
	obj.ID = int(id)
	r, err := ctl.client.Research.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, r)
}


// NewResearchController creates and registers handles for the research controller
func NewResearchController(router gin.IRouter, client *ent.Client) *ResearchController {
	rc := &ResearchController{
		client: client,
		router: router,
	}
	rc.register()
	return rc
}

// InitResearchController registers routes to the main engine
func (ctl *ResearchController) register() {
	researches := ctl.router.Group("/researches")
	researches.GET("", ctl.ListResearch)

    searchresearchs := ctl.router.Group("/searchresearchs")
	searchresearchs.GET("",ctl.GetSearchResearch)

	// CRUD
	researches.POST("", ctl.CreateResearch)
	researches.GET(":id", ctl.GetResearch)
	researches.PUT(":id", ctl.UpdateResearch)
	researches.DELETE(":id", ctl.DeleteResearch)
}

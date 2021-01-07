package controllers

import (
    "context"
    "fmt"
    "strconv"

    "github.com/team11/app/ent"
    "github.com/team11/app/ent/researchtype"
    "github.com/gin-gonic/gin"
)

// ResearchTypeController defines the struct for the researchtype controller
type ResearchTypeController struct {
    client *ent.Client
    router gin.IRouter
}

// CreateResearchType handles POST requests for adding researchtype entities
// @Summary Create researchtype
// @Description Create researchtype
// @ID create-researchtype
// @Accept   json
// @Produce  json
// @Param researchtype body ent.Researchtype true "Researchtype entity"
// @Success 200 {object} ent.Researchtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ResearchTypes [post]
func (ctl *ResearchTypeController) CreateResearchType(c *gin.Context) {
    obj := ent.Researchtype{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "Researchtype binding failed",
        })
        return
    }

    rt, err := ctl.client.Researchtype.
        Create().
        SetTYPENAME(obj.TYPENAME).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }

    c.JSON(200, rt)
}

// GetResearchType handles GET requests to retrieve a researchtype entity
// @Summary Get a researchtype entity by ID
// @Description get researchtype by ID
// @ID get-researchtype
// @Produce  json
// @Param id path int true "Researchtype ID"
// @Success 200 {object} ent.Researchtype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ResearchTypes/{id} [get]
func (ctl *ResearchTypeController) GetResearchType(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    rt, err := ctl.client.Researchtype.
        Query().
        Where(researchtype.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(200, rt)
}

// ListResearchType handles request to get a list of researchtype entities
// @Summary List researchtype entities
// @Description list researchtype entities
// @ID list-researchtype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Researchtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ResearchTypes [get]
func (ctl *ResearchTypeController) ListResearchType(c *gin.Context) {
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

    ResearchTypes, err := ctl.client.Researchtype.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, ResearchTypes)
}

// DeleteResearchType handles DELETE requests to delete a researchtype entity
// @Summary Delete a researchtype entity by ID
// @Description get researchtype by ID
// @ID delete-researchtype
// @Produce  json
// @Param id path int true "Researchtype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ResearchTypes/{id} [delete]
func (ctl *ResearchTypeController) DeleteResearchType(c *gin.Context) {
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

// UpdateResearchType handles PUT requests to update a researchtype entity
// @Summary Update a researchtype entity by ID
// @Description update researchtype by ID
// @ID update-researchtype
// @Accept   json
// @Produce  json
// @Param id path int true "Researchtype ID"
// @Param researchtype body ent.Researchtype true "Researchtype entity"
// @Success 200 {object} ent.Researchtype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /ResearchTypes/{id} [put]
func (ctl *ResearchTypeController) UpdateResearchType(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    obj := ent.Researchtype{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "Researchtype binding failed",
        })
        return
    }
    obj.ID = int(id)
    rt, err := ctl.client.Researchtype.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed"})
        return
    }

    c.JSON(200, rt)
}

// NewResearchTypeController creates and registers handles for the researchtype controller
func NewResearchTypeController(router gin.IRouter, client *ent.Client) *ResearchTypeController {
    rtc := &ResearchTypeController{
        client: client,
        router: router,
    }
    rtc.register()
    return rtc
}

// InitResearchTypeController registers routes to the main engine
func (ctl *ResearchTypeController) register() {
    ResearchTypes := ctl.router.Group("/Researchtypes")

    ResearchTypes.GET("", ctl.ListResearchType)

    // CRUD
    ResearchTypes.POST("", ctl.CreateResearchType)
    ResearchTypes.GET(":id", ctl.GetResearchType)
    ResearchTypes.PUT(":id", ctl.UpdateResearchType)
    ResearchTypes.DELETE(":id", ctl.DeleteResearchType)
}

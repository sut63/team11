package controllers

import (
    "context"
    "fmt"
    "strconv"

    "github.com/team11/app/ent"
    "github.com/team11/app/ent/author"
    "github.com/gin-gonic/gin"
)

// AuthorController defines the struct for the author controller
type AuthorController struct {
    client *ent.Client
    router gin.IRouter
}

// CreateAuthor handles POST requests for adding author entities
// @Summary Create author
// @Description Create author
// @ID create-author
// @Accept   json
// @Produce  json
// @Param author body ent.Author true "Author entity"
// @Success 200 {object} ent.Author
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Authors [post]
func (ctl *AuthorController) CreateAuthor(c *gin.Context) {
    obj := ent.Author{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "Author binding failed",
        })
        return
    }

    a, err := ctl.client.Author.
        Create().
        SetName(obj.Name).
        SetPosition(obj.Position).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{
            "error": "saving failed",
        })
        return
    }

    c.JSON(200, a)
}

// GetAuthor handles GET requests to retrieve a author entity
// @Summary Get a author entity by ID
// @Description get author by ID
// @ID get-author
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} ent.Author
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Authors/{id} [get]
func (ctl *AuthorController) GetAuthor(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    a, err := ctl.client.Author.
        Query().
        Where(author.IDEQ(int(id))).
        Only(context.Background())
    if err != nil {
        c.JSON(404, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(200, a)
}

// ListAuthor handles request to get a list of author entities
// @Summary List author entities
// @Description list author entities
// @ID list-author
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Author
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Authors [get]
func (ctl *AuthorController) ListAuthor(c *gin.Context) {
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

    Authors, err := ctl.client.Author.
        Query().
        Limit(limit).
        Offset(offset).
        All(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, Authors)
}

// DeleteAuthor handles DELETE requests to delete a author entity
// @Summary Delete a author entity by ID
// @Description get author by ID
// @ID delete-author
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Authors/{id} [delete]
func (ctl *AuthorController) DeleteAuthor(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    err = ctl.client.Author.
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

// UpdateAuthor handles PUT requests to update a author entity
// @Summary Update a author entity by ID
// @Description update author by ID
// @ID update-author
// @Accept   json
// @Produce  json
// @Param id path int true "Author ID"
// @Param author body ent.Author true "Author entity"
// @Success 200 {object} ent.Author
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Authors/{id} [put]
func (ctl *AuthorController) UpdateAuthor(c *gin.Context) {
    id, err := strconv.ParseInt(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(400, gin.H{
            "error": err.Error(),
        })
        return
    }

    obj := ent.Author{}
    if err := c.ShouldBind(&obj); err != nil {
        c.JSON(400, gin.H{
            "error": "Author binding failed",
        })
        return
    }
    obj.ID = int(id)
    a, err := ctl.client.Author.
        UpdateOne(&obj).
        Save(context.Background())
    if err != nil {
        c.JSON(400, gin.H{"error": "update failed"})
        return
    }

    c.JSON(200, a)
}

// NewAuthorController creates and registers handles for the author controller
func NewAuthorController(router gin.IRouter, client *ent.Client) *AuthorController {
    ac := &AuthorController{
        client: client,
        router: router,
    }
    ac.register()
    return ac
}

// InitAuthorController registers routes to the main engine
func (ctl *AuthorController) register() {
    Authors := ctl.router.Group("/Authors")

    Authors.GET("", ctl.ListAuthor)

    // CRUD
    Authors.POST("", ctl.CreateAuthor)
    Authors.GET(":id", ctl.GetAuthor)
    Authors.PUT(":id", ctl.UpdateAuthor)
    Authors.DELETE(":id", ctl.DeleteAuthor)
}

package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team11/app/ent"
   "github.com/team11/app/ent/category"
   "github.com/gin-gonic/gin"
)
 
// CategoryController defines the struct for the category controller
type CategoryController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateCategory handles POST requests for adding category entities
// @Summary Create category
// @Description Create category
// @ID create-category
// @Accept   json
// @Produce  json
// @Param category body ent.Category true "Category entity"
// @Success 200 {object} ent.Category
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /categorys [post]
func (ctl *CategoryController) CreateCategory(c *gin.Context) {
	obj := ent.Category{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "category binding failed",
		})
		return
	}
  
	u, err := ctl.client.Category.
		Create().
		SetCategoryName(obj.CategoryName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
// GetCategory handles GET requests to retrieve a category entity
// @Summary Get a category entity by ID
// @Description get category by ID
// @ID get-category
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} ent.Category
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /categorys/{id} [get]
func (ctl *CategoryController) GetCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Category.
		Query().
		Where(category.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
// ListCategory handles request to get a list of category entities
// @Summary List category entities
// @Description list category entities
// @ID list-category
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Category
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /categorys [get]
func (ctl *CategoryController) ListCategory(c *gin.Context) {
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
  
	categorys, err := ctl.client.Category.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, categorys)
 }
// DeleteCategory handles DELETE requests to delete a category entity
// @Summary Delete a category entity by ID
// @Description get category by ID
// @ID delete-category
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /categorys/{id} [delete]
func (ctl *CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Category.
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
// UpdateCategory handles PUT requests to update a category entity
// @Summary Update a category entity by ID
// @Description update category by ID
// @ID update-category
// @Accept   json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body ent.Category true "Category entity"
// @Success 200 {object} ent.Category
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /categorys/{id} [put]
func (ctl *CategoryController) UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Category{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "category binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Category.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }
// NewCategoryController creates and registers handles for the category controller
func NewCategoryController(router gin.IRouter, client *ent.Client) *CategoryController {
	uc := &CategoryController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitCategoryController registers routes to the main engine
 func (ctl *CategoryController) register() {
	categorys := ctl.router.Group("/categorys")
  
	categorys.GET("", ctl.ListCategory)
  
	// CRUD
	categorys.POST("", ctl.CreateCategory)
	categorys.GET(":id", ctl.GetCategory)
	categorys.PUT(":id", ctl.UpdateCategory)
	categorys.DELETE(":id", ctl.DeleteCategory)
 }
	   
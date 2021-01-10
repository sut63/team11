package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team11/app/ent"
   "github.com/team11/app/ent/location"
   "github.com/gin-gonic/gin"
)
 
// LocationController defines the struct for the location controller
type LocationController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateLocation handles POST requests for adding location entities
// @Summary Create location
// @Description Create location
// @ID create-location
// @Accept   json
// @Produce  json
// @Param Location body ent.Location true "location entity"
// @Success 200 {object} ent.Location
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /locations [post]
func (ctl *LocationController) CreateLocation(c *gin.Context) {
	obj := ent.Location{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "location binding failed",
		})
		return
	}
  
	cu, err := ctl.client.Location.
		Create().
		SetLocationName(obj.LocationName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, cu)
 }
// GetLocation handles GET requests to retrieve a location entity
// @Summary Get a locationr entity by ID
// @Description get location by ID
// @ID get-location
// @Produce  json
// @Param id path int true "Location ID"
// @Success 200 {object} ent.Location
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /locations/{id} [get]
func (ctl *LocationController) GetLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	cu, err := ctl.client.Location.
		Query().
		Where(location.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, cu)
 }
// ListLocation handles request to get a list of location entities
// @Summary List location entities
// @Description list location entities
// @ID list-location
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Location
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /locations [get]
func (ctl *LocationController) ListLocation(c *gin.Context) {
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
  
	locations, err := ctl.client.Location.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, locations)
 }
// DeleteLocation handles DELETE requests to delete a location entity
// @Summary Delete a location entity by ID
// @Description get location by ID
// @ID delete-location
// @Produce  json
// @Param id path int true "Location ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /locations/{id} [delete]
func (ctl *LocationController) DeleteLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Location.
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
// UpdateLocation handles PUT requests to update a location entity
// @Summary Update a location entity by ID
// @Description update location by ID
// @ID update-location
// @Accept   json
// @Produce  json
// @Param id path int true "Location ID"
// @Param location body ent.Location true "Location entity"
// @Success 200 {object} ent.Location
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /locations/{id} [put]
func (ctl *LocationController) UpdateLocation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Location{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "location binding failed",
		})
		return
	}
	obj.ID = int(id)
	cu, err := ctl.client.Location.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, cu)
 }
// NewLocationController creates and registers handles for the location controller
func NewLocationController(router gin.IRouter, client *ent.Client) *LocationController {
	cuc := &LocationController{
		client: client,
		router: router,
	}
	cuc.register()
	return cuc
 }
  
 // InitLocationController registers routes to the main engine
 func (ctl *LocationController) register() {
	locations := ctl.router.Group("/locations")
  
	locations.GET("", ctl.ListLocation)
  
	// CRUD
	locations.POST("", ctl.CreateLocation)
	locations.GET(":id", ctl.GetLocation)
	locations.PUT(":id", ctl.UpdateLocation)
	locations.DELETE(":id", ctl.DeleteLocation)
 }
package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/team11/app/ent"
   "github.com/team11/app/ent/roominfo"
   
   "github.com/gin-gonic/gin"
)
 
// RoominfoController defines the struct for the Roominfo controller
type RoominfoController struct {
   client *ent.Client
   router gin.IRouter
}


 
 // GetRoominfo handles GET requests to retrieve a roominfo entity
// @Summary Get a roominfo entity by ID
// @Description get roominfo by ID
// @ID get-roominfo
// @Produce  json
// @Param id path int true "Roominfo ID"
// @Success 200 {object} ent.Roominfo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfos/{id} [get]
func (ctl *RoominfoController) GetRoominfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Roominfo.
		Query().
		Where(roominfo.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }

 // ListRoominfo handles request to get a list of roominfo entities
// @Summary List roominfo entities
// @Description list roominfo entities
// @ID list-roominfo
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Roominfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfos [get]
func (ctl *RoominfoController) ListRoominfo(c *gin.Context) {
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
  
	roominfos, err := ctl.client.Roominfo.
		Query().
		Where(roominfo.RoomStatus("ว่าง")).
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, roominfos)
 }

 


 // DeleteRoominfo handles DELETE requests to delete a roominfo entity
// @Summary Delete a roominfo entity by ID
// @Description get roominfo by ID
// @ID delete-roominfo
// @Produce  json
// @Param id path int true "Roominfo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfos/{id} [delete]
func (ctl *RoominfoController) DeleteRoominfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Roominfo.
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

 // UpdateRoominfo handles PUT requests to update a roominfo entity
// @Summary Update a roominfo entity by ID
// @Description update roominfo by ID
// @ID update-roominfo
// @Accept   json
// @Produce  json
// @Param id path int true "Roominfo ID"
// @Param user body ent.Roominfo true "Roominfo entity"
// @Success 200 {object} ent.Roominfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roominfos/{id} [put]
func (ctl *RoominfoController) UpdateRoominfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Roominfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Roominfo binding failed",
		})
		return
	}
	obj.ID = int(id)
	cl, err := ctl.client.Roominfo.
		UpdateOne(&obj).
		SetRoomStatus(obj.RoomStatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, cl)
}

 // NewRoominfoController creates and registers handles for the roominfo controller
func NewRoominfoController(router gin.IRouter, client *ent.Client) *RoominfoController {
	uc := &RoominfoController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitRoominfoController registers routes to the main engine
 func (ctl *RoominfoController) register() {
	roominfos := ctl.router.Group("/roominfos")
  
	roominfos.GET("", ctl.ListRoominfo)
	
	// CRUD
	
	roominfos.GET(":id", ctl.GetRoominfo)
	roominfos.PUT(":id", ctl.UpdateRoominfo)
	roominfos.DELETE(":id", ctl.DeleteRoominfo)
 }
 
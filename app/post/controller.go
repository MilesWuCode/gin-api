package post

import (
	"net/http"

	"gin-api/model"
	"gin-api/plugin"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) List(c *gin.Context) {
	var service Service

	p := plugin.Pagination{Page: 1, Size: 3, Sort: 0}

	c.ShouldBind(&p)

	if list, err := service.List(p); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": list})
	}
}

func (ctrl *Controller) Create(c *gin.Context) {
	type CreateData struct {
		Title   string `form:"title" json:"title" validate:"required,max=20" label:"標題"`
		Content string `form:"email" json:"content" label:"內容"`
	}

	var data CreateData

	c.ShouldBind(&data)

	if err := plugin.Validate(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	var service Service

	var post model.Post

	post.Title = data.Title

	post.Content = data.Content

	if err := service.Create(&post); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": post})
	}
}

func (ctrl *Controller) Get(c *gin.Context) {
	id := c.Param("id")

	var service Service

	var post model.Post

	if err := service.Get(id, &post); err != nil {

		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": post})
	}
}

func (ctrl *Controller) Update(c *gin.Context) {
	type UpdateData struct {
		Title string `form:"title" json:"title" validate:"required,max=20" label:"標題"`
	}

	id := c.Param("id")

	var data UpdateData

	c.ShouldBind(&data)

	if err := plugin.Validate(data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	dataMap := plugin.StructToMapString(data)

	var service Service

	var post model.Post

	if err := service.Update(id, dataMap, &post); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": post})
	}
}

func (ctrl *Controller) Delete(c *gin.Context) {
	id := c.Param("id")

	var service Service

	if err := service.Delete(id); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}
}

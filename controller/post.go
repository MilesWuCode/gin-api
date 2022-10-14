package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
}

func (ctrl *PostController) List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

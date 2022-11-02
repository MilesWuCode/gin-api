package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) Page(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "test"})
}

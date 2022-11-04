package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"token": "token"})
}

func (ctrl *Controller) Me(c *gin.Context) {
	id, _ := c.Get("id")

	c.JSON(http.StatusOK, gin.H{"data": id})
}

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"token": "token"})
}

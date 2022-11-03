package test

import (
	"fmt"
	"gin-api/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) Page(c *gin.Context) {
	token, _ := auth.GenerateJWT()

	id, err := auth.ValidateToken(token)

	fmt.Println(id, err)

	c.JSON(http.StatusOK, gin.H{"data": token})
}

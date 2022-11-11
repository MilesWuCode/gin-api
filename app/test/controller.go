package test

import (
	"gin-api/auth"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) Page(c *gin.Context) {
	tokenDetail, _ := auth.CreateToken(3)
	auth.CreateAuth(3, tokenDetail)
}

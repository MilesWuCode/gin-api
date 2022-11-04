package auth

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	controller := Controller{}

	routerGroup := router.Group("/auth")
	{
		routerGroup.POST("/login", controller.Login)
	}
}

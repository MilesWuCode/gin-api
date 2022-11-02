package test

import (
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	controller := Controller{}

	routerGroup := router.Group("/test")
	{
		routerGroup.GET("/", controller.Page)
	}
}

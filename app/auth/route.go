package auth

import (
	"gin-api/auth"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	controller := Controller{}

	// JWT會存入redis
	routerGroup := router.Group("/auth")
	{
		routerGroup.POST("/login", controller.Login)

		// wip:套件無法針對 headers 做快取, 所以 Bearer JWT 無法做到不同人不同資料的快取
		routerGroup.GET("/me", auth.AuthMiddleware(), controller.Me)

		routerGroup.POST("/logout", auth.AuthMiddleware(), controller.Logout)

		routerGroup.POST("/refresh", auth.RefreshMiddleware(), controller.RefreshToken)
	}
}

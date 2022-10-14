package route

import (
	"gin-test/controller"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	// router.SetTrustedProxies([]string{"x.x.x.x"})
	router.SetTrustedProxies(nil)

	userController := controller.UserController{}

	user := router.Group("/user")
	{
		user.GET("/", userController.List)
		user.POST("/", userController.Create)
	}

	postController := controller.PostController{}

	post := router.Group("/post")
	{
		post.GET("/", postController.List)
	}

	return router
}

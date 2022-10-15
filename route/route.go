package route

import (
	"gin-test/controller"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Router() *gin.Engine {
	// gin.log
	// 範例 log.Printf("string")

	gin.DisableConsoleColor()

	f, _ := os.Create("./log/gin.log")

	// 寫入檔案, os.Stdout輸出畫面
	if gin.Mode() == gin.ReleaseMode {
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	// gin.router

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

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

	// log檔案位置
	f, _ := os.Create("./log/gin.log")

	// 寫入檔案, os.Stdout輸出畫面
	if gin.Mode() == gin.ReleaseMode {
		gin.DefaultWriter = io.MultiWriter(f)
	} else {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}

	// gin.router: gin.Default(), gin.New()
	// gin.Default() with Recovery & Logger
	// gin.New() witout Recovery & Logger
	router := gin.Default()

	// 若使用gin.New(),可以使用use加middleware回來
	// 或加入客製的middleware
	// router.Use(gin.Logger(), gin.Recovery())

	// router.SetTrustedProxies([]string{"x.x.x.x"})
	router.SetTrustedProxies(nil)

	userController := controller.UserController{}

	user := router.Group("/user")
	{
		user.GET("/", userController.List)
		user.POST("/", userController.Create)
		user.DELETE("/:id", userController.Delete)
	}

	postController := controller.PostController{}

	post := router.Group("/post")
	{
		post.GET("/", postController.List)
	}

	return router
}

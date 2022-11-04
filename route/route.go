package route

import (
	"io"
	"os"

	"gin-api/app/auth"
	"gin-api/app/test"
	"gin-api/app/user"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// gin.routerLog
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

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// auth
	auth.Route(router)

	// user
	user.Route(router)

	// test
	test.Route(router)

	return router
}

package main

import (
	"gin-test/database"
	"gin-test/route"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.mode: gin.TestMode, gin.DebugMode, gin.ReleaseMode
	gin.SetMode(gin.DebugMode)

	// 資料庫自動做migrate檢查
	database.AutoMigrate()

	// 啟動路由
	r := route.Router()
	// go r.RunTLS(":443", certFile, keyFile)
	r.Run(":8081")
}

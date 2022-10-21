package main

import (
	"fmt"
	"gin-api/database"
	"gin-api/plugin"
	"gin-api/route"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 初始化config
	plugin.InitConfig()

	// gin.mode: gin.TestMode, gin.DebugMode, gin.ReleaseMode
	gin.SetMode(gin.DebugMode)

	// 資料庫自動做migrate檢查
	database.AutoMigrate()

	// 路由
	r := route.Router()

	// 使用並發啟動tls
	// go r.RunTLS(":443", certFile, keyFile)

	// 啟動HTTP
	r.Run(fmt.Sprintf(":%d", viper.GetInt("gin.port")))
}

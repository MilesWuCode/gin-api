package main

import (
	"gin-test/database"
	"gin-test/route"
)

func main() {
	// 資料庫自動做migrate檢查
	database.AutoMigrate()

	// 啟動路由
	r := route.Router()
	r.Run(":8081")
}

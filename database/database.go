package database

import (
	"fmt"
	"gin-api/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// 初始化
func init() {
	user := "root"
	password := "password"
	host := "127.0.0.1"
	port := 3306
	database := "gin-api"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)
}

// 取得物件
func GetDB() *gorm.DB {
	// db.Debug()
	return db
}

// 自動轉移資料表
func AutoMigrate() {
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Tag{})
}

// 簡易查詢
func First(m interface{}) error {
	if result := db.First(&m); result.Error != nil {
		return result.Error
	}

	return nil
}

// 取得User
func GetUser(id uint) (model.User, error) {
	user := model.User{ID: id}

	if result := db.First(&user); result.Error != nil {
		return model.User{}, result.Error
	}

	return user, nil
}

package database

import (
	"gin-test/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	if err != nil {
		panic("gorm.Open()" + err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic("db.DB()" + err.Error())
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}

func Migrate() {
	db.AutoMigrate(&model.User{})
}

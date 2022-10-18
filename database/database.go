package database

import (
	"gin-api/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/plugin/dbresolver"
	"time"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	// 讀寫分離
	// db.Use(dbresolver.Register(dbresolver.Config{
	// 	// use `db2` as sources, `db3`, `db4` as replicas
	// 	Sources:  []gorm.Dialector{mysql.Open("db2_dsn")},
	// 	Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
	// 	// sources/replicas load balancing policy
	// 	Policy: dbresolver.RandomPolicy{},
	// }).Register(dbresolver.Config{
	// 	// use `db1` as sources (DB's default connection), `db5` as replicas for `User`, `Address`
	// 	Replicas: []gorm.Dialector{mysql.Open("db5_dsn")},
	// }, &User{}, &Address{}).Register(dbresolver.Config{
	// 	// use `db6`, `db7` as sources, `db8` as replicas for `orders`, `Product`
	// 	Sources:  []gorm.Dialector{mysql.Open("db6_dsn"), mysql.Open("db7_dsn")},
	// 	Replicas: []gorm.Dialector{mysql.Open("db8_dsn")},
	// }, "orders", &Product{}, "secondary"))

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

func AutoMigrate() {
	db.AutoMigrate(&model.User{})
}

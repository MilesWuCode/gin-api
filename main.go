package main

import (
	"gin-test/database"
	"gin-test/route"
)

func init() {
	database.Migrate()
}

func main() {
	r := route.Router()
	r.Run(":8081")
}

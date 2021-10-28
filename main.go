package main

import (
	"server/lib"
	"server/models"
	"server/router"
)

func main() {
	lib.InitDb()
	_ = lib.GetDBConn().AutoMigrate(&models.User{})

	router.InitRouter()

}

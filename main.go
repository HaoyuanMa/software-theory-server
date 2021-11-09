package main

import (
	"server/lib"
	"server/models"
	"server/router"
)

func main() {
	lib.InitDb()
	_ = lib.GetDBConn().AutoMigrate(&models.User{}, &models.Staff{}, &models.Record{})

	router.InitRouter()

}

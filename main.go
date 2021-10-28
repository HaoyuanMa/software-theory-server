package main

import (
	"server/router"
)

func main() {
	//lib.InitDb()
	//_ = lib.GetDBConn().AutoMigrate(&models.User{})

	router.InitRouter()

}

package api

import (
	"github.com/gin-gonic/gin"
	. "server/models"
)

func Test_get(c *gin.Context) {
	var user = &User{
		UserName: "mhy",
		Password: "password",
	}
	c.JSON(200, gin.H{
		"status ": 200,
		"message": "success",
		"user":    user,
	})
}

func Test_post(c *gin.Context) {
	var user User
	_ = c.ShouldBindJSON(&user)
	c.JSON(200, gin.H{
		"status ": 200,
		"message": "hello " + user.UserName,
	})
}

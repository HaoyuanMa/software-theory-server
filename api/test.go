package api

import (
	"github.com/gin-gonic/gin"
	. "server/models"
)

func TestGet(c *gin.Context) {
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

func TestPost(c *gin.Context) {
	var user User
	_ = c.ShouldBindJSON(&user)
	c.JSON(200, gin.H{
		"status ": 200,
		"message": "hello " + user.UserName,
	})
}

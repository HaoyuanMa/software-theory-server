package api

import (
	"github.com/gin-gonic/gin"
	"server/lib"
	"server/models"
	"server/protocol"
)

func GetNoMaskList(c *gin.Context) {
	db := lib.GetDBConn()
	var records []models.Record
	err := db.Find(&records).Error
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"data": &protocol.NoMaskList{
			Total: len(records),
			Items: records,
		},
	})
}

func GetStaffList(c *gin.Context) {
	db := lib.GetDBConn()
	var staffs []models.Staff
	err := db.Find(&staffs).Error
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "ok",
		"data": &protocol.StaffList{
			Total: len(staffs),
			Items: staffs,
		},
	})
}

func Login(c *gin.Context) {
	var login protocol.Login
	_ = c.ShouldBindJSON(login)
	db := lib.GetDBConn()
	var user models.User
	db.Where("user_name = ?", login.UserName).Find(&user)
	if user.Password == login.PassWord {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "ok",
		})
	} else {
		c.JSON(403, gin.H{
			"code":    403,
			"message": "failed",
		})
	}
}

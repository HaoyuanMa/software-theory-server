package api

import (
	"github.com/gin-gonic/gin"
	"server/lib"
	"server/models"
	"server/protocol"
)

func Forward(c *gin.Context) {

}

func Register(c *gin.Context) {
	var faceInput protocol.FaceInputProto
	_ = c.BindJSON(&faceInput)
	var curStaff models.Staff
	var db = lib.GetDBConn()
	var err = db.Where("id= ?", faceInput.Id).First(&curStaff).Updates(&models.Staff{FaceId: faceInput.FaceId}).Error
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "input failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "ok",
	})
}

func Record(c *gin.Context) {

}

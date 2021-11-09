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
	var recordInput protocol.RecordInputProto
	_ = c.BindJSON(&recordInput)
	staff := getStaffByFaceId(recordInput.FaceId)
	err := lib.GetDBConn().Create(&models.Record{
		StaffId:   staff.ID,
		StaffName: staff.Name,
	}).Error
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

func getStaffByFaceId(faceId string) models.Staff {
	var staffs []models.Staff
	db := lib.GetDBConn()
	_ = db.Find(&staffs).Error
	minDistance := -1
	resultStaff := staffs[0]
	for _, staff := range staffs {
		curDistance := calDistance(faceId, staff.FaceId)
		if curDistance < minDistance {
			minDistance = curDistance
			resultStaff = staff
		}
	}
	return resultStaff
}

func calDistance(id string, id2 string) int {
	//TODO: cal Distance
	return 0
}

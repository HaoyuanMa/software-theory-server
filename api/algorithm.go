package api

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"server/lib"
	"server/models"
	"server/protocol"
	"strconv"
)

var StreamChan chan protocol.ImageForwardProto

func init() {
	StreamChan = make(chan protocol.ImageForwardProto, 10)
}

func BuildConnection(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.Abort()
		return
	}
	ws.SetCloseHandler(func(code int, text string) error {
		close(StreamChan)
		return nil
	})
	go func() {
		for {
			data, ok := <-StreamChan
			if !ok {
				break
			}
			ws.WriteJSON(data)
		}
	}()

}

func Forward(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "input failed",
		})
		return
	}
	x1, _ := strconv.ParseFloat(c.PostForm("x1"), 64)
	x2, _ := strconv.ParseFloat(c.PostForm("x2"), 64)
	x3, _ := strconv.ParseFloat(c.PostForm("x3"), 64)
	x4, _ := strconv.ParseFloat(c.PostForm("x4"), 64)
	im, _ := strconv.ParseInt(c.PostForm("is_mask"), 10, 64)
	var isMask bool
	if im == 1 {
		isMask = false
	} else {
		isMask = true
	}

	if len(StreamChan) == cap(StreamChan) {
		//若管道已满则丢弃最旧的数据
		_ = <-StreamChan
	}

	fileContent, _ := file.Open()
	var byteContainer []byte
	byteContainer = make([]byte, 500000)
	fileContent.Read(byteContainer)

	img := base64.StdEncoding.EncodeToString(byteContainer)
	//将数据放入管道
	data := protocol.ImageForwardProto{
		Img:    img,
		X1:     x1,
		X2:     x2,
		X3:     x3,
		X4:     x4,
		IsMask: isMask,
	}
	StreamChan <- data
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

func InputRecord(c *gin.Context) {
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

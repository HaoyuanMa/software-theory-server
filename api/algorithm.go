package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math"
	"net/http"
	"server/lib"
	"server/models"
	"server/protocol"
	"strconv"
	"strings"
	"time"
)

var StreamChan chan protocol.ImageForwardProto
var count int
var lastStaffId uint
var lastStaffTime int64

func init() {
	lastStaffId = 0
	lastStaffTime = time.Now().Unix()
	count = 0
	StreamChan = make(chan protocol.ImageForwardProto, 1000)
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
		//close(StreamChan)
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
	//file, err := c.FormFile("file")
	//if err != nil {
	//	c.JSON(500, gin.H{
	//		"status":  500,
	//		"message": "input failed",
	//	})
	//	return
	//}
	x1, _ := strconv.ParseFloat(c.PostForm("x1"), 64)
	x2, _ := strconv.ParseFloat(c.PostForm("x2"), 64)
	y1, _ := strconv.ParseFloat(c.PostForm("y1"), 64)
	y2, _ := strconv.ParseFloat(c.PostForm("y2"), 64)
	im, _ := strconv.ParseInt(c.PostForm("is_mask"), 10, 64)
	img := c.PostForm("img")
	var isMask bool
	if im == 2 {
		isMask = true
	} else {
		isMask = false
	}

	if len(StreamChan) == cap(StreamChan) {
		//若管道已满则丢弃最旧的数据
		_ = <-StreamChan
	}

	//fileContent, _ := file.Open()
	//var byteContainer []byte
	//byteContainer = make([]byte, 500000)
	//fileContent.Read(byteContainer)

	//img := base64.StdEncoding.EncodeToString(byteContainer)

	//将数据放入管道
	data := protocol.ImageForwardProto{
		Img:    img,
		X1:     x1,
		X2:     x2,
		Y1:     y1,
		Y2:     y2,
		IsMask: isMask,
	}
	count += 1
	fmt.Printf("{No.%d x1: %8f x2: %8f y1: %8f y2: %8f isMask: %t}\n", count, data.X1, data.X2, data.Y1, data.Y2, data.IsMask)
	StreamChan <- data
}

func Register(c *gin.Context) {
	var faceInput protocol.FaceInputProto
	faceInput.FaceId = c.PostForm("feature")
	faceInput.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	fmt.Println(faceInput)

	var curStaff models.Staff

	var db = lib.GetDBConn()
	err := db.Where("id= ?", faceInput.Id).First(&curStaff).Updates(&models.Staff{FaceId: faceInput.FaceId}).Error
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
	recordInput.FaceId = c.PostForm("feature")
	fmt.Println(recordInput)
	staff, err := getStaffByFaceId(recordInput.FaceId)
	if err != nil {
		fmt.Println("staff not found")
		c.JSON(200, gin.H{
			"status":  200,
			"message": "ok",
		})
		return
	}
	fmt.Println("last: ")
	fmt.Println(lastStaffId)
	fmt.Println(staff.ID)
	if staff.ID == lastStaffId {
		if time.Now().Unix()-lastStaffTime > 30*60 {
			lastStaffTime = time.Now().Unix()
			lastStaffId = staff.ID
			err := lib.GetDBConn().Create(&models.Record{
				StaffId:   staff.ID,
				StaffName: staff.Name,
				Email:     staff.Email,
				Gender:    staff.Gender,
			}).Error
			if err != nil {
				c.JSON(500, gin.H{
					"status":  500,
					"message": "input failed",
				})
				return
			}
		}
	} else {
		lastStaffTime = time.Now().Unix()
		lastStaffId = staff.ID
		err := lib.GetDBConn().Create(&models.Record{
			StaffId:   staff.ID,
			StaffName: staff.Name,
			Email:     staff.Email,
			Gender:    staff.Gender,
		}).Error
		if err != nil {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "input failed",
			})
			return
		}
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "ok",
	})
}

func getStaffByFaceId(faceId string) (models.Staff, error) {
	var staffs []models.Staff
	db := lib.GetDBConn()
	_ = db.Find(&staffs).Error
	minDistance := 0.45
	resultStaff := staffs[0]
	for _, staff := range staffs {
		if len(staff.FaceId) <= 0 {
			continue
		}
		curDistance := calDistance(faceId, staff.FaceId)
		fmt.Println(curDistance)
		if curDistance < minDistance {
			minDistance = curDistance
			resultStaff = staff
		}
	}
	if minDistance >= 0.45 {
		return resultStaff, errors.New("not found")
	}
	return resultStaff, nil
}

func calDistance(id string, id2 string) float64 {
	v1 := strings.Fields(id)
	v2 := strings.Fields(id2)
	sum := 0.0
	for i := 0; i < 128; i++ {
		t1, _ := strconv.ParseFloat(v1[i], 64)
		t2, _ := strconv.ParseFloat(v2[i], 64)
		sum += math.Pow(t1-t2, 2)
	}
	return math.Pow(sum, 0.5)
}

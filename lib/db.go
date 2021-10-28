package lib

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
	"os"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open("mysql", "root:bhjx@309@tcp(10.20.13.33:3306)/software_theory?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

}

func GetDBConn() *gorm.DB {
	return db
}

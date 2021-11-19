package models

import "github.com/jinzhu/gorm"

type Staff struct {
	gorm.Model
	Name   string
	Gender string
	FaceId string `gorm:"type:Text"`
}

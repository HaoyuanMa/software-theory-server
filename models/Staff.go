package models

import "github.com/jinzhu/gorm"

type Staff struct {
	gorm.Model
	Name   string
	Gender string
	Email  string
	FaceId string `gorm:"type:Text"`
}

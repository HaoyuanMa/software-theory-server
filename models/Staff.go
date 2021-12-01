package models

import "github.com/jinzhu/gorm"

type Staff struct {
	gorm.Model
	Name       string
	Address    string
	Gender     string
	Email      string
	Age        int
	HealthCode string `gorm:"default:'green'"`
	FaceId     string `gorm:"type:Text"`
}

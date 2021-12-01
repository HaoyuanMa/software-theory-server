package models

import "github.com/jinzhu/gorm"

type Record struct {
	gorm.Model
	StaffId   uint
	StaffName string
	Email     string
	Gender    string
	IsWarning bool `gorm: "default:false"`
}

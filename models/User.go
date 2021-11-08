package models

import (
	"github.com/jinzhu/gorm"
	"server/lib"
)

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UpdateUserVM struct {
	Id      uint `json:"id"`
	NewUser User `json:"user"`
}

func hasUser(name string) bool {
	var curUser User
	var err = lib.GetDBConn().Where("user_name= ?", name).First(&curUser).Error
	if err != nil {
		return false
	}
	if curUser.ID < 0 {
		return false
	}
	return true
}

func (user User) CreatUser() bool {
	if hasUser(user.UserName) {
		return false
	}
	var err = lib.GetDBConn().Create(&user).Error
	if err != nil {
		return false
	}
	return true
}

func UpdateUser(id uint, user User) bool {

	//fmt.Println(user.UserName)

	var curUser User
	var err = lib.GetDBConn().Model(&curUser).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return false
	}

	return true
}

func DeleteUser(id uint) bool {
	var user User
	var err = lib.GetDBConn().Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return false
	}
	return true
}

func GetUser(id uint) (User, bool) {
	var curUser User
	var err = lib.GetDBConn().Where("id= ?", id).First(&curUser).Error
	if err != nil {
		return curUser, false
	}
	return curUser, true
}

func GetUsers() []User {
	var users []User
	var err = lib.GetDBConn().Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

func ValidatePassword(name string, password string) (User, bool) {
	var curUser User
	var err = lib.GetDBConn().Where("user_name = ?", name).First(&curUser).Error
	if err != nil || curUser.Password != password {
		return curUser, false
	}
	return curUser, true
}

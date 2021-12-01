package protocol

import "server/models"

type Login struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

type NoMaskList struct {
	Total int             `json:"total"`
	Items []models.Record `json:"items"`
}

type StaffList struct {
	Total int            `json:"total"`
	Items []models.Staff `json:"items"`
}

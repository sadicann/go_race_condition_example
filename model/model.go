package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name    string `json:"name"`
	Balance int    `json:"status" gorm:"default:1"`
}

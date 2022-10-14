package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" form:"name" json:"name" binding:"required" validate:"required" label:"名稱"`
	Email    string `gorm:"unique;not null" form:"email" json:"email" binding:"required" validate:"required,email" label:"帳號"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required" validate:"required" label:"密碼"`
	// Posts    []Post `gorm:"foreignkey:UserID"`
}

func (t User) TableName() string {
	return "users"
}

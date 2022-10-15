package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	// gorm.model
	ID uint `gorm:"primarykey" json:"id"`

	// custom
	Name     string `gorm:"not null" form:"name" json:"name" binding:"required" validate:"required" label:"名稱"`
	Email    string `gorm:"unique;not null" form:"email" json:"email" binding:"required" validate:"required,email" label:"帳號"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required" validate:"required" label:"密碼"`

	// gorm.model
	CreatedAt time.Time      `json:"created_at" label:"建立時間"`
	UpdatedAt time.Time      `json:"updated_at" label:"更新時間"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" label:"刪除時間"`

	// relationships
	// Posts    []Post `gorm:"foreignkey:UserID"`
}

func (t User) TableName() string {
	return "users"
}

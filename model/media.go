package model

import (
	"time"
)

type Media struct {
	// gorm欄位
	// gorm.Model
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 自訂義欄位,"-":不顯示
	Model    string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`

	// 自訂義關聯
}

// table-name
func (t *Media) TableName() string {
	// println("model.User.tablename")

	return "medias"
}

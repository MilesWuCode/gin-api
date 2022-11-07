package model

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	// gorm欄位
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// 自訂義欄位
	Title   string `gorm:"not null;comment:標題" json:"tile"`
	Content string `gorm:"type:text;comment:內文" json:"content"`
	State   bool   `gorm:"default:true" json:"state"`

	// 自訂義關聯
	UserID uint
}

// table-name
func (t *Post) TableName() string {
	return "posts"
}

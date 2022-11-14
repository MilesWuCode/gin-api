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
	Title   string `gorm:"not null;comment:標題" json:"title"`
	Content string `gorm:"type:text;comment:內文" json:"content"`
	State   bool   `gorm:"default:true" json:"state"`

	// 自訂義關聯
	UserID uint   `gorm:"not null;comment:會員" json:"user_id"`
	User   User   `gorm:"foreignkey:UserID" json:"user,omitempty"`
	Tag    []*Tag `gorm:"many2many:post_tags;" json:"tags,omitempty"`
}

// table-name
func (t *Post) TableName() string {
	return "posts"
}

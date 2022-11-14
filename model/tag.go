package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	// gorm欄位
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// 自訂義欄位
	Name string `gorm:"not null;comment:名稱" json:"name"`

	// 自訂義關聯
	Posts []*Post `gorm:"many2many:post_tags;" json:"posts,omitempty"`
}

// table-name
func (t *Tag) TableName() string {
	return "tags"
}

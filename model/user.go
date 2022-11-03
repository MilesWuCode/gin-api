package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm欄位
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// 自訂義欄位,"-":不顯示
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`

	// 自訂義關聯
	// Posts []Post `gorm:"foreignkey:UserID"`
}

// table name
func (t *User) TableName() string {
	// println("model.User.tablename")

	return "users"
}

// events:BeforeSave, BeforeUpdate, AfterSave, AfterUpdate

func (t *User) AfterUpdate(tx *gorm.DB) (err error) {
	// if t.Role == "admin" {
	// 	return errors.New("admin user not allowed to update")
	// }

	println("model.User.AfterUpdate")

	return
}

package user

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	// gorm欄位
	// gorm.Model
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// 自訂義欄位
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`

	// 自訂義關聯
	// Posts []Post `gorm:"foreignkey:UserID"`
}

// table name
func (t *UserModel) TableName() string {
	log.Println("user.model.tablename")

	return "users"
}

// events:BeforeSave, BeforeUpdate, AfterSave, AfterUpdate

func (t *UserModel) AfterUpdate(tx *gorm.DB) (err error) {
	// if t.Role == "admin" {
	// 	return errors.New("admin user not allowed to update")
	// }

	log.Println("user.model.AfterUpdate")

	return
}

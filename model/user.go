package model

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	// gorm field
	gorm.Model

	// custom field
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`

	// relationships
	// Posts    []Post `gorm:"foreignkey:UserID"`
}

// table name
func (t User) TableName() string {
	return "users"
}

// events:BeforeSave, BeforeUpdate, AfterSave, AfterUpdate

func (t *User) AfterUpdate(tx *gorm.DB) (err error) {
	// if t.Role == "admin" {
	// 	return errors.New("admin user not allowed to update")
	// }

	log.Println("User.AfterUpdate")

	return
}

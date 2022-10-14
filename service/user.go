package service

import (
	"gin-test/database"
	"gin-test/model"
	"github.com/gin-gonic/gin"
)

type UserService struct{}

type User model.User

func (service UserService) All() ([]User, error) {
	db := database.GetDB()

	var userData []User

	// if err := db.Preload("Books").Find(&userData).Error; err != nil {
	// 	return nil, err
	// }


	if err := db.Find(&userData).Error; err != nil {
		return nil, err
	}

	return userData, nil
}

// func (service UserService) Get(id string) (User, error) {
// 	db := database.GetDB()

// 	var userData User

// 	if err := db.Preload("Books").Where("id = ?", id).First(&userData).Error; err != nil {
// 		return userData, err
// 	}

// 	return userData, nil
// }

func (service UserService) Create(c *gin.Context) (User, error) {
	var userData User

	if err := c.Bind(&userData); err != nil {
		return userData, err
	}

	db := database.GetDB()

	if err := db.Create(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

// func (service UserService) Update(id string, c *gin.Context) (User, error) {
// 	db := database.GetDB()

// 	var userData User

// 	if err := db.Where("id = ?", id).First(&userData).Error; err != nil {
// 		return userData, err
// 	}

// 	if err := c.Bind(&userData); err != nil {
// 		return userData, err
// 	}

// 	db.Save(&userData)

// 	return userData, nil
// }

// func (service UserService) Delete(id string) error {
// 	db := database.GetDB()

// 	var userData User

// 	if err := db.Where("id = ?", id).Delete(&userData).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

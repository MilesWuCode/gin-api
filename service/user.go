package service

import (
	"gin-test/database"
	"gin-test/model"
)

type UserService struct{}

func (service UserService) All() ([]model.User, error) {
	db := database.GetDB()

	var user []model.User

	// 預載入
	// if err := db.Preload("Books").Find(&userData).Error; err != nil {
	// 	return nil, err
	// }

	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// func (service UserService) Get(id string) (User, error) {
// 	db := database.GetDB()

// 	var userData User

// 	if err := db.Preload("Books").Where("id = ?", id).First(&userData).Error; err != nil {
// 		return userData, err
// 	}

// 	return userData, nil
// }

func (service UserService) Create(user *model.User) error {
	db := database.GetDB()

	if err := db.Create(&user).Error; err != nil {
		return err
	} else {
		return nil
	}
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

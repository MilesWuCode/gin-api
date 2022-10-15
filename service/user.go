package service

import (
	"gin-test/database"
	"gin-test/model"
	"gin-test/plugin"
)

type UserService struct{}

func (service UserService) List(p plugin.Pagination) ([]model.User, error) {
	db := database.GetDB()

	var user []model.User

	limit, offset, err := p.Ready()

	if err != nil {
		return user, err
	}

	if err := db.Debug().Limit(limit).Offset(offset).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (service UserService) Create(user *model.User) error {
	db := database.GetDB()

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (service UserService) Get(id string, user *model.User) error {
	db := database.GetDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	return nil
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

func (service UserService) Delete(id string) error {
	db := database.GetDB()

	var user model.User

	// 使用Unscoped()查詢所有資料
	// 範例 db.Unscoped().Where("age = 20").Find(&users)
	// 執行 SELECT * FROM users WHERE age = 20;

	// 若有deleted_at則scope查詢
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
		// return errors.New("User not found")
	}

	// 永久刪除
	// 範例 db.Unscoped().Delete(&order)
	// 執行 DELETE FROM orders WHERE id=10;

	// 若有deleted_at則軟刪除
	if err := db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

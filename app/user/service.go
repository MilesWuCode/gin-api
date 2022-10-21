package user

import (
	"gin-api/database"
	"gin-api/model"
	"gin-api/plugin"
)

type Service struct{}

// 清單
func (service *Service) List(p plugin.Pagination) ([]model.User, error) {
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

// 新增
func (service *Service) Create(user *model.User) error {
	db := database.GetDB()

	if err := db.Debug().Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// 取值
func (service *Service) Get(id string, user *model.User) error {
	db := database.GetDB()

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	return nil
}

// 修改
func (service *Service) Update(id string, data map[string]interface{}, user *model.User) error {
	db := database.GetDB()

	if err := db.Debug().Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// 刪除
func (service *Service) Delete(id string) error {
	db := database.GetDB()

	var user model.User

	// 使用Unscoped()查詢所有資料
	// 範例 db.Unscoped().Where("age = 20").Find(&users)
	// 執行 SELECT * FROM users WHERE age = 20;

	// 若有deleted_at則scope查詢
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
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

package post

import (
	"gin-api/database"
	"gin-api/model"
	"gin-api/plugin"
)

type Service struct{}

// 清單
func (service *Service) List(p plugin.Pagination) ([]model.Post, error) {
	db := database.GetDB()

	var post []model.Post

	limit, offset, err := p.Ready()

	if err != nil {
		return post, err
	}

	query := db.Preload("User").Limit(limit).Offset(offset)

	// 代號轉規則
	switch p.Sort {
	case 1:
		query.Order("id desc")
	case 2:
		query.Order("updated_at desc")
	case 3:
		query.Order("id desc").Order("updated_at desc")
	}

	if err := query.Find(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

// 新增
func (service *Service) Create(post *model.Post) error {
	db := database.GetDB()

	if err := db.Create(&post).Error; err != nil {
		return err
	}

	return nil
}

// 取值
func (service *Service) Get(id interface{}, post *model.Post) error {
	db := database.GetDB()

	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		return err
	}

	return nil
}

// 修改
func (service *Service) Update(id interface{}, data map[string]interface{}, post *model.Post) error {
	db := database.GetDB()

	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		return err
	}

	if err := db.Model(&post).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// 刪除
func (service *Service) Delete(id interface{}) error {
	db := database.GetDB()

	var post model.Post

	// 使用Unscoped()查詢所有資料
	// 範例 db.Unscoped().Where("age = 20").Find(&posts)
	// 執行 SELECT * FROM posts WHERE age = 20;

	// 若有deleted_at則scope查詢
	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		return err
	}

	// 永久刪除
	// 範例 db.Unscoped().Delete(&order)
	// 執行 DELETE FROM orders WHERE id=10;

	// 若有deleted_at則軟刪除
	if err := db.Delete(&post).Error; err != nil {
		return err
	}

	return nil
}

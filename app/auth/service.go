package auth

import (
	"gin-api/auth"
	"gin-api/database"
	"gin-api/model"
)

type Service struct{}

func (service *Service) CheckIdentity(email string, password string, user *model.User) bool {
	db := database.GetDB()

	if err := db.Where("email = ?", email).Find(&user).Error; err != nil {
		return false
	}

	return auth.CheckPassword(password, user.Password)
}

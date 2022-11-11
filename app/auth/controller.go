package auth

import (
	"gin-api/app/user"
	"gin-api/auth"
	"gin-api/model"
	"gin-api/plugin"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// JWT簡單版
// func (ctrl *Controller) Login(c *gin.Context) {

// 	type FormData struct {
// 		Email    string `form:"email" json:"email" validate:"required,email" label:"帳號"`
// 		Password string `form:"password" json:"password" validate:"required" label:"密碼"`
// 	}

// 	var data FormData

// 	c.ShouldBind(&data)

// 	if err := plugin.Validate(data); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})

// 		return
// 	}

// 	var service Service

// 	var user model.User

// 	if !service.CheckIdentity(data.Email, data.Password, &user) {
// 		c.AbortWithStatus(http.StatusUnauthorized)

// 		return
// 	}

// 	token, expire, err := auth.GenerateJWT(&user)

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusInternalServerError)
// 	}

// 	c.JSON(http.StatusOK, gin.H{"type": "Bearer", "token": token, "expire": expire})
// }

// JWT簡單版
func (ctrl *Controller) Me(c *gin.Context) {
	userID := c.GetUint("userID")

	var userService user.Service

	var user model.User

	err := userService.Get(userID, &user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl *Controller) Login(c *gin.Context) {

	type FormData struct {
		Email    string `form:"email" json:"email" validate:"required,email" label:"帳號"`
		Password string `form:"password" json:"password" validate:"required" label:"密碼"`
	}

	var data FormData

	c.ShouldBind(&data)

	if err := plugin.Validate(data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	var service Service

	var user model.User

	if !service.CheckIdentity(data.Email, data.Password, &user) {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	tokenDetail, err := auth.CreateToken(user.ID)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	if err := auth.CreateAuth(user.ID, tokenDetail); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, tokenDetail)
}

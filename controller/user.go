package controller

import (
	"gin-test/model"
	"gin-test/plugin"
	"gin-test/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func init() {
	logger = plugin.Log()

	defer logger.Sync()
}

type UserController struct{}

func (ctrl *UserController) List(c *gin.Context) {
	var userService service.UserService

	// 預設值
	p := plugin.Pagination{Page: 1, Size: 3}

	// Bind query string or post data
	c.ShouldBind(&p)

	if list, err := userService.List(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": list})
	}
}

func (ctrl *UserController) Create(c *gin.Context) {
	var user model.User

	// Bind form-data request
	// c.Bind(&user)

	// Bind query string or post data
	c.ShouldBind(&user)

	// 回傳簡單錯誤,gin預設
	// if err := c.ShouldBind(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	// 	return
	// }

	// 回傳複雜錯誤,validator套件
	if err := plugin.Validate(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	var userService service.UserService

	if err := userService.Create(&user); err != nil {
		logger.Error("userService.Create", zap.String("err", err.Error()))

		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

package controller

import (
	"fmt"
	"gin-test/model"
	"gin-test/plugin"
	"gin-test/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func init() {
	// Log
	var logger = plugin.Log()

	defer logger.Sync()
}

type UserController struct{}

type Pagination struct {
	Page  int `form:"page" default:1`
	Limit int `form:"limit" default:10`
}

func (ctrl *UserController) List(c *gin.Context) {
	var userService service.UserService

	var p Pagination

	// Bind query string or post data
	c.ShouldBind(&p)

	fmt.Printf("%+v", p)

	if list, err := userService.All(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, list)
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
		c.JSON(http.StatusOK, user)
	}
}

package controller

import (
	"gin-test/model"
	"gin-test/plugin"
	"gin-test/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
}

func (ctrl *UserController) List(c *gin.Context) {

	var userService service.UserService

	userData, err := userService.All()

	if err != nil {
		// logrus.Warn(err.Error())

		c.JSON(404, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, userData)
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

	log.Printf("%T", user)

	if err := userService.Create(&user); err != nil {
		// logrus.Warn(err.Error())
		log.Println(err.Error())

		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	} else {
		log.Println(user)

		c.JSON(http.StatusOK, user)
	}
}

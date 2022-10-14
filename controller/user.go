package controller

import (
	"encoding/json"
	"gin-test/model"
	"gin-test/plugin"
	"gin-test/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (ctrl *UserController) List(c *gin.Context) {

	var userService service.UserService

	userData, err := userService.All()

	if err != nil {
		// logrus.Warn(err.Error())
		e, _ := json.Marshal(err.Error())

		c.JSON(404, gin.H{"error": e})
	} else {
		c.JSON(200, userData)
	}
}

func (ctrl *UserController) Create(c *gin.Context) {
	var user model.User

	err := plugin.Validate(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})

		return
	}

	// var userService service.UserService

	// userData, err := userService.Create(c)

	// if err != nil {
	// 	// logrus.Warn(err.Error())
	// 	fmt.Println(err)

	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	// } else {
	// 	c.JSON(201, userData)
	// }
}

package test

import (
	"fmt"
	"net/http"

	"gin-api/app/user"
	"gin-api/auth"
	"gin-api/model"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (ctrl *Controller) Page(c *gin.Context) {
	var userService user.Service

	var user model.User

	if err := userService.Get("2", &user); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	token, expire, _ := auth.GenerateJWT(&user)

	id, err := auth.ValidateJWT(token)

	fmt.Println(id, err)

	c.JSON(http.StatusOK, gin.H{"token": token, "expire": expire})
}

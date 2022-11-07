package post

import (
	"fmt"
	"gin-api/database"
	"gin-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPolicy() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Request.Method;
		id := c.Param("id")
		i, _ := strconv.ParseUint(id, 10, 64)

		// wip:換成jwt判別身份
		user := model.User{ID: uint(i)}
		database.First(&user)
		fmt.Println(user)

		item := model.User{ID: uint(i)}
		database.First(&item)
		fmt.Println(item)

		if user.ID != item.ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})

			return
		}

		c.Next()
	}
}

func UpdatePolicy() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		i, _ := strconv.ParseUint(id, 10, 64)

		// wip:換成jwt判別身份
		user := model.User{ID: uint(i)}
		database.First(&user)

		item := model.User{ID: uint(i)}
		database.First(&item)

		if user.ID != item.ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})

			return
		}

		c.Next()
	}
}

func DeletePolicy() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		i, _ := strconv.ParseUint(id, 10, 64)

		// wip:換成jwt判別身份
		user := model.User{ID: uint(i)}
		database.First(&user)

		item := model.User{ID: uint(i)}
		database.First(&item)

		if user.ID != item.ID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})

			return
		}

		c.Next()
	}
}

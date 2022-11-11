package post

import (
	"errors"
	"gin-api/database"
	"gin-api/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdatePolicy() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userID := c.GetUint("userID")

		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		item := model.Post{ID: uint(id)}

		db.First(&item)

		if userID != item.UserID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "permission denied"})

			return
		}

		c.Next()
	}
}

func DeletePolicy() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		userID := c.GetUint("userID")

		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		item := model.Post{ID: uint(id)}

		if err := db.First(&item).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}

			return
		}

		if userID != item.UserID {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Next()
	}
}

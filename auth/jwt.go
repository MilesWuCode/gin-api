package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"gin-api/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type authClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func GenerateJWT(user *model.User) (string, int64, error) {
	expire := time.Now().Add(24 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
		},
		UserID: user.ID,
	})

	tokenString, err := token.SignedString([]byte(viper.GetString("app.key")))

	if err != nil {
		return "", 0, err
	}

	return tokenString, expire, nil
}

func ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("app.key")), nil
	})

	if claims, ok := token.Claims.(*authClaims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, err
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		tokenArr := strings.Split(tokenString, " ")

		if len(tokenArr) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		authType := strings.Trim(tokenArr[0], "\n\r\t")

		if !strings.EqualFold(authType, "Bearer") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.Trim(tokenArr[1], "\n\t\r")

		if id, err := ValidateJWT(token); err == nil {
			c.Set("userID", id)
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Next()
	}
}

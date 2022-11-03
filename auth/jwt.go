package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type authClaims struct {
	jwt.StandardClaims
	UserID uint `json:"user_id"`
}

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
		UserID: 123,
	})

	tokenString, err := token.SignedString([]byte("my_secret_key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("my_secret_key"), nil
	})

	if claims, ok := token.Claims.(*authClaims); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return 0, err
	}
}

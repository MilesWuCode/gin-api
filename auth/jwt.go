package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"gin-api/database"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

type AccessClaims struct {
	jwt.StandardClaims
	UserID     uint `json:"user_id"`
	AccessUUID string
}

type RefreshClaims struct {
	jwt.StandardClaims
	UserID      uint `json:"user_id"`
	RefreshUUID string
}
type TokenDetail struct {
	Type         string `json:"type"`
	AccessToken  string `json:"access_token"`
	AccessUUID   string `json:"-"`
	AtExpires    int64  `json:"access_expire"`
	RefreshToken string `json:"refresh_token"`
	RefreshUUID  string `json:"-"`
	RtExpires    int64  `json:"refresh_expire"`
}

// 建立token
func CreateToken(userID uint) (*TokenDetail, error) {
	detail := &TokenDetail{Type: "Bearer"}
	detail.AccessUUID = uuid.New().String()
	detail.AtExpires = time.Now().Add(15 * time.Minute).Unix()
	detail.RefreshUUID = uuid.New().String()
	detail.RtExpires = time.Now().Add(7 * 24 * time.Hour).Unix()

	// AccessToken
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, AccessClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: detail.AtExpires,
		},
		UserID:     userID,
		AccessUUID: detail.AccessUUID,
	})

	// AccessToken
	if tokenString, err := accessToken.SignedString([]byte(viper.GetString("app.key"))); err == nil {
		detail.AccessToken = tokenString
	} else {
		return &TokenDetail{}, err
	}

	// RefreshToken
	refresToken := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: detail.RtExpires,
		},
		UserID:      userID,
		RefreshUUID: detail.RefreshUUID,
	})

	// RefreshToken
	if tokenString, err := refresToken.SignedString([]byte(viper.GetString("app.key"))); err == nil {
		detail.RefreshToken = tokenString
	} else {
		return &TokenDetail{}, err
	}

	return detail, nil
}

// 存入redis
func CreateAuth(userID uint, detail *TokenDetail) error {
	at := time.Unix(detail.AtExpires, 0)
	rt := time.Unix(detail.RtExpires, 0)
	now := time.Now()

	rdb := database.GetRdb()
	ctx := context.Background()

	if err := rdb.Set(ctx, detail.AccessUUID, fmt.Sprintf("%d", userID), at.Sub(now)).Err(); err != nil {
		return err
	}

	if err := rdb.Set(ctx, detail.RefreshUUID, fmt.Sprintf("%d", userID), rt.Sub(now)).Err(); err != nil {
		return err
	}

	return nil
}

type AccessDetail struct {
	UserID     uint
	AccessUUID string
}

type RefreshDetail struct {
	UserID      uint
	RefreshUUID string
}

// 提取Access
func ExtractAccessDetail(tokenString string) (*AccessDetail, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("app.key")), nil
	})

	if claims, ok := token.Claims.(*AccessClaims); ok && token.Valid {
		return &AccessDetail{
			UserID:     claims.UserID,
			AccessUUID: claims.AccessUUID,
		}, nil
	} else {
		return nil, err
	}
}

// 驗證Access
func ValidAccessToken(ad *AccessDetail) bool {
	rdb := database.GetRdb()
	ctx := context.Background()

	userid, err := rdb.Get(ctx, ad.AccessUUID).Result()

	if err != nil {
		return false
	}

	if userid != fmt.Sprintf("%d", ad.UserID) {
		return false
	}

	return true
}

// 提取Refresh
func ExtractRefreshDetail(tokenString string) (*RefreshDetail, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("app.key")), nil
	})

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return &RefreshDetail{
			UserID:      claims.UserID,
			RefreshUUID: claims.RefreshUUID,
		}, nil
	} else {
		return nil, err
	}
}

// 驗證Refresh
func ValidRefreshToken(ad *RefreshDetail) bool {
	rdb := database.GetRdb()
	ctx := context.Background()

	userid, err := rdb.Get(ctx, ad.RefreshUUID).Result()

	if err != nil {
		return false
	}

	if userid != fmt.Sprintf("%d", ad.UserID) {
		return false
	}

	return true
}

// 中介層
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

		accessDetail, err := ExtractAccessDetail(token)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		if ValidAccessToken(accessDetail) {
			c.Set("userID", accessDetail.UserID)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Next()
	}
}

// 中介層
func RefreshMiddleware() gin.HandlerFunc {
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

		refreshDetail, err := ExtractRefreshDetail(token)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		if ValidRefreshToken(refreshDetail) {
			c.Set("userID", refreshDetail.UserID)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		c.Next()
	}
}

// 刪除RedisUUID
func DeleteAuth(uuid string) bool {
	rdb := database.GetRdb()
	ctx := context.Background()

	if _, err := rdb.Del(ctx, uuid).Result(); err != nil {
		return false
	}

	return true
}

package database

import (
	"context"
	"fmt"
	"gin-api/plugin"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func init() {
	plugin.Config()

	redisAddr := fmt.Sprintf("%v:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))

	if len(redisAddr) == 0 {
		redisAddr = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(context.Background()).Result()

	if err != nil {
		panic(err)
	}
}

func GetRdb() *redis.Client {
	return rdb
}

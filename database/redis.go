package database

import (
	"context"
	"fmt"
	"gin-api/plugin"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	plugin.InitConfig()

	addr := fmt.Sprintf("%v:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))

	if len(addr) == 0 {
		addr = "localhost:6379"
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}
}

func GetRdb() *redis.Client {
	return rdb
}

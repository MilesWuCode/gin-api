package plugin

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

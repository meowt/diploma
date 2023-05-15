package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./pkg/config")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	log.Println("Config initialised")
	return
}

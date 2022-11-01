package config

import (
	"github.com/spf13/viper"
	"log"
)

func Init() {
	viper.SetConfigName("config")

	viper.AddConfigPath("config/")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file, ", err)
		return
	}

}

package main

import (
	"log"

	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./cmd/server")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Config file not found")
		} else {
			log.Printf("config read failed with error: %v", err)
		}
	}
}

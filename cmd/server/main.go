package main

import (
	"github.com/spf13/viper"
	"log"
	"net/http"
)

type ServerConfig struct {
	Port string
}

func serverConfig() ServerConfig {
	viper.SetDefault("server.port", "9090")
	return ServerConfig{
		Port: viper.GetString("server.port"),
	}
}

func main() {
	loadConfig()
	config := serverConfig()

	log.Print("Welcome to KVStore")
	log.Printf("Started on port %s", config.Port)

	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":"+config.Port, nil)
}

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	_, _ = writer.Write([]byte("pong"))
}

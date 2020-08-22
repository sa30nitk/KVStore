package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Welcome to KVStore")
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8080", nil)
}

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	_, _ = writer.Write([]byte("pong"))
}

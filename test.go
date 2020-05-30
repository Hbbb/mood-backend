package main

import (
	"log"
	"net/http"

	"server/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load environment")
	}

	http.HandleFunc("/", server.SaveMood)
	http.ListenAndServe(":80", nil)
}

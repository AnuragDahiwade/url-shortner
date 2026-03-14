package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AnuragDahiwade/url-shortner/database"
	"github.com/AnuragDahiwade/url-shortner/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	database.ConnectDB()

	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	log.Println("Server running on port", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

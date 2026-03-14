package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AnuragDahiwade/url-shortner/database"
	"github.com/AnuragDahiwade/url-shortner/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	port := ":" + os.Getenv("APP_PORT")

	database.ConnectDB()

	http.HandleFunc("/shorten", handlers.ShortenURL)
	http.HandleFunc("/", handlers.RedirectURL)

	log.Println("Server running on", port)

	log.Fatal(http.ListenAndServe(port, nil))
}

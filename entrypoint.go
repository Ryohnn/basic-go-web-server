package main

import (
	"os"

	"github.com/Ryohnn/basic-go-web-server/internal/middleware"
	"github.com/Ryohnn/basic-go-web-server/internal/mux"

	"github.com/joho/godotenv"

	"log"
	"net/http"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, relying on container environment variables")
	}
}

func main() {
	loadEnv()
	serveMux := mux.SetupRoutes()
	handler := middleware.Cors(serveMux)

	log.Fatal(
		http.ListenAndServe(
			":"+os.Getenv("APP_PORT"),
			handler,
		),
	)
}

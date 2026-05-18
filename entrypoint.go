package main

import (
	"fmt"
	"os"

	"github.com/Ryohnn/basic-go-web-server/internal/middleware"
	"github.com/Ryohnn/basic-go-web-server/internal/mux"

	"github.com/joho/godotenv"

	"log"
	"net/http"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprint(os.Stderr, "Error loading .env file")
	}
}

func main() {
	loadEnv()
	mux := mux.SetupRoutes()
	handler := middleware.Cors(mux)

	log.Fatal(
		http.ListenAndServe(
			":"+os.Getenv("APP_PORT"),
			handler,
		),
	)
}

package main

import (
	"fmt"
	"os"

	"github.com/Ryohnn/basic-go-web-server/Mux/Handlers"
	"github.com/joho/godotenv"

	"log"
	"net/http"
)

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", handlers.HomeHandler{})
	return mux
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprint(os.Stderr, "Error loading .env file")
	}
}

func main() {
	loadEnv()

	log.Fatal(
		http.ListenAndServe(
			":"+os.Getenv("LOCAL_IP"),
			setupRoutes(),
		),
	)
}

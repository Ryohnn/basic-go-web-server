package main

import (
	"log"
	"net/http"

	"github.com/Ryohnn/basic-go-web-server/Mux/Handlers"
)

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", handlers.HomeHandler{})
	return mux
}

func main() {
	mux := setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", mux))
}

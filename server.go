package main

import (
	"fmt"
	"log"
	"net/http"
)

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello!!!")
		if err != nil {
			return
		}
	})

	return mux
}

func main() {
	mux := setupRoutes()

	log.Fatal(http.ListenAndServe(":8080", mux))
}

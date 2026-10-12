package main

import (
	"os"

	"github.com/Ryohnn/pixul-be/internal/entrypoint"

	"log"
	"net/http"
)

func main() {
	log.Fatal(
		http.ListenAndServe(
			":"+os.Getenv("APP_PORT"),
			entrypoint.Setup(),
		),
	)
}

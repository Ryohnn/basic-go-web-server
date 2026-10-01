package main

import (
	"os"

	"github.com/Ryohnn/basic-go-web-server/internal/entrypoint"

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

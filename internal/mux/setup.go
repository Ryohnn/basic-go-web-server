package mux

import (
	"net/http"

	"github.com/Ryohnn/basic-go-web-server/internal/mux/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/test", handlers.TestHandler{})
	return mux
}

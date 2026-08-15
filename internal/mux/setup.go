package mux

import (
	"database/sql"
	"net/http"

	"github.com/Ryohnn/basic-go-web-server/internal/mux/handlers"
)

func SetupRoutes(DB *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/test", handlers.TestHandler{DB: DB})
	return mux
}

package mux

import (
	"net/http"

	"github.com/Ryohnn/basic-go-web-server/internal/mux/handlers"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(DB *pgx.Conn) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/games", handlers.GamesHandler{DB: DB})
	return mux
}

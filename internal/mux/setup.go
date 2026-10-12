package mux

import (
	"net/http"

	"github.com/Ryohnn/pixul-be/internal/mux/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRoutes(DB *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/games", handlers.GamesHandler{DB: DB})
	return mux
}

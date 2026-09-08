package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

type GamesHandler struct {
	DB *pgx.Conn
}

type Game struct {
	result []string
	id int64
	title string
	slug string
	description string
	published bool
	created_at time.Time
	updated_at time.Time
}

func (t GamesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rows, err := t.DB.Query(context.Background(), `SELECT * FROM "games" LIMIT 10`)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var result []string

	for rows.Next() {
		var g Game
		err := rows.Scan(&g.id, &g.title, &g.slug, &g.description, &g.published, &g.created_at, &g.updated_at)
		if (err != nil) {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = append(result, g.title)
	}

	json.NewEncoder(w).Encode(result)
}

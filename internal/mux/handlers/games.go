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
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Published   bool      `json:"publish"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (t GamesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rows, err := t.DB.Query(context.Background(), `SELECT * FROM "games" LIMIT 10`)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var result = []Game{}

	for rows.Next() {
		var g Game
		err := rows.Scan(&g.Id, &g.Title, &g.Slug, &g.Description, &g.Published, &g.CreatedAt, &g.UpdatedAt)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		result = append(result, g)
	}

	json.NewEncoder(w).Encode(result)
}

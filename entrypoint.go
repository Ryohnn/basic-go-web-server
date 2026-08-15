package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/Ryohnn/basic-go-web-server/internal/middleware"
	"github.com/Ryohnn/basic-go-web-server/internal/mux"

	"log"
	"net/http"
)

func setupDB() *sql.DB {
	dbString := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbString)

	if err != nil {
		log.Println(err.Error())
		log.Fatal("Unable to connect to DB.")
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	log.Println("Successfully connected to the database!")

	return db
}

func main() {
	serveMux := mux.SetupRoutes(setupDB())
	handler := middleware.Cors(serveMux)

	log.Fatal(
		http.ListenAndServe(
			":"+os.Getenv("APP_PORT"),
			handler,
		),
	)
}

package entrypoint

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Ryohnn/basic-go-web-server/internal/middleware"
	"github.com/Ryohnn/basic-go-web-server/internal/mux"
)

func Setup() http.Handler {
	return middleware.Cors(mux.SetupRoutes(setupDB()))
}

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

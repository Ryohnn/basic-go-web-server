package entrypoint

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Ryohnn/basic-go-web-server/internal/middleware"
	"github.com/Ryohnn/basic-go-web-server/internal/mux"
	"github.com/jackc/pgx/v5"
)

func Setup() http.Handler {
	return middleware.Cors(mux.SetupRoutes(setupDB()))
}

func setupDB() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Println(err.Error())
		log.Fatal("Unable to connect to DB.")
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to connect to DB: %v", err)
	}

	log.Println("Successfully connected to the database!")

	return db
}

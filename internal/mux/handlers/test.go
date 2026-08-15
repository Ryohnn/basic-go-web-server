package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type TestHandler struct{
	DB *sql.DB
}

func (t TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := t.DB.Ping()

	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode([]string{"Test One", "Test Two"})
}

package handlers

import (
	"encoding/json"
	"net/http"
)

type TestHandler struct{}

func (t TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]string{"Test One", "Test Two"})
}

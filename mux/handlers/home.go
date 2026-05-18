package handlers

import (
	"encoding/json"
	"net/http"
)

type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]string{"Test One", "Test Two"})
}

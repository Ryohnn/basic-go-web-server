package handlers

import (
	"fmt"
	"net/http"
)

type HomeHandler struct{}

func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello!!!")
	if err != nil {
		return
	}
}

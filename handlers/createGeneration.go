package handlers

import (
	"log"
	"net/http"
)

func CreateGeneration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	log.Printf("%s -- /api/createGeneration", r.Method)
}

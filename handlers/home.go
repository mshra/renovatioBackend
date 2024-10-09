package handlers

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><body><a href=\"https://renovatio-design.vercel.app/\">Renovatio </a></body></html>"))

	log.Printf("%s -- /", r.Method)
}

package handlers

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><body><a href=\"https://renovatio-design.vercel.app/\">Renovatio </a></body></html>"))

	log.Printf("%s -- /", r.Method)
}

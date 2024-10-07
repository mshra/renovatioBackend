package main

import (
	"errors"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><body><a href=\"https://renovatio-design.vercel.app/\">Renovatio </a></body></html>"))

	log.Printf("%s -- /", r.Method)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)

	log.Print("-- Server running on PORT:8000")
	err := http.ListenAndServe(":8000", mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server Closed!")
	} else if err != nil {
		log.Fatalf("error starting server -- %s", err)
	}
}

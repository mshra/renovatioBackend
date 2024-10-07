package main

import (
	"errors"
	"log"
	"net/http"
)

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><body>Hello, world</body></html>"))

	log.Printf("%s -- /api/greet", r.Method)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/greet", Greet)

	log.Print("-- Server running on http://localhost:8000/")
	err := http.ListenAndServe(":8000", mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server Closed!")
	} else if err != nil {
		log.Fatalf("error starting server -- %s", err)
	}
}

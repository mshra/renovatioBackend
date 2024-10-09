package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/mshra/renovatioBackend/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/api/createGeneration", handlers.CreateGeneration)

	log.Print("-- Server running on PORT:8000")
	err := http.ListenAndServe("localhost:8000", mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server Closed!")
	} else if err != nil {
		log.Fatalf("error starting server -- %s", err)
	}
}

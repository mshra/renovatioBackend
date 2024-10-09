package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mshra/renovatioBackend/handlers"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var PORT string

	PORT, isPORT := os.LookupEnv("PORT")

	if !isPORT {
		PORT = "localhost:8000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)

	log.Print("-- Server running on PORT:8000")
	err = http.ListenAndServe(PORT, mux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server Closed!")
	} else if err != nil {
		log.Fatalf("error starting server -- %s", err)
	}
}

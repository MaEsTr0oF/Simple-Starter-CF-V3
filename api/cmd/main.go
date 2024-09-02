package main

import (
	"log"
	"net/http"
	"os"

	"api/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/api", handlers.HandleCors(http.HandlerFunc(handlers.HandleUserData)))
	
	err := http.ListenAndServe("0.0.0.0:" + port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

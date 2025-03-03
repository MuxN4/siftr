package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Missing PORT")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Health Router (Subrouter for Readiness)
	healthRouter := chi.NewRouter()
	healthRouter.Get("/ready", readinessHandler)

	router.Mount("/health", healthRouter)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Server running on port", portString)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

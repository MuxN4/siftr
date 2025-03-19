package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MuxN4/siftr/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Missing PORT")
	}

	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("Missing DB_URI")
	}

	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}

	queries := db.New(conn)
	apiCfg := apiConfig{DB: queries}

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
	healthRouter.Get("/error", errorHandler)

	healthRouter.Post("/users", apiCfg.createUserHandler)
	healthRouter.Get("/users", apiCfg.middlewareAuth(apiCfg.GerUserHandler))

	healthRouter.Post("/feeds", apiCfg.middlewareAuth(apiCfg.createFeedHandler))
	healthRouter.Get("/feeds", apiCfg.getFeedsHandler)

	healthRouter.Post("/feed_followers", apiCfg.middlewareAuth(apiCfg.createFeedFollowersHandler))

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

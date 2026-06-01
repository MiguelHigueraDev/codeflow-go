package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/miguelhigueradev/codeflow/services/auth-service/internal/db"
)

type Config struct {
	Port               string
	DatabaseURL        string
	GitHubClientID     string
	GitHubClientSecret string
}

func loadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
		os.Exit(1)
	}

	githubClientID := os.Getenv("GITHUB_CLIENT_ID")
	if githubClientID == "" {
		log.Fatal("GITHUB_CLIENT_ID is not set")
		os.Exit(1)
	}

	githubClientSecret := os.Getenv("GITHUB_CLIENT_SECRET")
	if githubClientSecret == "" {
		log.Fatal("GITHUB_CLIENT_SECRET is not set")
		os.Exit(1)
	}

	return Config{
		Port:               port,
		DatabaseURL:        databaseURL,
		GitHubClientID:     githubClientID,
		GitHubClientSecret: githubClientSecret,
	}
}

func main() {
	// Load environment variables
	err := godotenv.Load("services/auth-service/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load configuration
	cfg := loadConfig()

	// Create database pool
	pool, err := db.NewPool(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	// Store

	// Server
	server := &http.Server{
		Addr: ":8080",
	}

	log.Println("server running on port 8080")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

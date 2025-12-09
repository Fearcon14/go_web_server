package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fearcon14/go_web_server/cmd/internal/config"
	"github.com/Fearcon14/go_web_server/cmd/internal/database"
	"github.com/Fearcon14/go_web_server/cmd/internal/handlers"
	"github.com/Fearcon14/go_web_server/cmd/internal/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file from current working directory (project root)
	godotenv.Load(".env")
	// Load database connection string from environment variable
	dbURL := os.Getenv("DB_URL")

	// Open database connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open database connection:", err)
	}

	// Create database queries instance
	dbQueries := database.New(db)

	cfg := &config.ApiConfig{
		DatabaseConnection: dbURL,
		DB:                 dbQueries,
	}

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("GET /api/healthz", handlers.ReadinessHandler)
	serveMux.HandleFunc("GET /admin/metrics", handlers.MetricsHandler(cfg))
	serveMux.HandleFunc("POST /admin/reset", handlers.ResetMetricsHandler(cfg))
	serveMux.HandleFunc("POST /api/validate_chirp", handlers.ValidateChirpHandler)

	serveMux.Handle("/app/", middleware.MiddlewareMetricsInc(cfg)(http.StripPrefix("/app", http.FileServer(http.Dir("../../web")))))

	server := &http.Server{
		Handler: serveMux,
		Addr:    ":8080",
	}

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}

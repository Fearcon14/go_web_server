package main

import (
	"fmt"
	"net/http"

	"github.com/Fearcon14/go_web_server/cmd/internal/config"
	"github.com/Fearcon14/go_web_server/cmd/internal/handlers"
	"github.com/Fearcon14/go_web_server/cmd/internal/middleware"
)

func main() {
	cfg := &config.ApiConfig{}

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

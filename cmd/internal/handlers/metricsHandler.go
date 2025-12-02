package handlers

import (
	"fmt"
	"net/http"

	"github.com/Fearcon14/go_web_server/cmd/internal/config"
)

func MetricsHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.FileserverHits.Load())))
	}
}

func ResetMetricsHandler(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg.FileserverHits.Store(0)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Metrics reset"))
	}
}

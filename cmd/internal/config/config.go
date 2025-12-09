package config

import (
	"sync/atomic"

	"github.com/Fearcon14/go_web_server/cmd/internal/database"
)

// ApiConfig holds application configuration and shared state
type ApiConfig struct {
	FileserverHits     atomic.Int32
	DatabaseConnection string // PostgreSQL connection string
	DB                 *database.Queries
}

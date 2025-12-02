package config

import "sync/atomic"

// ApiConfig holds application configuration and shared state
type ApiConfig struct {
	FileserverHits atomic.Int32
}

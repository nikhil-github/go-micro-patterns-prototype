package database

import "github.com/yourusername/shared-foundation/core"

// Database interface for data persistence
type Database interface {
	core.Service
	Connect() error
	Disconnect() error
	Health() error
	// Add more database-specific methods as needed
	// Query(query string, args ...interface{}) (Result, error)
	// Transaction() (Transaction, error)
}

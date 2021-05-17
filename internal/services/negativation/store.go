package negativations

import "database/sql"

type INegativationStore interface {
	NewStore(db *sql.DB) *Store
}

// Store provides all functions to execute db queries
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore instantiates a NegativationStore.
func NewStore(db *sql.DB) *Store {
	negStore := &Store{
		Queries: New(db),
		db:      db,
	}

	return negStore
}

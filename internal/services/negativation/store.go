package negativations

import "database/sql"

type Store interface {
	Querier
	//NewStore(db *sql.DB) Store
}

// Store provides all functions to execute db queries
type sNegativationStore struct {
	*Queries
	db *sql.DB
}

// NewStore instantiates a NegativationStore.
func NewStore(db *sql.DB) Store {
	negStore := &sNegativationStore{
		Queries: New(db),
		db:      db,
	}

	return negStore
}

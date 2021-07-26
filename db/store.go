// Package db implements a database interface for storing user and puzzle data.
package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store represents a datastore object with a SQL database, for easier execution of transactions.
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore returns a reference to a Store object, constructed with db.
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// execTx executes txFunc (and its queries) as one atomic transaction, with context.
func (s *Store) execTx(ctx context.Context, txFunc func(*Queries) error) error {
	// Start a transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Execute the passed queries func as a transaction
	queries := New(tx)
	if err := txFunc(queries); err != nil {
		// Error occured, rollback changes
		if rbError := tx.Rollback(); rbError != nil {
			// Rollback error occured (in addition to a tx error)
			return fmt.Errorf("rollback err: %w, tx err: %v", rbError, err)
		}
		return err
	}

	// Commit transaction
	return tx.Commit()
}

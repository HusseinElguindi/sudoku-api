package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (s *Store) execTx(ctx context.Context, txFunc func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	if err := txFunc(queries); err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("rollback err: %w, tx err: %v", rbError, err)
		}
		return err
	}

	return tx.Commit()
}

package databasedriver

import (
	"context"
	"social/internal/store"

	"gorm.io/gorm"
)

var _ store.Store = (*Store)(nil)

type Store struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) NewTx(ctx context.Context) (store.Store, error) {
	tx := s.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &Store{db: tx}, nil
}

func (s *Store) Rollback() error {
	return s.db.Rollback().Error
}

func (s *Store) Commit() error {
	return s.db.Commit().Error
}

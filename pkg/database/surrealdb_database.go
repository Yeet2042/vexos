package database

import (
	"context"
	"fmt"

	surrealdb "github.com/surrealdb/surrealdb.c.go"
)

type SurrealdbConfig struct {
	Path      string
	Namespace string
	Database  string
}

type Surrealdb struct {
	db surrealdb.Driver
}

func NewSurrealdbInstance(
	ctx context.Context,
	config *SurrealdbConfig,
) (
	DatabaseInstance,
	error,
) {
	if config == nil {
		return nil, fmt.Errorf("[pkg/database]: config is nil")
	}

	dbURL := fmt.Sprintf("surrealkv://%s", config.Path)

	db, err := surrealdb.Open(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("[pkg/database]: failed to open embedded db: %w", err)
	}

	if err = db.Use(ctx, config.Namespace, config.Database); err != nil {
		db.Close()
		return nil, fmt.Errorf("[pkg/database]: failed to select namespace and database: %w", err)
	}

	return &Surrealdb{db: db}, nil
}

func (s *Surrealdb) GetDB() surrealdb.Driver {
	return s.db
}

func (s *Surrealdb) Close() {
	if s.db != nil {
		s.GetDB().Close()
	}
}

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

type sdb struct {
	db surrealdb.Driver
}

func New(
	ctx context.Context,
	config *SurrealdbConfig,
) (
	Database,
	error,
) {
	if config == nil {
		return nil, fmt.Errorf("[pkg/database]: config is nil")
	} else if config.Namespace == "" {
		return nil, fmt.Errorf("[pkg/database]: namespace is empty")
	} else if config.Database == "" {
		return nil, fmt.Errorf("[pkg/database]: database is empty")
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

	return &sdb{db: db}, nil
}

func (s *sdb) GetDB() surrealdb.Driver {
	return s.db
}

func (s *sdb) Close() {
	if s.db != nil {
		s.GetDB().Close()
	}
}

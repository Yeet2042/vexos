package database

import surrealdb "github.com/surrealdb/surrealdb.c.go"

type DatabaseInstance interface {
	GetDB() surrealdb.Driver
	Close()
}

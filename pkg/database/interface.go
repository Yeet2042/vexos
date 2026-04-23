package database

import surrealdb "github.com/surrealdb/surrealdb.c.go"

type Database interface {
	GetDB() surrealdb.Driver
	Close()
}

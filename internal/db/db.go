package db

import (
	"fmt"
)

type DBModel interface {
	Create(string) string
}

type DB struct {}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) CreateDB(data string) string {
	return fmt.SPrintf("%s %s", "inserted", data)
}

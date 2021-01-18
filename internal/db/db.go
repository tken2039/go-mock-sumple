package db

import (
	"fmt"
)

type DBModel interface {
	CreateDB(string) string
}

type DB struct {}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) CreateDB(data string) string {
	return fmt.Sprintf("%s %s", "inserted", data)
}

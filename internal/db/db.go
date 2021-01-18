package db

import (
	"fmt"
)

// DBModel は CreateDB 関数を宣言する interface
type DBModel interface {
	CreateDB(string) string
}

// DB は CreateDB メソッドを持つクラス
type DB struct {}

// NewDB は *DB を返却する
func NewDB() *DB {
	return &DB{}
}

// CreateDB は DB クラスのメソッド
// このメソッドを実装している DB クラスは、暗黙的に interface "DBModel" を実装している
func (db *DB) CreateDB(data string) string {
	return fmt.Sprintf("%s %s", "inserted", data)
}

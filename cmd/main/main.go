package main

import (
	"net/http"
	"github.com/tken2039/go-mock-sumple/pkg/db"
	"github.com/tken2039/go-mock-sumple/pkg/handler"
)

func main() {
	// DB を取得
	// DB は interface "DBModel" を暗黙的に実装している
	d := db.NewDB()

	// Handler を初期化して取得
	// NewHandler は interface "DBModel" を引数に取る関数であるので
	// ここでは、DBModel を実装する DB 型の変数 d を引数に取ることができている
	h := handler.NewHandler(d)

	// h は Handler クラスなので、 CreateHandler メソッドを呼び出すことができる
	http.HandleFunc("/", h.CreateHandler)
	http.ListenAndServe(":8080", nil)
}

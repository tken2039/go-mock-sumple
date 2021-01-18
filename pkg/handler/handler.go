package handler

import (
	"fmt"
        "net/http"
        "github.com/tken2039/go-mock-sumple/internal/db"
)

// Handler は interface "DBModel" 型の変数をメンバに持つ構造体
type Handler struct {
        db db.DBModel
}

// NewHandler は interface "DBModel" を引数に持ち、 Handler を返却する
// 引数が interface であるということは、この interface （つまりは DBModel ）
// を実装する型の変数を自由に引数に取ることができる
func NewHandler(d db.DBModel) *Handler {
        return &Handler{d}
}

// CreateHandler は Handler クラスのメソッド
func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	data := "Data"
	fmt.Fprintf(w, "%s", h.db.CreateDB(data))
}

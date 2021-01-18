package handler

import (
	"fmt"
        "net/http"
        "github.com/tken2039/go-mock-sumple/internal/db"
)

type Handler struct {
        db db.DBModel
}

func NewHandler(d db.DBModel) *Handler {
        return &Handler{d}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	data := "Data"
	fmt.Fprintf(w, "%s", h.db.CreateDB(data))
}

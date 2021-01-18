package main

import (
        "net/http"
        "internal/db"
)

type Handler struct {
        db db.DBModel
}

func NewHandler(d db.DBModel) *Handler {
        return &Handler{d}
}

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
        result := h.DBModel.Create("data")
}

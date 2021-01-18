package main

import (
	"net/http"
	"github.com/tken2039/go-mock-sumple/internal/db"
	"github.com/tken2039/go-mock-sumple/internal/handler"
)

func main() {
	d := db.NewDB()
	h := handler.NewHandler(d)
	http.HandleFunc("/", h.CreateHandler)
	http.ListenAndServe(":8080", nil)
}

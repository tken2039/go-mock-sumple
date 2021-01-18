package main

import (
	"net/http"
	"internal/db"
	"internal/handler"
)

func main() {
	d := db.NewDB()
	h := handler.NewHandler(d)
	http.HandleFunc("/", h.CreateHandler)
	http.ListenAndServe(":8080", nil)
}

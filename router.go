package main

import (
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /posts", GetAllHandler)

	router.HandleFunc("GET /posts/{id}", FindPostHandler)

	router.HandleFunc("POST /posts", AddPost)

	return router
}

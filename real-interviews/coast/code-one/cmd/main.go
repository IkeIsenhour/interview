package main

import (
	"fmt"
	"log"
	"net/http"

	"code-one/internal/api/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shortener", handler.ShortenerPost)
	mux.HandleFunc("GET /shortener/{url}", handler.ShortenerGet)

	fmt.Println("listening on port :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

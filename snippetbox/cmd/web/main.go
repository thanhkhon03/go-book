package main

import (
	"log"
	"net/http"
	"snippetBox/internal/pkg"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pkg.Home)
	mux.HandleFunc("/snippet/view", pkg.SnippetView)
	mux.HandleFunc("/snippet/create", pkg.SnippetCreate)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

package main

import (
	"log"
	"net/http"
	"snippetBox/internal/pkg"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", pkg.Home)
	mux.HandleFunc("/snippet/view", pkg.SnippetView)
	mux.HandleFunc("/snippet/create", pkg.SnippetCreate)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

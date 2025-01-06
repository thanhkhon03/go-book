package main

import (
	"flag"
	"log"
	"net/http"
	"snippetBox/internal/pkg"
)

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", pkg.Home)
	mux.HandleFunc("/snippet/view", pkg.SnippetView)
	mux.HandleFunc("/snippet/create", pkg.SnippetCreate)

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

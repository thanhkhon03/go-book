package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Starting server on :8080") // Chay tai host:8080
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

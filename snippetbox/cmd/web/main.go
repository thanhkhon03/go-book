package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"snippetBox/internal/pkg"

)

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address") // Them addr...
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", pkg.Home)
	mux.HandleFunc("/snippet/view", pkg.SnippetView)
	mux.HandleFunc("/snippet/create", pkg.SnippetCreate)

	str := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := str.ListenAndServe()
	errorLog.Fatal(err)
}

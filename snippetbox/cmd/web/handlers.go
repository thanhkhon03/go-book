/*package main

import (
	"fmt"
	"html/template" // new import 
	"log" // new import
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
		ts, err := template.ParseFiles("./ui/html/pages/home.html")
        if err != nil {
             log.Println(err.Error())
             http.Error(w, "Internal Server Error", 500)
             return
}
			 err = ts.Execute(w, nil)
			 ìf err != nil {
			 log.Println(err.Error())
			 http.Error(w, "Internal Server Error", 500)
	}		 
}

	w.Write([]byte("Hello from Snippetbox"))
}

func SnippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}

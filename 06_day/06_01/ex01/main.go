package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/snippet/create", createSnippet)
	http.HandleFunc("/snippet/write", writeSnippet)
	http.HandleFunc("/showpost", showSnippet)

	fileServer := http.FileServer(http.Dir("./"))

	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Go Backend: { HTTPVersion = 1 }; serving on http://localhost:8888/")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/", home)

	log.Println("Listening on :8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

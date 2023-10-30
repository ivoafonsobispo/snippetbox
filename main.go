package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	addr := ":4000"
	log.Println("starting server on ", addr)

	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/ep", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

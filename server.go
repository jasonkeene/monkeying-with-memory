package main

import (
	"log"
	"net/http"

	"github.com/jasonkeene/slides"
)

func main() {
	h := slides.Handler("Monkeying with Memory", "README.md")
	fsh := http.FileServer(http.Dir("."))
	mux := http.NewServeMux()
	mux.Handle("/", h)
	mux.Handle("/assets/", fsh)

	addr := "localhost:8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

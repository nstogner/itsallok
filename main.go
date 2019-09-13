package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("listening on address ($ADDR) %q", addr)
	http.ListenAndServe(addr, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("received request: %v", r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("It's all OK.\n"))
		},
	))
}

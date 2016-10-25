package main

import (
	"fmt"
	"net/http"
)

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello!")
}

type worldHandler struct{}

func (h *worldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "World")
}

func main() {
	hello := helloHandler{}
	world := worldHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	// createCertAndKey()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServe()
}

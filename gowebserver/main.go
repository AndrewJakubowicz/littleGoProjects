package main

import "net/http"

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	handler := myHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	// createCertAndKey()
	// server.ListenAndServeTLS("cert.pem", "key.pem")
	server.ListenAndServe()
}

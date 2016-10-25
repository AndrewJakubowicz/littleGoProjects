package main

import (
	"fmt"
	"net/http"
)

// Returns all the headers of a request.
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

// Shows how to access individual request headers.
func encodingHeader(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, "r.Header[\"Accept-Encoding\"] gives:", h)
	h2 := r.Header.Get("Accept-Encoding")
	fmt.Fprintln(w, "r.Header.Get(\"Accept-Encoding\") gives:", h2)
}

func reqBody(w http.ResponseWriter, r *http.Request) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/encoding", encodingHeader)
	http.HandleFunc("/reqBody", reqBody)
	http.HandleFunc("/process", process)

	fmt.Println("Starting server on [", server.Addr, "]")
	server.ListenAndServe()
}

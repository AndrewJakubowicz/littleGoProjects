package main

import (
	"fmt"
	"io/ioutil"
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

// process the form.
func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, "Using r.ParseForm():\n============")
	fmt.Fprintln(w, r.Form, "\n\nUse r.PostForm() to ignore URL key-value pairs.\n============\n", r.PostForm)
	fmt.Fprintln(w, "\n\nr.Form[\"FirstName\"]\n============\n", r.Form["FirstName"])
}

// This function returns the form.
func serveForm(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("minimalRequest.html")
	fmt.Fprintf(w, string(file))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/encoding", encodingHeader)
	http.HandleFunc("/reqBody", reqBody)
	http.HandleFunc("/process", process)
	http.HandleFunc("/", serveForm)

	fmt.Println("Starting server on [", server.Addr, "]")
	server.ListenAndServe()
}

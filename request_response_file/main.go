package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Process a multipart form.
func process(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if len(r.MultipartForm.File["uploaded"]) < 1 {
		http.Error(w, "Error: Please submit a file!", 400)
		return
	}
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, string(data))

}

// This function returns the form.
func serveIndex(w http.ResponseWriter, r *http.Request) {
	file, _ := ioutil.ReadFile("index.html")
	w.Write(file)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/", serveIndex)

	fmt.Println("Starting server on [", server.Addr, "]")
	server.ListenAndServe()
}

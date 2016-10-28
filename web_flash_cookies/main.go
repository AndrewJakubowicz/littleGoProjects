package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// This is where a cookie is set with a message.
// This cookie is a session cookie, therefore it'll persist as long as browser stays open.
// However if the showMessage function below is called, the cookie will be used up.
func setMessage(w http.ResponseWriter, r *http.Request) {
	m := []byte("Super secret message!")
	c := http.Cookie{
		Name:  "flashCookie",
		Value: base64.URLEncoding.EncodeToString(m),
	}
	http.SetCookie(w, &c)
}

// showMessage uses up the "flashCookie" cookie and displays its message.
func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flashCookie")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "There is no message! It is gone!!! >.<")
			return
		}
		http.Error(w, "Unknown cookie related error!", 500)
		return
	}
	// We have a cookie.
	// Let's replace it with an expired cookie (with same name) to delete the cookie.
	killCookie := http.Cookie{
		Name:    "flashCookie",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	}
	http.SetCookie(w, &killCookie)

	// Finally, we'll show the user the flash message.
	m, err := base64.URLEncoding.DecodeString(c.Value)
	if err != nil {
		http.Error(w, "Cookie message decoding error!", 500)
		return
	}

	fmt.Fprintln(w, string(m))
}

func main() {
	s := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set", setMessage)
	http.HandleFunc("/get", showMessage)

	fmt.Println("Starting server on [", s.Addr, "]...")
	s.ListenAndServe()
}

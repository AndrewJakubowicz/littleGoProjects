// The handlers you see here are heavily inspired by:
// http://www.alexedwards.net/blog/a-recap-of-request-handling
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

// This should be shared between the handler functions.
// Ideal over global scope.
type cookieHandler struct {
	value string
}

// setCookie sets and sends a cookie to the client.
func (c *cookieHandler) setCookie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// First look for name in url.
		r.ParseForm()
		name := r.FormValue("name")
		if name == "" {
			// Assign default cookie name if none supplied.
			name = "DefaultName"
		}

		c1 := http.Cookie{
			Name:     name,
			Value:    c.value,
			HttpOnly: true,
		}
		http.SetCookie(w, &c1)
		fmt.Fprintln(w, "Set cookie! Check out /get to see the set cookie.")
		fmt.Fprintln(w, "Set your own cookie name using, /set?name=<nameHere>\n\nHave fun! :3")
	}
}

// Gets all the cookies and displays them in the browser.
func (*cookieHandler) getCookie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Your Cookies:\n\n", r.Cookies())
	}
}

// Prints to the terminal the name of the function called.
func logFunc(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Printing connections to terminal.
		fmt.Print("Incomming request from:\n", r.Header, "\n\n=================\n")
		fmt.Println("Calling handler: ", getFunctionName(h))
		fmt.Print("=================\n\n")
		h(w, r)
	}
}

// Returns the name of the function.
// http://stackoverflow.com/a/7053871
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func main() {
	// Sets up server.
	context := cookieHandler{value: "A fun context!"}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	// Handlers for get/set of cookies.
	// Handlers are chained through a trivial log function.
	http.Handle("/set", logFunc(context.setCookie()))
	http.Handle("/get", logFunc(context.getCookie()))

	// Start server
	fmt.Println("Running server on [", server.Addr, "]...")
	server.ListenAndServe()
}

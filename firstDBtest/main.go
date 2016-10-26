// This is a toy application and not suitable for real use.
package main

import "net/http"
import "fmt"
import "log"
import "github.com/boltdb/bolt"

var DB *bolt.DB

var somebucket = []byte("someBucket")

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("USAGE:\n\n\n/add?key=<yourKey>&value=<someValue>\n\n/view?key=<keyToRetrieve>"))
}

// Adds a user defined value to the database.
func add(w http.ResponseWriter, r *http.Request) {
	// Simple key/value addition.
	r.ParseForm()
	key := r.Form["key"][0]
	value := r.Form["value"][0]
	fmt.Fprintln(w, "USAGE: /add?key=<yourKey>&value=<someValue>\n\nkey:", key, "\nvalue:", value)

	// Early escape if no key/value
	if len(key) == 0 && len(value) == 0 {
		http.Error(w, "Please provide BOTH key and value", 400)
		return
	}

	err := DB.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(somebucket)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintln(w, "Successfully added \"", key, ":", value, "\"")
}

// Returns a user defined value from the database.
func view(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "USAGE: /view?key=<keyToRetrieve>")

	// Grab key
	r.ParseForm()
	key := r.Form["key"][0]

	// Retrieve value
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(somebucket)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", somebucket)
		}
		value := bucket.Get([]byte(key))
		fmt.Fprintln(w, "Found your value:", string(value))
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	DB = db // store open db in global scope.
	defer db.Close()

	http.HandleFunc("/", index)
	http.HandleFunc("/add", add)
	http.HandleFunc("/view", view)
	fmt.Println("Starting server....")
	server.ListenAndServe()
}

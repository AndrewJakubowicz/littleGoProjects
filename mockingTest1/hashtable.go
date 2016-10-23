package mockingTest1

import "errors"

var (
	// ErrNotFound returned if Get method fails to lookup a value.
	ErrNotFound = errors.New("not found")
)

// HashTable is an interface for a simple hash table.
// It's designed so that some libraries immediately adhere to it.
// E.g. http://godoc.org/github.com/hoisie/redis
type HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

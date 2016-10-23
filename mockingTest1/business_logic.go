package mockingTest1

// BusinessLogic is a sample function doing business transactions.
func BusinessLogic(h HashTable) {
	h.Set("hello", []byte("world"))
}

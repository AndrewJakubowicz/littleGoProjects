// Small internet cafe has 8 computers.
// Computers available as a first come, first serve basis.
// When all computers are taken, the next person waits till one is available.
// 25 people are waiting.
// Each person spends between 15 minutes and 2 hours online.
// Simulate!
// Source: whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Refactored by William Kennedy.
func main() {

	// How many people want in?
	people := rand.Perm(25)

	// Create the waiting list.
	waitingList := make(chan int)

	// Set the waitgroup for the number of people
	// to wait on before ending the app.
	var wg sync.WaitGroup
	wg.Add(len(people))

	// Add the 8 computers into the room.
	for i := 0; i < 8; i++ {

		// One goroutine per computer.
		go func() {

			// Wait for people to request a computer.
			for p := range waitingList {

				// Calculate stats.
				t := (rand.Int() % 105) + 15
				fmt.Printf("Tourist %d is online\n", p)
				time.Sleep(time.Duration(t*10) * time.Millisecond)
				fmt.Printf("Tourist %d spent %d minutes online\n", p, t)

				// Report the user is done with this computer.
				wg.Done()
			}
		}()
	}

	// Start serving computers to people
	for i := 0; i < len(people); i++ {
		waitingList <- people[i]
	}

	// All people have been served.
	fmt.Println("Closing the waiting list")
	close(waitingList)

	// Wait for all of them to report they are
	// done using the computer.
	wg.Wait()
}

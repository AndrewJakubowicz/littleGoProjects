// Exercise 3: Internet Cafe
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

var wg sync.WaitGroup

func main() {
	waitingList := make(chan int)
	people := rand.Perm(25)
	wg.Add(25)
	// This is the line of people waiting to get on the computer.
	go allocatePeople(people, waitingList)

	// Initiate the 8 computers.
	for i := 0; i < 8; i++ {
		go useComputer(waitingList)
	}
	wg.Wait()
}

func allocatePeople(people []int, c chan int) {
	for i := 0; i < len(people); i++ {
		c <- people[i]
	}
}

func useComputer(c chan int) {
	for {
		person := <-c
		t := (rand.Int() % 105) + 15
		fmt.Printf("Tourist %d is online\n", person)
		time.Sleep(time.Duration(t) * time.Millisecond)
		fmt.Printf("Tourist %d spent %d minutes online\n", person, t)
		wg.Done()
	}
}

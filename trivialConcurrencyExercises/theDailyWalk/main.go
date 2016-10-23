// Exercise 1: The Daily Walk.
// Every morning, Alice and Bob go for a walk.
// They both grab sunglasses, perhaps a belt, closing open windows, turning off ceiling fans,
// and pocketing their phones and keys.
// Once they're BOTH ready, which typically takes between 60 - 90 seconds,
// they arm the alarm, which has a 60 second delay.
// While the alarm counts down they both put on their shoes (35 - 45 seconds).
// They then leave the house together and lock the door before the alarm has finished the countdown.
// Source: whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Getting ready
	wg.Add(2)
	go getReady("Alice", &wg)
	go getReady("Bob", &wg)
	wg.Wait()

	// Arm alarm
	fmt.Println("Arming alarm!")
	go armAlarm()

	// Put on shoes
	wg.Add(2)
	go putOnShoes("Alice", &wg)
	go putOnShoes("Bob", &wg)

	wg.Wait()

	fmt.Println("Exiting and locking the door")

	runtime.Goexit()
}

func getReady(name string, w *sync.WaitGroup) {
	defer w.Done()
	randomNum := rand.Intn(31) + 60
	fmt.Printf("%s started getting ready.\n", name)
	time.Sleep(time.Millisecond * time.Duration(randomNum))
	fmt.Printf("%s spent %d seconds getting ready.\n", name, randomNum)
}

func armAlarm() {
	fmt.Println("Alarm counting down.")
	time.Sleep(60 * time.Millisecond)
	fmt.Println("Alarm is armed.")

	os.Exit(0) // Exit program
}

func putOnShoes(name string, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Printf("%s started putting on shoes\n", name)
	timeSpent := rand.Intn(11) + 35
	time.Sleep(time.Duration(timeSpent) * time.Millisecond)
	fmt.Printf("%s spent %d seconds putting on shoes.\n", name, timeSpent)
}

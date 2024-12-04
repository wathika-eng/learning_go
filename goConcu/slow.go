package main

import (
	"fmt"
	"time"
)

// blocks greet, until it runs
//
// we can defer it or use go routines
func slowGreet(phrase string, doneChan chan bool) {
	// simulated, can be an API call or calculation
	time.Sleep(2 * time.Second)
	fmt.Println("Hello, ", phrase)
	// once done, send data back to
	doneChan <- true
	// once done close
	close(doneChan)
}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hey, ", phrase)
	doneChan <- true
}

// avoid race condition, where programs compete to finish first
// wait multiple times

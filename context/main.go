package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func openConnections(done chan bool) {
	fmt.Println("Attempting to connect...")
	if rand.Intn(100) <= 50 {
		fmt.Println("Oops! Error occurred")
		time.Sleep(10000 * time.Hour)
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("Connected...")
	}
	done <- true
}

func openWithTimeout() {
	// wrap context with timeout, cancel function must be called before
	// use defer to cancel context awaiting
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// run as a go routine
	done := make(chan bool)
	go openConnections(done)
	// select checks the 2 cases and which returns first
	select {
	case <-done:
		fmt.Println("Connection successful")
	case <-ctx.Done():
		fmt.Println("Timeout")
	}
	close(done)
}

func main() {
	openWithTimeout()
	//fmt.Println(rand.Intn(100))
}

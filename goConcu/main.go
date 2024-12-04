package main

func main() {
	// use slices to avoid race
	// a
	// dones := make([]chan bool, 2)
	done := make(chan bool)
	go slowGreet("Mia", done)
	// doesn't depend on slowGreet
	//done = make(chan bool)
	go greet("Kamaa", done)
	// variable := data type to be sent through channel
	// now reverse arrow to read data being sent
	// go will end once it receives something from the channel
	// for _, done := range dones {
	// 	<-done
	// }
	// fmt.Println(done)
	// range over channels, avoid deadlock (not knowing when done)
	// close the slowest 
	for range done {
		//
	}
}

// go routines are parallel, non-blocking
// doesn't return a value or wait for return
// dispatch and go
// we use channels instead to transmit data

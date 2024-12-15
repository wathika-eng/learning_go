package main

import "fmt"

func panicPeter() {
	fmt.Println("Hello Peter")

	//must be inside defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("End panic")
		}
	}()
	// not advised
	// will end program immediately but allow defer to run before end
	panic("Panicked")
}

func main() {
	panicPeter()
	fmt.Println("The end...")
}

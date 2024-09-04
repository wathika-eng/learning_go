package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var a string
	// fmt.Printf("%T\t %v \n", a, a)
	reader := bufio.NewReader(os.Stdin)

	// Read input from the user
	input, err := reader.ReadString('\n')
	if err != nil { // Check for errors during input reading
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert input to a floating-point number
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil { // If error is present, print it
		fmt.Println("Error converting input to number:", err)
	} else {
		// If no error, print the number incremented by 1
		fmt.Println("Added to", numRating+1)
	}
}

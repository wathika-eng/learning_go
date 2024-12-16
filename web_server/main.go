package main

import "fmt"

func main() {
	var result []int // Slice to hold the results
	var i uint64
	// Sequentially append values to the slice
	for i = 0; i < 18446744073709551615; i++ {
		result = append(result, int(i)) // Append to the slice
	}

	// Print the result slice (it will contain numbers from 0 to 999)
	fmt.Println("Result slice:", result)
}

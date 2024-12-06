// Generics allow writing functions and types that work with any data type
// while ensuring type safety.
//
// This improves reusability.

package main

import (
	"fmt"
)

type Num interface {
	int | float64 | uint | string
}

// Add is a generic function that works for types satisfying the Num constraint
func Add[T Num](a, b T) T {
	return a + b
}

func main() {
	// Adding integers
	intResult := Add(1, 2)
	fmt.Println("Integer addition:", intResult)

	// Concatenating strings
	strResult := Add("He", "llo")
	fmt.Println("String concatenation:", strResult)

	// Adding floats
	floatResult := Add(1.5, 2.3)
	fmt.Println("Float addition:", floatResult)

	// This won't compile: unsupported type
	//Add(1, true)

	results := Multiplier([]float64{1.2, 3.4}, func(n float64) float64 {
		return n * 2
	})
	fmt.Println(results)
}

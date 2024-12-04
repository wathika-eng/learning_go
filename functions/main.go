// functions as values
package main

import "fmt"

// custom type
type transformFn func(int) int

func main() {
	// slice of numbers
	numbers := []int{1, 2}
	tripple := createTransformer(3)
	fourth := createTransformer(4)
	// don't execute the function just use name not double()
	doubled := transformNumbers(numbers, func(number int) int { // anonymous func
		return number * 2
	})
	trippled := transformNumbers(numbers, tripple)

	fourthed := transformNumbers(numbers, fourth)
	fmt.Println(doubled)
	fmt.Println(trippled)
	fmt.Println(fourthed)
	// fmt.Println(getTransformerFunction)
	fmt.Println(factorial(5))
	fmt.Println(addUp(1, 2, 3))
	// how to pass arrays to a fn
	num := []int{1, 2}
	fmt.Println(addUp(num...))
}

// function to double/tripple numbers, takes in a slice of numbers and func param
func transformNumbers(numbers []int, transform transformFn) []int {
	// new slice of numbers to append the double
	dNumbers := []int{}
	for _, v := range numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}

// double ints
// func double(number int) int {
// 	return number * 2
// }
// func tripple(number int) int {
// 	return number * 3
// }

// return functions as return values
// func getTransformerFunction(numbers []int) transformFn {
// 	if numbers[0] == 1 {
// 		return double
// 	} else {
// 		return tripple
// 	}
// }

// returns functions with input as int and output as int
//
// function to produce other functions
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

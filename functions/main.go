package main

import (
	"errors"
	"fmt"
)

func main() {
	// declare a value which you will pass to the value
	a := 8
	b := -0.8
	results, remainder, err := printMe(a, int(b)) //err is always returned nil or with a value
	if err != nil {                               //if error has a value other than nil
		fmt.Printf("Error found, %v\n", err)
	} else { // else proceed
		fmt.Printf("Results are %v\t %v\n", results, remainder)
	}
}

/*
func name(params type) return type if there is {
you can also return multiple values (int, int)
errors cause panic, default value is nil but must be returned
}
*/
func printMe(a, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("can't divide by 0")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}

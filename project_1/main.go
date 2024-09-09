package main

import (
	"fmt"
	"time"
)

func main() {
	// declare birthyear var of type int
	var birthYear int
	fmt.Printf("Input your year of birth: ")
	// ensure input is of type int
	fmt.Scanf("%d", &birthYear)
	// ensure birth year is logically correct
	if birthYear < 1900 || birthYear > 2024 {
		fmt.Println("Input a correct year")
		// if not, show error and exit
		return
	}
	// now calculate the age taking time stdlib found in Go packages
	age := time.Now().Year() - birthYear
	fmt.Printf("Your age is: %d\n", age)
}

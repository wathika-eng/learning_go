package main

import "fmt"

func main() {
	age := 32
	fmt.Printf("Memory address of age (in main): %v\n", &age)
	agePointer := &age
	fmt.Printf("Dereferenced value of agePointer: %v\n", *agePointer)
	fmt.Printf("Memory address of agePointer (in main): %v\n", &agePointer)

	newAge := adultYears(&age)
	fmt.Printf("Returned value: %v\n", newAge)
}

// this age will be a copy, hence a different place in memory
// we can get a pointer to the age instead
func adultYears(age *int) int {
	// dereference the pointer and print the memory address of the original age variable
	fmt.Printf("Memory address of age (inside adultYears): %v\n", age)
	return *age - 18
	// we can avoid returning, and just overwrite the initial value
	// *age = *age -18
}

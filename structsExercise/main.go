package main

import "fmt"

const PI = 3.14

type Circle struct {
	Radius float64
}

// Calculate the circumference of the circle
func (c Circle) Circumference() float64 {
	return 2 * PI * c.Radius
}

// Calculate the area of the circle
func (c Circle) Area() float64 {
	return PI * (c.Radius * c.Radius)
}

func main() {
	// Create a circle with a specific radius
	ourCircle := Circle{Radius: 4}

	// Print the calculated values
	fmt.Printf("Circumference: %.2f\n", ourCircle.Circumference())
	fmt.Printf("Area: %.2f\n", ourCircle.Area())
}

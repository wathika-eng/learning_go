package main

import "fmt"

func main() {
	var amount, tip float64
	fmt.Printf("Enter the amount in Ksh: ")
	fmt.Scanf("%f", &amount)
	fmt.Printf("Enter tip for waiter: ")
	fmt.Scanf("%f", &tip)
	if amount < 0 || tip < 0 {
		fmt.Println("Enter correct amount of money!")
		return
	}
	fmt.Printf("Amount is: Ksh.%.2f\t Tip is: Ksh.%.2f\n", amount, tip)
	tipPercent := amount * (tip / 100)
	bill := amount + tipPercent
	fmt.Printf("Total bill: Ksh.%.2f\n", bill)
}

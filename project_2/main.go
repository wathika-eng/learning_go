package main

import "fmt"

func main() {
	var option int
	var weight, height float64

	fmt.Printf("Hey, input your weight, option 1(in Pounds) or option 2(in kgs): ")
	fmt.Scanf("%d", &option)
	if option == 1 {
		fmt.Printf("Input weight in Pounds: ")
		fmt.Scanf("%f", &weight)
		kg := weight * 0.453592
		weight = kg
		fmt.Printf("Your weight in Kilograms: %.1f\n", weight)
	} else if option == 2 {
		fmt.Printf("Input weight in Kgs: ")
		fmt.Scanf("%f", &weight)
		fmt.Printf("Your weight in Kgs: %.1f\n", weight)
	} else {
		fmt.Println("Input correct option")
		return
	}
	fmt.Printf("Input height in metres: ")
	fmt.Scanf("%f", &height)
	BMI := weight / (height * height)
	fmt.Printf("Your BMI is %.1f\n", BMI)
	if BMI >= 30 {
		fmt.Println("You are obese😧")
	} else if BMI >= 25 {
		fmt.Println("You are overweight😔")
	} else if BMI >= 18.5 {
		fmt.Println("Normal lifestyle 🙂")
	} else {
		fmt.Println("You are underweight😟")
	}
}

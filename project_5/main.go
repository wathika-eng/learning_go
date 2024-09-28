package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const rounds int = 3

func main() {

	var userChoice string
	for i := 0; i < rounds; i++ {
		compChoiceNum := rand.Intn(3)
		fmt.Printf("Round %d - Enter your choice (Rock, Paper, Scissors): \n", i+1)
		fmt.Scanln(&userChoice)
		userChoice = strings.Title((strings.ToLower(userChoice)))

		var compChoice string
		switch compChoiceNum {
		case 0:
			compChoice = "Rock"
		case 1:
			compChoice = "Scissors"
		case 2:
			compChoice = "Paper"
		}

		fmt.Printf("User chose: %v \t Computer chose: %v\n", userChoice, compChoice)

		switch {
		case userChoice == compChoice:
			fmt.Println("It's a tie!")
		case userChoice == "Rock" && compChoice == "Scissors",
			userChoice == "Paper" && compChoice == "Rock",
			userChoice == "Scissors" && compChoice == "Paper":
			fmt.Println("User won!")
		case compChoice == "Rock" && userChoice == "Scissors",
			compChoice == "Paper" && userChoice == "Rock",
			compChoice == "Scissors" && userChoice == "Paper":
			fmt.Println("Computer won!")
		default:
			fmt.Println("Invalid user choice. Please enter Rock, Paper, or Scissors.")
			return
		}
	}
	fmt.Println("Game over!")
}

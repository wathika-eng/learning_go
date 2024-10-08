package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var withdrawAmount, depositAmount float64
	bankBalance := 300.20

	pl("Welcome to Go Bank!")
	for {
		pl("What do you want to do?: ")
		fmt.Println("1. Check balance")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Deposit")
		fmt.Println("4. Exit")
		fmt.Print("Your choice? ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		userChoice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		switch userChoice {
		case 1:
			fmt.Printf("Your balance is Ksh. %.2f\n", bankBalance)

		case 2:
			fmt.Print("Enter amount to withdraw: ")
			withdrawInput, _ := reader.ReadString('\n')
			withdrawInput = strings.TrimSpace(withdrawInput)
			withdrawAmount, err = strconv.ParseFloat(withdrawInput, 64)

			if err != nil || withdrawAmount <= 0 {
				pl("Invalid amount! Can't withdraw a negative or non-numeric amount.")
				continue
			}

			if withdrawAmount > bankBalance {
				pl("Not enough balance!")
			} else {
				bankBalance -= withdrawAmount
				fmt.Printf("Withdrew Ksh. %.2f, balance is Ksh. %.2f\n", withdrawAmount, bankBalance)
			}

		case 3:
			fmt.Print("Enter amount to deposit: ")
			depositInput, _ := reader.ReadString('\n')
			depositInput = strings.TrimSpace(depositInput)
			depositAmount, err = strconv.ParseFloat(depositInput, 64)

			if err != nil || depositAmount <= 0 {
				pl("Invalid amount! Can't deposit a negative or non-numeric amount.")
				continue
			}

			bankBalance += depositAmount
			fmt.Printf("Deposited Ksh. %.2f, balance is Ksh. %.2f\n", depositAmount, bankBalance)

		case 4:
			pl("Exiting...")
			return

		default:
			pl("Invalid choice! Please select a valid option.")
		}
	}
}

func pl(input string) {
	fmt.Println(input)
}

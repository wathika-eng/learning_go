// banking system
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const accountBalance = "balance.txt"

// write the balance to a file
func writeBalanceToFile(balance float64) {
	// file name will be created if not present, bytes will be written
	formattedBalance := fmt.Sprint(balance)
	// convert string to byte and declare permission
	os.WriteFile(accountBalance, []byte(formattedBalance), 0644)
}

// read balance from a file
func readBalanceFromFile() (float64, error) {
	// convert from byte to string then float
	// error is always a value in Go
	data, err := os.ReadFile(accountBalance)
	//we have a error if not nil
	if err != nil {
		// return 2 values
		return 0, errors.New("failed to read the file")
	}
	balanceString := string(data)
	balanceFloat, _ := strconv.ParseFloat(balanceString, 64)
	// return value and nil (to show no error of if available)
	return balanceFloat, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var withdrawAmount, depositAmount float64
	//receive 2 values
	bankBalance, err := readBalanceFromFile()
	if err != nil {
		// you can use os.Exit || panic() || return || log.Fatal()
		// panic not preferred for simple errors
		log.Fatal(err)
	}
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
				writeBalanceToFile(bankBalance)
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
			writeBalanceToFile(bankBalance)
			fmt.Printf("Deposited Ksh. %.2f, balance is Ksh. %.2f\n", depositAmount, bankBalance)

		case 4:
			pl("Exiting...")
			return

		default:
			pl("Invalid choice! Please select a valid option.")
		}
	}
}

// custom function to print with a new char but shortened
func pl(input string) {
	fmt.Println(input)
}

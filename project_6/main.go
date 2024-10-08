package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const bankBalance float64 = 400.2

func main() {
	reader := bufio.NewReader(os.Stdin)
	var userChoice int
	var withdrawAmount float64
	var err error
	pl("Welcome to Go Bank!")
	for {
		pl("What do you want to do?: ")
		fmt.Printf("1.Check balance\n2.Withdraw\n3.Deposit\n4.Exit\n")
		fmt.Printf("Your choice? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		userChoice, _ = strconv.Atoi(input)
		if err == nil {
			panic("Error encountered")
		}
		switch userChoice {
		case 1:
			fmt.Printf("Your balance is Ksh.%v\n", bankBalance)
			// return
		case 2:
			fmt.Printf("Enter amount to withdraw: ")
			fmt.Scanf("%d", &withdrawAmount)
			if withdrawAmount > bankBalance {
				pl("Not enough amount!")

			}
			if withdrawAmount <= 0 {
				pl("Can't withdraw a negative amount")
			}
			balance := bankBalance - withdrawAmount
			fmt.Printf("Withdrew Ksh.%v, balance is %v\n", withdrawAmount, balance)
		case 4:
			pl("Exiting...")
			return
		}
	}
}

func pl(input string) {
	fmt.Println(input)
}

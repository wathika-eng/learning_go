package balance

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalance = "balance.txt"

// write the balance to a file
func WriteBalanceToFile(balance float64) {
	// file name will be created if not present, bytes will be written
	formattedBalance := fmt.Sprint(balance)
	// convert string to byte and declare permission
	os.WriteFile(accountBalance, []byte(formattedBalance), 0644)
}

// read balance from a file
func ReadBalanceFromFile() (float64, error) {
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
func Balance() {
	currentBalance, err := ReadBalanceFromFile()
	if err != nil {
		fmt.Println("No existing balance file found, setting to default balance of Ksh. 0.00.")
		currentBalance = 0.00
		WriteBalanceToFile(currentBalance)
	}

	fmt.Printf("The current balance is: %.2f\n", currentBalance)
}

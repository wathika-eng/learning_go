package main

// comma, err syntax
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// just type package name without importing first
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rate me!")
	/*
		comma, ok syntax we don't have a try catch, error are variables e.g true or false
		hold it in an err, else in _ if you don't get feedback
		when dealing with stdin, we shall obviously have errors so catch it with _
		err will be an issue coz maybe you will not use it (unused vars
		make code not to compile) red error
	*/
	user_input, _ := reader.ReadString('\n') //read until user enters a new line
	fmt.Println("Thanks for rating,", user_input)

}

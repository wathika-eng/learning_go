package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
if string is returned, print it, else print the error
*/
func main() {
	var input string
	fmt.Println("Enter your name:")
	fmt.Scanln(&input)

	boy, err := name(input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(boy)
	}
}

/*
this function returns a string and an error incase
return 2 values either the string or error
*/
func name(you string) (string, error) {
	if you != "kamaa" {
		return "", errors.New("You are not Kamaa")
	} else {
		return strings.ToUpper(you), nil
	}
}

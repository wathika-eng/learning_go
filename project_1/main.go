package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// declare input var of type int
	// var input int
	reader := bufio.NewReader(os.Stdin)

	// Read input from the user

	fmt.Printf("Input your year of birth: ")
	// ensure input is of type int
	input, _ := reader.ReadString('\n')
	// Trim the newline character that gets included in the input
	input = strings.TrimSpace(input)

	age, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	} else {

		// fmt.Scanf("%d", &input)
		// ensure birth year is logically correct
		if age < 1900 || age > 2024 {
			fmt.Println("Input a correct year")
			// if not, show error and exit
			return
		}
		// now calculate the age taking time stdlib found in Go packages
		old := int64(time.Now().Year()) - age
		fmt.Printf("Your age is: %d\n", old)
	}
}

package main

import (
	"fmt"
	"strings"
)

/*
we have to rep strings as binary, but some chars aren't in ascii, so no more bits
if we increase bits we waste memory, that is UTF-32
so instead we have UTF-8, anything in ascii uses 1byte
runes are int 32, can be '', cast strings to rune
strings are immutable
instead of concatenating strings, use strings.Builder which is faster
*/

func main() {
	myString := []rune("Hey")
	for char, index := range myString {
		fmt.Printf("%v\t %v\n", index, char)
	}
	var stringSlice = []string{"h", "e", "y"}
	var stringBuilder strings.Builder
	for i := range stringSlice {
		stringBuilder.WriteString(stringSlice[i])
	}
	// appends vars
	var catStr = stringBuilder.String()
	fmt.Printf("%v", catStr)
}

package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

/*
for vals outside main, give the type
consts must have type always
consts won't throw error if unused
*/
// const sample := "Won't work"
var sampleAno string = "Works"

// anotherOne := "won't work"
const works string = "Works"

func main() { // opening curly brace must be on same line
	// the larger the int size, the more memory used
	// if you add more than limit it will cause buffer overflow
	// at compile time error is thrown if excess numbers
	var a int16 = 32767
	// at runtime no error but will lead to wrong output
	a = a + 2
	// unsigned int only store positives
	// you can declare int(127 limit) only but not float only
	var b int
	// var c float32
	fmt.Printf("Type of B: %T\t Size: %d", b, unsafe.Sizeof(b))
	fmt.Println(a)
	/*
		to get len of string import unicode/utf8
		if you use inbuilt len fn, it will get number of bytes
		any weird chars will lead to wrong results
		Go uses utf-8 for chars outside ascii
		for rune use single strings ''
	*/
	var myString string
	myString = "Hey😂"
	fmt.Printf("Length of string: %v\n", len(myString))
	fmt.Printf("Exact length: %v\n", utf8.RuneCountInString(myString))
}

/*
You can't perform operations with mixed types
int divisions result in ints being returned after rounding off
default vals for int, float, rune is 0, for strings is "", for bool is false
*/

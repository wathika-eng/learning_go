package main

import "fmt"

type str string

func (e str) customErr() {
	fmt.Println(e)
}

func ErrFn() {
	var err str = "Abort now"
	err.customErr()
}

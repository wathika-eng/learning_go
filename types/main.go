package main

import (
	"fmt"
	"math/rand"
)

type integer = int32

// var x integer = 10

func main() {
	min := 1
	max := 100
	age := rand.Intn(max-min) + min
	ag, _ := ageNum(age)
	fmt.Println(ag)
}

func ageNum(age int) (string, error) {
	if age > 50 {
		panic("Too old")

	}
	return "You are young", nil
}

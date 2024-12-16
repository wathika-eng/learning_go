package main

import (
	"fmt"
	"reflect"
	"time"
)

type order struct {
	name  string
	price int32
	date  string
}

func createOrder(q interface{}) {
	// type
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	// data structure
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)
}

func main() {
	i := 10
	// with this, we get the type at compile time, not runtime
	fmt.Printf("%d %T\n", i, i)
	ordered := order{
		"Rice",
		100,
		time.Now().Local().Local().String(),
	}
	createOrder(ordered)
	fmt.Printf("struct: %v\n", ordered)
}

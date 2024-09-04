package main

import (
	"fmt"
	"unsafe"
)

/*
	ARRAYS

fast, you declare size, hence static, can check cap()
fixed length, same type specified
indexable [0], continious mem ()

	Slices

dynamic, more mem, a bit slow
are just  arrays under the hood
*/
func main() {
	intArr := [...]int64{1, 2, 3}
	var sample = [3]int{1, 2, 3}
	fmt.Printf("size of array %d\n", unsafe.Sizeof(intArr))
	fmt.Println(sample, intArr)
	// use make type, length, capacity
	// setting the capacity is efficient for speed
	var newSlice []int = make([]int, 3, 4)

	println(newSlice)

}

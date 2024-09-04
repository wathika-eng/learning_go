package main

import (
	"fmt"
	"unsafe"
)

/*
fixed length, same type specified
indexable [0], continious mem ()
*/
func main() {

	intArr := [...]int64{1, 2, 3}
	var sample = [3]int{1, 2, 3}
	fmt.Printf("size of array %d\n", unsafe.Sizeof(intArr))
	fmt.Println(sample, intArr)
}

package main

import "fmt"

func main() {
	productNames := []string{"Apple", "Banana", "Carrot"}
	pl(productNames[1:]) // Banana, Carrot
	pl(productNames[:1]) //apple
	productNames[0] = "Cherry"
	pl(productNames) // array is overwritten
	newProductsList := append(productNames, "Kiwi")
	fmt.Printf("Len: %v, Cap: %v\n", len(newProductsList), cap(newProductsList)) // 4, capacity becomes x2 or x1.5 of len
	pl(newProductsList)                                                          // kiwi comes last
	pl(newProductsList[len(newProductsList)-1])                                  // we have 4 elements but we need last index and we start counting with 0
	fmt.Printf("Type is %T", productNames)
}

//short fmt.Println(s any)
func pl(s any) {
	fmt.Println(s)
}

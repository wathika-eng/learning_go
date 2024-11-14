package main

import "fmt"

func main() {
	productNames := []string{"Apple", "Banana", "Carrot"}
	pl(productNames[1:]) // Banana, Carrot
	pl(productNames[:1]) //apple
	productNames[0] = "Cherry"
	pl(productNames) // array is overwritten
	newProductsList := append(productNames, "Kiwi")
	fmt.Printf("Len: %v, Cap: %v", len(productNames), cap(newProductsList)) // 3, capacity becomes x2 of len
}

//short fmt.Println(s any)
func pl(s any) {
	fmt.Println(s)
}

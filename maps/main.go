package main

import "fmt"

/*
key value pairs
to loop through use range keyword, no order is preserver
*/

func main() {
	var myMap = map[string]int{"Me": 21, "another": 22}
	fmt.Println(myMap)
	for name := range myMap {
		fmt.Printf("Name: %v \n", name)
	}
	newMap := make(map[string]int)
	newMap["Kamaa"] = 10
	fmt.Println(newMap)
}

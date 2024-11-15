package Hobbies

import (
	"fmt"
	"unsafe"
)

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func LearningMaps() {
	websites := map[string]string{
		"Google": "https://www.google.com",
	}
	websites["Twitter"] = "https://x.com"
	// fmt.Println(websites)
	// delete(websites, "Twitter")
	// fmt.Println(websites)
	fmt.Printf("Size %v,Type %T\n", unsafe.Sizeof(websites), websites)
	for key, value := range websites {
		fmt.Printf("%v: %v\n", key, value)
	}
	newPrices := make(floatMap, 3)
	newPrices["Shoes"] = 20.1
	newPrices.output()
}

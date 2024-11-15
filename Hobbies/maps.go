package Hobbies

import (
	"fmt"
	"unsafe"
)

func LearningMaps() {
	websites := map[string]string{
		"Google": "https://www.google.com",
	}
	websites["Twitter"] = "https://x.com"
	fmt.Println(websites)
	delete(websites, "Twitter")
	fmt.Println(websites)
	fmt.Printf("Size %v,Type %T\n", unsafe.Sizeof(websites), websites)
}

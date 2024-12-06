package main

// type Num interface {
// 	int | float64 | uint | string
// }

func Multiplier[T Num](values []T, multi func(n T) T) []T {
	newArray := make([]T, 0, 3)
	for _, v := range values {
		valuesArr := multi(v)
		newArray = append(newArray, valuesArr)
	}
	return newArray
}

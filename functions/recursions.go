package main

// all recursive fns must have an exit statement
//
// calling yourself
func factorial(num int) int {
	// since the fn calls itself, when num is 0, stop
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
	// result := 1
	// for i := 1; i <= num; i++ {
	// 	result = result * i
	// }
	// return result
}

// variadic fn has ... to allow more than 1 parameter value
func addUp(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

package main

// 1. update state
// 2. Reduce memory instead of copying whole struct

func main() {
	// a, b := 1, 2
	// fmt.Println(a, b)
	// // & returns pointer address of value
	// p := &a
	// fmt.Println(p)
	// //returns value of what the pointer is pointing to
	// fmt.Println(*p)
	// // modifying *p changes the initial value a
	// // will still be in same memory, reduces copying
	// *p = 3
	// fmt.Println(a)
	// SquareAdd(2)
	user := User{
		"jowienew@gmail.com",
		"Kamaa",
		21,
	}
	user.SetUser(user)
	// fmt.Println(user.Email())
	user.UpdateEmail("Kamaa@gmail.com")
	user.JsonifyUser()
	//fmt.Println(user.UserEmail())
	// fmt.Println(user.JsonifyUser())
}

package main

import (
	"fmt"
	"log"

	"example.com/structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	// outPutUserData(appUser)
	//var appUser **User
	appUser, err := user.New(userfirstName, userlastName, userbirthDate)
	admin, err2 := user.NewAdmin("examplegmail.com", "123j")
	if err != nil || err2 != nil {
		log.Fatal(err)
		log.Fatal(err2)
	} else {

		appUser.OutPutUserData()
		appUser.ClearUsername()
		// appUser.OutPutUserData()
	}
	admin.OutPutUserData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	// scan doesn't allow empty value
	// scanln tells program that you are done when you enter a new line
	fmt.Scanln(&value)
	return value
}

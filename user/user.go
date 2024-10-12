package user

import (
	"errors"
	"fmt"
	"time"
)

// define your own type which shall be a struct or blueprint
// methods - functions attached to structs
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// sample struct showing inheritance
type Admin struct {
	email    string
	password string
	// to inherit values from another struct, we embed
	User
	// use anonymous
	// sometime avoid declaring name so as to easily calling struct with name
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" {
		return nil, errors.New("email is required")
	} else {
		return &Admin{
			email:    email,
			password: password,
			User: User{
				firstName: "Admin",
				lastName:  "kamaa",
				birthDate: "09/10/2002",
				createdAt: time.Now(),
			},
		}, nil
	}
}

// creating a constructor function, which will create our struct
func New(userfirstName, userlastName, userbirthDate string) (*User, error) {
	//check for errors before creating a new user

	if userfirstName == "" || userlastName == "" || userbirthDate == "" {
		return nil, errors.New("should not be nil")
	}
	// create a variable based on your struct
	// instead of returning value, return a pointer to avoid copying
	return &User{
		// you can remove the struct fields if you will assign values based on order
		firstName: userfirstName,
		lastName:  userlastName,
		birthDate: userbirthDate,
		createdAt: time.Now(),
	}, nil
}

// instead of passing values one by one, pass the struct type
// to make it a method, add a pair of () before fnname and specify struct, and get rid of param
// method receiver argument, u is the instance
func (u User) OutPutUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

// methods and functions receive copies, which might waste memory, instead use pointers
// sometimes, you might edit copies
func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

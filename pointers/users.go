package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Email    string `json:email`
	Username string `json:username`
	Age      int    `json:age`
}

func (u *User) SetUser(data User) *User {
	//fmt.Printf("%v\n", unsafe.Sizeof(u))
	u = &data
	return u
}

func (u *User) UpdateEmail(email string) {
	//fmt.Printf("%v\n", unsafe.Sizeof(u))
	u.Email = email
}

func (u User) JsonifyUser() {
	data, _ := json.MarshalIndent(u, "", " ")
	fmt.Println(string(data))
}

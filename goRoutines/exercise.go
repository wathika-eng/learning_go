package main

import "time"

func getUserByFriends(id int) []string {
	time.Sleep(time.Second * 3)
	return []string{
		"Kamaa",
		"John",
	}
}

func getUserChats(id int) []string {
	time.Sleep(time.Second * 3)
	chats := make(map[int][]string)
	chats[1] = []string{"Hey", "Come"}
	chats[2] = []string{"Niko Poa", "Nakuja"}
	// access value in a map based on id
	if chat, ok := chats[id]; ok {
		return chat
	}
	// return empty slice if not present
	return []string{}
}

func getUserByName(name string) string {
	return "Kamaa"
	// "John",
}

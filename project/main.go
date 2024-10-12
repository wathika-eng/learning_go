// A note taking app which stores them in JSON format
package main

import (
	"errors"
	"fmt"
	"log"
	"noteApp/note"
)

func main() {
	title, content, _ := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		log.Fatal(err)
	} else {
		userNote.DisplayNote()
	}
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var input string
	fmt.Scanln(&input)
	if input == "" {
		return "", errors.New("content cannot be blank")
	}
	return input, nil
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title: ")
	content, err2 := getUserInput("Your content: ")
	if err != nil || err2 != nil {
		log.Fatal(err)
		log.Fatal(err2)
	}
	return title, content, nil
}

// A note taking app which stores them in JSON format
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"noteApp/note"
	"os"
	"strings"
)

func main() {
	title, content, _ := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		// exit and display error
		log.Fatal(err)
	}
	// userNote.DisplayNote()
	data, err := userNote.Save()
	if err != nil {
		log.Print("Saving the file failed")
	}
	fmt.Printf("Json file data:\n%v", data)
}

// gets user input with a reader and passes it to the note data function
func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var err error
	var input string
	reader := bufio.NewReader(os.Stdin)
	// reader will capture blank error automatically
	input, err = reader.ReadString('\n')
	if err != nil {
		return "", errors.New("content cannot be blank")
	}
	formattedInput := strings.TrimSpace(input)
	return formattedInput, nil
}

// gets the data from user input on the terminal, returns title, content and error if any
func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Title: ")
	content, err2 := getUserInput("Your content: ")
	if err != nil || err2 != nil {
		log.Fatal(err)
		log.Fatal(err2)
	}
	return title, content, nil
}

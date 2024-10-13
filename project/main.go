// A note taking app which stores them in JSON format
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"noteApp/todo"
	"os"
	"strings"
)

// interface which dictates a set of methods other structs must use or a contract
// if it has one method, name it method name + "er"
type saver interface {
	// a method Save, it returns string or lserror
	Save() (string, error)
}

func main() {
	todoText, err := getUserInput("Todo text: ")
	todoData, err := todo.New(todoText)
	if err != nil {
		// exit and display error
		log.Fatal(err)
	}
	// userNote.DisplayNote()
	saveDate(todoData)
}

// the data var is of type saver
func saveDate(data saver) {
	// you can call the methods defined
	saveData, err := data.Save()
	if err != nil {
		log.Print("Saving the file failed")
	}
	fmt.Printf("Json file data:\n%v", saveData)
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
// func getNoteData() (string, error) {
// 	content, err := getUserInput("Your content: ")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return content, nil
// }

package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// note struct which contains title, content and auto generated time
type Note struct {
	title   string
	content string
	// instead of time.Now() use a string
	createdAt string
}

// save each note to a file in the system
func (n Note) Save() error {
	// format the file well to allow constitency
	fileName := strings.ReplaceAll(n.title, " ", "_")
	fileName = strings.ToLower(fileName)
	// format to json, will aready be byte
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}
	// filename, format, permissions
	savedFile := os.WriteFile(fileName, json, 0644)
	return savedFile
}

// creates a new note \n
// using such pointers creates an extra overhead
func New(noteTitle, noteContent string) (*Note, error) {
	if noteTitle == "" || noteContent == "" {
		return &Note{}, errors.New("cannot be empty")
	}
	now := time.Now()
	return &Note{

		title:   noteTitle,
		content: noteContent,
		// changed the format of time displayed
		createdAt: now.Format(time.UnixDate),
	}, nil

}

// a method for struct Note which displays notes content, instead of having structs fields as public
func (n Note) DisplayNote() {
	fmt.Println(n.title, n.content, n.createdAt)
}

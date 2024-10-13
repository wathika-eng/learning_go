package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

// note struct which contains title, Content and auto generated time
type Note struct {
	// must convert to uppercase when saving/json encoding
	// backticks add struct tags/metadata
	Title   string `json:"title"`
	Content string `json:"content"`
	// instead of time.Now() use a string
	CreatedAt string `json:"created_at"`
}

// save each note to a file in the system
func (note Note) Save() (string, error) {
	// format the filename well to allow constitency
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	// format to json, will already be byte
	// we can indent to 4 spaces
	jsonData, err := json.MarshalIndent(note, "", "    ")
	if err != nil {
		// return error and empty string
		return "", err
	}
	// filename, format, permissions
	savedFile := os.WriteFile(fileName, jsonData, 0644)
	// read the filename saved, error is unused
	readFile, _ := os.ReadFile(fileName)
	// since returned data is bytes, we convert to string format
	data := string(readFile)
	// return file name and data
	return data, savedFile
}

// creates a new note \n
// using such pointers creates an extra overhead
func New(noteTitle, noteContent string) (*Note, error) {
	if noteTitle == "" || noteContent == "" {
		//disallows saving empty files, very important
		return &Note{}, errors.New("all fields are required")
	}
	now := time.Now()
	return &Note{

		Title:   noteTitle,
		Content: noteContent,
		// changed the format of time displayed
		CreatedAt: now.Format(time.UnixDate),
	}, nil

}

// a method for struct Note which displays notes Content, instead of having structs fields as public
// func (n Note) DisplayNote() {
// 	fmt.Println(n.Title, n.Content)
// }

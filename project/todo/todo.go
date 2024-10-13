package todo

import (
	"encoding/json"
	"errors"
	"os"
)

type Note struct {
	Content string `json:"content"`
}

// save each note to a file in the system
func (note Note) Save() (string, error) {
	fileName := "todo.json"
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
func New(noteContent string) (*Note, error) {
	if noteContent == "" {
		//disallows saving empty files, very important
		return &Note{}, errors.New("all fields are required")
	}
	return &Note{
		Content: noteContent,
	}, nil

}

// a method for struct Note which displays notes Content, instead of having structs fields as public
// func (n Note) DisplayNote() {
// 	fmt.Println(n.Title, n.Content)
// }

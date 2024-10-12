package note

import (
	"errors"
	"fmt"
	"time"
)

// note struct which contains title, content and auto generated time
type Note struct {
	title   string
	content string
	// instead of time.Now() use a string
	createdAt string
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

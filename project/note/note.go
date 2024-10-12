package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title   string
	content string
	// instead of time.Now() use a string
	createdAt string
}

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

func (n Note) DisplayNote() {
	fmt.Println(n.title, n.content, n.createdAt)
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	user := User{
		firstName: "John",
		lastName:  "Doe",
		birthDate: "01/01/1990",
		createdAt: time.Now(),
	}

	if user.firstName != "John" {
		t.Errorf("Expected firstName to be 'John', got '%s'", user.firstName)
	}
	if user.lastName != "Doe" {
		t.Errorf("Expected lastName to be 'Doe', got '%s'", user.lastName)
	}
	if user.birthDate != "01/01/1990" {
		t.Errorf("Expected birthDate to be '01/01/1990', got '%s'", user.birthDate)
	}
	if user.createdAt.IsZero() {
		t.Error("Expected createdAt to be set, but it's zero")
	}
}

func TestGetUserData(t *testing.T) {
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	tests := []struct {
		name        string
		input       string
		promptText  string
		expectedOut string
	}{
		{"First Name", "Alice\n", "Please enter your first name: ", "Alice"},
		{"Last Name", "Smith\n", "Please enter your last name: ", "Smith"},
		{"Birthdate", "05/15/1985\n", "Please enter your birthdate (MM/DD/YYYY): ", "05/15/1985"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdin = r
			go func() {
				w.Write([]byte(tt.input))
				w.Close()
			}()

			result := getUserData(tt.promptText)
			if result != tt.expectedOut {
				t.Errorf("Expected '%s', got '%s'", tt.expectedOut, result)
			}
		})
	}
}

func TestOutPutUserData(t *testing.T) {
	user := User{
		firstName: "Jane",
		lastName:  "Doe",
		birthDate: "02/14/1995",
		createdAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outPutUserData(user)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := fmt.Sprintf("Jane Doe 02/14/1995 %s\n", user.createdAt.String())
	if output != expected {
		t.Errorf("Expected output '%s', got '%s'", expected, output)
	}
}

func TestMain(t *testing.T) {
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	input := "John\nDoe\n01/01/1990\n"
	r, w, _ := os.Pipe()
	os.Stdin = r
	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	or, ow, _ := os.Pipe()
	os.Stdout = ow

	main()

	ow.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, or)
	output := buf.String()

	expectedPrompts := "Please enter your first name: Please enter your last name: Please enter your birthdate (MM/DD/YYYY): "
	if !strings.Contains(output, expectedPrompts) {
		t.Errorf("Expected output to contain '%s', but it didn't", expectedPrompts)
	}

	expectedOutput := "John Doe 01/01/1990"
	if !strings.Contains(output, expectedOutput) {
		t.Errorf("Expected output to contain '%s', but it didn't", expectedOutput)
	}
}

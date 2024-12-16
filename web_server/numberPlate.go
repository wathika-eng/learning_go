package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generateNumberPlate generates a random number plate in the format "KAB 603A"
func generateNumberPlate() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	// Generate three random letters
	letterPart := "K" +
		string(letters[rand.Intn(len(letters))]) +
		string(letters[rand.Intn(len(letters))])

	// Generate a three-digit number between 100 and 999
	numberPart := rand.Intn(900) + 100

	// Generate one random letter
	lastLetter := string(letters[rand.Intn(len(letters))])

	return fmt.Sprintf("%s %03d%s", letterPart, numberPart, lastLetter)
}

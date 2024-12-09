// common utility functions
package utils

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const TIMEZONE int = 3 * 60 * 60
const TZ = TIMEZONE / 3600 // +3
const ROUNDS int = 10

// return time in GMT +3 format
func DBTime() string {
	return time.Now().In(time.FixedZone("GMT +3", TIMEZONE)).Format("2 Jan 2006 15.04") + " GMT+3"
}

// hash the password before saving to DB
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), ROUNDS)
	return string(bytes), err
}

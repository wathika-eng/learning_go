package utils

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

func DBTime() string {
	return time.Now().In(time.FixedZone("GMT+3", 3*60*60)).Format("2 Jan 2006 15.04") + " GMT+3"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

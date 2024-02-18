package utils

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateTimeNowJakarta() time.Time {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println(err)
	}
	now := time.Now().In(loc)
	return now
}

func GenerateBool(value int) bool {
	return value == 1
}

func GenerateInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

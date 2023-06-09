package helpers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
func VerifyPassword(userPassword, providePassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providePassword), []byte(userPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("Password is incorect")
		check = false
	}
	return check, msg
}

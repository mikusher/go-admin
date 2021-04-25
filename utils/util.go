package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashAndSaltPassword(password string) []byte {
	// generate password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return hash
}

func CheckPasswordHash(hashPass string, plainPwd []byte) bool {
	byteHash := []byte(hashPass)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

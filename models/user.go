package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
}

func (user *User) SetPassword(password []byte) {
	hashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user.Password = hashPassword
}

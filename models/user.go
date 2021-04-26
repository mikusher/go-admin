package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	RoleId    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleId"`
}

// set string password
func (user *User) SetStringPassword(password string) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = hashPassword
}

// SetHashPassword set hash password
func (user *User) SetHashPassword(password []byte) {
	hashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	user.Password = hashPassword
}

// SetPassword
func (user *User) SetPassword(password string) {
	user.Password = []byte(password)
}

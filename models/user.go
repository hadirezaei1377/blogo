package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username" gorm:"uniqueIndex"`
	Password  []byte    `json:"-"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Posts     []Post    `gorm:"foreignKey:AuthorID"`
	Comments  []Comment `gorm:"foreignKey:UserID"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	RoleID    uint      `json:"role_id"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Printf("Error hashing password for user %s: %s", user.Username, err.Error())
		return
	}
	user.Password = hashedPassword
	log.Printf("Password set for user %s", user.Username)
}

func (user *User) ComparePasswords(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		log.Printf("Error: Password comparison failed for user %s: %s", user.Username, err.Error())
	}
	return err
}

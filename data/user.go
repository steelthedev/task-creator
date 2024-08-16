package data

import (
	"fmt"

	"github.com/steelthedev/task-handler/utils"
)

type User struct {
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password" gorm:"column:password"`
}

type SafeUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u *User) ToSafeUser() *SafeUser {
	return &SafeUser{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	}
}

func (u *User) ValidateEmail() error {
	if !utils.EmailIsValid(u.Email) {
		return fmt.Errorf("invalid email")
	}
}

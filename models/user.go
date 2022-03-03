package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User is the user model
type User struct {
	ID           uint   `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	Username     string `json:"username,omitempty"`
	Email        string `json:"email,omitempty"`
	Picture      string `json:"picture,omitempty"`
	Password     string `json:"password,omitempty"`
	PasswordHash string `json:"-"`
	IsActive     bool   `json:"is_active,omitempty"`
	IsAdmin      bool   `json:"is_admin,omitempty"`
	Base
}

// HashPassword hashes the password
func (u *User) HashPassword() error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.PasswordHash = string(passwordHash)

	return nil
}

func (u *User) PasswordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	return err == nil
}

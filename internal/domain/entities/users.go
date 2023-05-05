package entities

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint   `json:"ID"`
	Email     string `json:"Email"`
	Name      string `json:"Name"`
	LastName  string `json:"LastName"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Projects  []Project      `json:"Projects"`
}

func (u *User) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

// compare hashedpassword with possible password in plain text
func (u *User) VerifyPassword(passwordToCompare string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwordToCompare))
	return err == nil, err
}

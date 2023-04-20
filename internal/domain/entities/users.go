package entities

import (
	"time"
)

type User struct {
	ID        string    `json:"ID"`
	Email     string    `json:"Email"`
	Name      string    `json:"Name"`
	LastName  string    `json:"Last Name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	// Projects  []Project      `json:"Projects"`
}

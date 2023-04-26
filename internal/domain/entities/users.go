package entities

import (
	"time"
)

type User struct {
	ID        uint   `json:"ID"`
	Email     string `json:"Email"`
	Name      string `json:"Name"`
	LastName  string `json:"LastName"`
	Password  string `json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// Projects  []Project      `json:"Projects"`
}

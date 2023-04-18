package domain

type User struct {
	Id       string    `json:"ID"`
	Name     string    `json:"Name"`
	LastName string    `json:"Last Name"`
	Projects []Project `json:"Projects"`
}

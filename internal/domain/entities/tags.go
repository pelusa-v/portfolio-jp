package entities

type Tag struct {
	ID          uint   `json:"ID"`
	Name        string `json:"Name" validate:"required"`
	Description string `json:"Description"`
}

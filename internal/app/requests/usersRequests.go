package requests

import "github.com/pelusa-v/portfolio-jp/internal/domain/entities"

type UserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}

func (userReq *UserRequest) MapToEntity() entities.User {
	return entities.User{
		Email:    userReq.Email,
		Name:     userReq.Name,
		LastName: userReq.LastName,
		Password: userReq.Password,
	}
}

type UserValidationRequest struct {
	Password string `json:"password"`
}

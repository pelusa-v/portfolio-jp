package models

import (
	"time"

	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"type:varchar(200);not null;uniqueIndex"`
	Name      string `gorm:"type:varchar(200);not null"`
	LastName  string `gorm:"type:varchar(200);not null"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Projects  []Project      `gorm:"type:"`
}

func MapEntityToModel(userEntity *entities.User) *User {
	return &User{
		ID:        userEntity.ID,
		Email:     userEntity.Email,
		Name:      userEntity.Name,
		LastName:  userEntity.LastName,
		Password:  userEntity.Password,
		CreatedAt: userEntity.CreatedAt,
		UpdatedAt: userEntity.UpdatedAt,
	}
}

func (userModel *User) MapModelToEntity() entities.User {
	return entities.User{
		ID:        userModel.ID,
		Email:     userModel.Email,
		Name:      userModel.Name,
		LastName:  userModel.LastName,
		Password:  userModel.Password,
		CreatedAt: userModel.CreatedAt,
		UpdatedAt: userModel.UpdatedAt,
	}
}

type ListUser []User

func (listUserModel *ListUser) MapModelToEntity() []entities.User {
	var userEntities []entities.User
	for _, user := range *listUserModel {
		userEntities = append(userEntities, user.MapModelToEntity())
	}
	return userEntities
}

package models

import (
	"time"

	"github.com/pelusa-v/portfolio-jp/internal/domain/entities"
)

type User struct {
	ID        string    `gorm:"primaryKey"`
	Email     string    `gorm:"type:varchar(200);not null;uniqueIndex"`
	Name      string    `gorm:"type:varchar(200);not null"`
	LastName  string    `gorm:"type:varchar(200);not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	// Projects  []Project      `gorm:"type:"`
}

func MapEntityToModel(userEntity *entities.User) User {
	return User{
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

package repository

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth"
)

type User struct {
	ID           uint
	Name         string
	Email        string `gorm:"unique"`
	Password     string
	Phone_Number string `gorm:"unique"`
	Role         string
	Address      string
	Home_Address string
	Team         string
	Status       string
	Gender       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ToCore(userModel User) auth.Core {
	return auth.Core{
		ID:       userModel.ID,
		Role:     userModel.Role,
		Email:    userModel.Email,
		Password: userModel.Password,
	}
}

func ToCoreLogin(userModel User) auth.Login {
	return auth.Login{
		ID:    userModel.ID,
		Role:  userModel.Role,
		Email: userModel.Email,
	}
}

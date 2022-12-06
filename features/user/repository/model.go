package repository

import (
	"time"

	_user "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Phone_Number string
	Role         string
	Address      string
	Home_Address string
	Team         string
	Status       string
	Gender       string
	Log          []Log
}

type Class struct {
	gorm.Model
	Name          string
	StartDate     time.Time
	GraduatedDate time.Time
	UserID        uint
}

type Log struct {
	gorm.Model
	UserID   uint
	MenteeID uint
	Notes    string
	Status   string
}

func fromCore(dataCore _user.Core) User {
	userGorm := User{
		Name:         dataCore.Name,
		Email:        dataCore.Email,
		Password:     dataCore.Password,
		Phone_Number: dataCore.Phone_Number,
		Role:         dataCore.Role,
		Address:      dataCore.Address,
		Home_Address: dataCore.Home_Address,
		Team:         dataCore.Team,
		Status:       dataCore.Status,
		Gender:       dataCore.Gender,
	}
	return userGorm
}

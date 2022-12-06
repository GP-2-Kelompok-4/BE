package repository

import (
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

func (dataModel *User) toCore() _user.Core {
	return _user.Core{
		ID:           dataModel.ID,
		Name:         dataModel.Name,
		Email:        dataModel.Email,
		Password:     dataModel.Password,
		Phone_Number: dataModel.Phone_Number,
		Role:         dataModel.Role,
		Address:      dataModel.Address,
		Home_Address: dataModel.Home_Address,
		Team:         dataModel.Team,
		Status:       dataModel.Status,
		Gender:       dataModel.Gender,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []User) []_user.Core {
	var dataCore []_user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

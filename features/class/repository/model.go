package repository

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
	UserID       uint
	User         User
}

type User struct {
	gorm.Model
	Name  string
	Class []Class
}

type ClassDetail struct {
	gorm.Model
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
	PIC          string
}

//from core to model

func fromCore(dataCore class.ClassCore) Class {
	userGorm := Class{
		Name:         dataCore.Name,
		StartDate:    dataCore.StartDate,
		GraduateDate: dataCore.GraduateDate,
		UserID:       dataCore.UserID,
	}
	return userGorm
}

//from model to core

// func (dataModel *User) toCoreUser() class.User {
// 	return class.User{
// 		ID:   dataModel.ID,
// 		Name: dataModel.Name,
// 	}
// }

func (dataModel *Class) toCore() class.ClassCore {
	return class.ClassCore{
		ID:           dataModel.ID,
		Name:         dataModel.Name,
		StartDate:    dataModel.StartDate,
		GraduateDate: dataModel.GraduateDate,
		UserID:       dataModel.UserID,
		// User:         dataModel.User.toCoreUser(),
		User: class.User{
			ID:   dataModel.User.ID,
			Name: dataModel.User.Name,
		},
	}
}

func toCoreList(dataModel []Class) []class.ClassCore {
	var dataCore []class.ClassCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

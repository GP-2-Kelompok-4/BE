package repository

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"

	"gorm.io/gorm"
)

type ClassModel struct {
	gorm.Model
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
	UserID       uint
	Mentee       []Mentee
}

type Mentee struct {
	gorm.Model
	Name                   string
	Nickname               string
	Email                  string
	Gender                 string
	PhoneNumber            string
	Address                string
	HomeAddress            string
	TelegramAccount        string
	DiscordAccount         string
	Status                 string
	EducationType          string
	EducationMajor         string
	Graduate               time.Time
	Institution            string
	EmergencyContact       string
	EmergencyContactStatus string
	ClassID                uint
}

type User struct {
	gorm.Model
	Name  string
	Class []ClassModel
}

//from core to model

func fromCore(dataCore class.ClassCore) ClassModel {
	userGorm := ClassModel{
		Name:         dataCore.Name,
		StartDate:    dataCore.StartDate,
		GraduateDate: dataCore.GraduateDate,
		UserID:       dataCore.UserID,
	}
	return userGorm
}

//from model to core

func (dataModel *ClassModel) toCore() class.ClassCore {
	return class.ClassCore{
		ID:           dataModel.ID,
		Name:         dataModel.Name,
		StartDate:    dataModel.StartDate,
		GraduateDate: dataModel.GraduateDate,
		UserID:       dataModel.UserID,
		CreatedAt:    dataModel.CreatedAt,
		UpdatedAt:    dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []ClassModel) []class.ClassCore {
	var dataCore []class.ClassCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

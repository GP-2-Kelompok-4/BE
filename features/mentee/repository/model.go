package repository

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee"
	"gorm.io/gorm"
)

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
	ClassId                uint
	Log                    []Log
	LogStruct              Log
	Class                  Class
}

type Class struct {
	gorm.Model
	Name   string
	Mentee []Mentee
}

type Log struct {
	gorm.Model
	UserId uint
	Status string
	Notes  string
	User   User
}

type User struct {
	gorm.Model
	Name string
	Log  []Log
}

//from core to model

func fromCore(dataCore mentee.MenteeCore) Mentee {
	menteeGorm := Mentee{
		Name:                   dataCore.Name,
		Nickname:               dataCore.Nickname,
		Email:                  dataCore.Email,
		Gender:                 dataCore.Gender,
		PhoneNumber:            dataCore.PhoneNumber,
		Address:                dataCore.Address,
		HomeAddress:            dataCore.HomeAddress,
		TelegramAccount:        dataCore.TelegramAccount,
		DiscordAccount:         dataCore.DiscordAccount,
		Status:                 dataCore.Status,
		EducationType:          dataCore.EducationType,
		EducationMajor:         dataCore.EducationMajor,
		Graduate:               dataCore.Graduate,
		Institution:            dataCore.Institution,
		EmergencyContact:       dataCore.EmergencyContact,
		EmergencyContactStatus: dataCore.EmergencyContactStatus,
		ClassId:                dataCore.ClassId,
	}
	return menteeGorm
}

func (dataModel *Mentee) toCore() mentee.MenteeCore {
	return mentee.MenteeCore{
		ID:                     dataModel.ID,
		Name:                   dataModel.Name,
		Nickname:               dataModel.Nickname,
		Email:                  dataModel.Email,
		Gender:                 dataModel.Gender,
		PhoneNumber:            dataModel.PhoneNumber,
		Address:                dataModel.Address,
		HomeAddress:            dataModel.HomeAddress,
		TelegramAccount:        dataModel.TelegramAccount,
		DiscordAccount:         dataModel.DiscordAccount,
		Status:                 dataModel.Status,
		EducationType:          dataModel.EducationType,
		EducationMajor:         dataModel.EducationMajor,
		Graduate:               dataModel.Graduate,
		Institution:            dataModel.Institution,
		EmergencyContact:       dataModel.EmergencyContact,
		EmergencyContactStatus: dataModel.EmergencyContactStatus,
		ClassId:                dataModel.ClassId,
		Class: mentee.Class{
			ID:   dataModel.Class.ID,
			Name: dataModel.Class.Name,
		},
		LogStruct: mentee.Log{
			ID:        dataModel.LogStruct.ID,
			UserId:    dataModel.LogStruct.UserId,
			Status:    dataModel.LogStruct.Notes,
			Notes:     dataModel.LogStruct.Notes,
			CreatedAt: dataModel.LogStruct.CreatedAt,
			User:      mentee.User{},
		},
	}
}

func toCoreList(dataModel []Mentee) []mentee.MenteeCore {
	var dataCore []mentee.MenteeCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

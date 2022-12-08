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
	TelegramAccount        string
	DiscordAccount         string
	ClassId                uint
	Address                string
	HomeAddress            string
	Status                 string
	EducationType          string
	EducationMajor         string
	Institution            string
	Graduate               time.Time
	EmergencyContact       string
	EmergencyContactName   string
	EmergencyContactStatus string
	Log                    []Log
	// LogStruct              Log
	Class Class
}

type Class struct {
	gorm.Model
	Name   string
	Mentee []Mentee
}

type Log struct {
	gorm.Model
	UserId   uint
	MenteeId uint
	Status   string
	Notes    string
	User     User
	Mentee   Mentee
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
		TelegramAccount:        dataCore.TelegramAccount,
		DiscordAccount:         dataCore.DiscordAccount,
		ClassId:                dataCore.ClassId,
		Address:                dataCore.Address,
		HomeAddress:            dataCore.HomeAddress,
		Status:                 dataCore.Status,
		EducationType:          dataCore.EducationType,
		EducationMajor:         dataCore.EducationMajor,
		Institution:            dataCore.Institution,
		Graduate:               dataCore.Graduate,
		EmergencyContact:       dataCore.EmergencyContact,
		EmergencyContactName:   dataCore.EmergencyContactName,
		EmergencyContactStatus: dataCore.EmergencyContactStatus,
	}
	return menteeGorm
}

// func (dataModel *Log) toCoreLog() mentee.Log {
// 	return mentee.Log{
// 		ID:     dataModel.ID,
// 		UserId: dataModel.UserId,
// 		Status: dataModel.Status,
// 		Notes:  dataModel.Notes,
// 		User: mentee.User{
// 			ID:   dataModel.User.ID,
// 			Name: dataModel.User.Name,
// 		},
// 	}
// }

func (dataModel *Mentee) toCore() mentee.MenteeCore {
	var arrLogs []mentee.Log
	for _, val := range dataModel.Log {
		arrLogs = append(arrLogs, mentee.Log{
			ID:     val.ID,
			Notes:  val.Notes,
			Status: val.Status,
		})
	}
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
		LogStruct: arrLogs,
		// LogStruct: mentee.Log{
		// 	ID:        dataModel.LogStruct.ID,
		// 	UserId:    dataModel.LogStruct.UserId,
		// 	Status:    dataModel.LogStruct.Notes,
		// 	Notes:     dataModel.LogStruct.Notes,
		// 	CreatedAt: dataModel.LogStruct.CreatedAt,
		// 	User:      dataModel.LogStruct.toCoreLog().User,
		// },
	}
}

func toCoreList(dataModel []Mentee) []mentee.MenteeCore {
	var dataCore []mentee.MenteeCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

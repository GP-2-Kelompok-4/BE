package repository

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	UserID   uint
	MenteeID uint
	Notes    string
	Status   string
}

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
	Log                    []Log
}

func fromCore(dataCore log.CoreLog) Log {
	logGorm := Log{
		UserID:   dataCore.UserID,
		MenteeID: dataCore.MenteeID,
		Notes:    dataCore.Notes,
		Status:   dataCore.Status,
	}
	return logGorm
}

func toCore(dataCore Log) log.CoreLog {
	logGorm := log.CoreLog{
		UserID:   dataCore.UserID,
		MenteeID: dataCore.MenteeID,
		Notes:    dataCore.Notes,
		Status:   dataCore.Status,
	}
	return logGorm
}

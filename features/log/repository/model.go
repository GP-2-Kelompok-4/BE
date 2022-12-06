package repository

import (
	"time"

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

package repository

import (
	"time"

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
	ClassID                uint
}

type Class struct {
	gorm.Model
	Name          string
	StartDate     time.Time
	GraduatedDate time.Time
	UserID        uint
	Mentee        []Mentee
}

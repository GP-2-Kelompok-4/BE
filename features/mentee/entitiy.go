package mentee

import (
	"time"
)

type MenteeCore struct {
	ID                     uint
	Name                   string
	Nickname               string
	Email                  string
	Gender                 string
	PhoneNumber            string
	ClassId                uint
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
	LogStruct              Log
	Class                  Class
}

type User struct {
	ID   uint
	Name string
}

type Log struct {
	ID        uint
	UserId    uint
	Status    string
	Notes     string
	CreatedAt time.Time
	User      User
}

type Class struct {
	ID   uint
	Name string
}

type ServiceInterface interface {
	Create(input MenteeCore) (err error)
}

type RepositoryInterface interface {
	Create(input MenteeCore) (row int, err error)
}

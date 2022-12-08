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
	LogStruct              []Log
	Class                  Class
}

type User struct {
	ID   uint
	Name string
}

type Log struct {
	ID        uint
	UserName  string
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

	GetAll(queryClass, queryEducationType, queryStatus string) (data []MenteeCore, err error)
	GetMentee(id uint) (data MenteeCore, err error)
	DeleteMentee(id uint) (err error)
	UpdateMentee(input MenteeCore, id uint) (err error)
}

type RepositoryInterface interface {
	Create(input MenteeCore) (row int, err error)
	GetAll(queryClass, queryEducationType, queryStatus string) (data []MenteeCore, err error)
	// GetAllFiltering(queryClass, queryEducationType, queryStatus string) (data []MenteeCore, err error)
	GetMentee(id uint) (data MenteeCore, err error)
	DeleteMentee(id uint) (err error)
	UpdateMentee(input MenteeCore, id uint) (err error)
}

package mentee

import "time"

type MenteeCore struct {
	ID                     uint
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
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              time.Time
}

type ServiceInterface interface {
}

type RepositoryInterface interface {
}

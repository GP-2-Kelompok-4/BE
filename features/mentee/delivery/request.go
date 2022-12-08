package delivery

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee"
)

type MenteeRequest struct {
	Name                   string `json:"name" form:"name"`
	Nickname               string `json:"nickname" form:"nickname"`
	Email                  string `json:"email" form:"email"`
	Gender                 string `json:"gender" form:"gender"`
	PhoneNumber            string `json:"phone_number" form:"phone_number"`
	TelegramAccount        string `json:"telegram_account" form:"telegram_account"`
	DiscordAccount         string `json:"discord_account" form:"discord_account"`
	ClassID                uint   `json:"class_id" form:"class_id"`
	Address                string `json:"address" form:"address"`
	HomeAddress            string `json:"home_address" form:"home_address"`
	Status                 string `json:"status" form:"status"`
	EducationType          string `json:"education_type" form:"education_type"`
	EducationMajor         string `json:"education_major" form:"education_major"`
	Institution            string `json:"institution" form:"institution"`
	Graduate               string `json:"graduate" form:"graduate"`
	EmergencyContact       string `json:"emergency_contact" form:"emergency_contact"`
	EmergencyContactName   string `json:"emergency_contact_name" form:"emergency_contact_name"`
	EmergencyContactStatus string `json:"emergency_contact_status" form:"emergency_contact_status"`
}

var layout2 = "2006-01-02"

func toCore(data MenteeRequest) mentee.MenteeCore {
	graduate, _ := time.Parse(layout2, data.Graduate)
	return mentee.MenteeCore{
		// ID:                     0,
		Name:                   data.Name,
		Nickname:               data.Nickname,
		Email:                  data.Email,
		Gender:                 data.Gender,
		PhoneNumber:            data.PhoneNumber,
		TelegramAccount:        data.TelegramAccount,
		DiscordAccount:         data.DiscordAccount,
		ClassId:                uint(data.ClassID),
		Address:                data.Address,
		HomeAddress:            data.HomeAddress,
		Status:                 data.Status,
		EducationType:          data.EducationType,
		EducationMajor:         data.EducationMajor,
		Institution:            data.Institution,
		Graduate:               graduate,
		EmergencyContact:       data.EmergencyContact,
		EmergencyContactName:   data.EmergencyContactName,
		EmergencyContactStatus: data.EmergencyContactStatus,
		// LogStruct:              mentee.Log{},
		// Class:                  mentee.Class{},
	}
}

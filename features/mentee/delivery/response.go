package delivery

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee"
)

type ClassResponse struct {
}

type ResponseDataCore struct {
	Name                   string    `json:"name" form:"name"`
	Nickname               string    `json:"nickname" form:"nickname"`
	Email                  string    `json:"email" form:"email"`
	Gender                 string    `json:"gender" form:"gender"`
	PhoneNumber            string    `json:"phone_number" form:"phone_number"`
	TelegramAccount        string    `json:"telegram_account" form:"telegram_account"`
	DiscordAccount         string    `json:"discord_account" form:"discord_account"`
	ClassID                uint      `json:"class_id" form:"class_id"`
	Address                string    `json:"address" form:"address"`
	HomeAddress            string    `json:"home_address" form:"home_address"`
	Status                 string    `json:"status" form:"status"`
	EducationType          string    `json:"education_type" form:"education_type"`
	EducationMajor         string    `json:"education_major" form:"education_major"`
	Institution            string    `json:"institution" form:"institution"`
	Graduate               time.Time `json:"graduate" form:"graduate"`
	EmergencyContact       string    `json:"emergency_contact" form:"emergency_contact"`
	EmergencyContactName   string    `json:"emergency_contact_name" form:"emergency_contact_name"`
	EmergencyContactStatus string    `json:"emergency_contact_status" form:"emergency_contact_status"`
}

func fromCoreDataResponse(dataCore mentee.MenteeCore) ResponseDataCore {
	return ResponseDataCore{
		Name:                   dataCore.Name,
		Nickname:               dataCore.Nickname,
		Email:                  dataCore.Email,
		Gender:                 dataCore.Gender,
		PhoneNumber:            dataCore.PhoneNumber,
		TelegramAccount:        dataCore.TelegramAccount,
		DiscordAccount:         dataCore.DiscordAccount,
		ClassID:                dataCore.ClassId,
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
}

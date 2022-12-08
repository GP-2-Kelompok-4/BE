package delivery

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee"
)

type ResponseDataCore struct {
	Name                   string        `json:"name" form:"name"`
	Nickname               string        `json:"nickname" form:"nickname"`
	Email                  string        `json:"email" form:"email"`
	Gender                 string        `json:"gender" form:"gender"`
	PhoneNumber            string        `json:"phone_number" form:"phone_number"`
	TelegramAccount        string        `json:"telegram_account" form:"telegram_account"`
	DiscordAccount         string        `json:"discord_account" form:"discord_account"`
	ClassID                uint          `json:"class_id" form:"class_id"`
	Address                string        `json:"address" form:"address"`
	HomeAddress            string        `json:"home_address" form:"home_address"`
	Status                 string        `json:"status" form:"status"`
	EducationType          string        `json:"education_type" form:"education_type"`
	EducationMajor         string        `json:"education_major" form:"education_major"`
	Institution            string        `json:"institution" form:"institution"`
	Graduate               time.Time     `json:"graduate" form:"graduate"`
	EmergencyContact       string        `json:"emergency_contact" form:"emergency_contact"`
	EmergencyContactName   string        `json:"emergency_contact_name" form:"emergency_contact_name"`
	EmergencyContactStatus string        `json:"emergency_contact_status" form:"emergency_contact_status"`
	ClassName              string        `json:"class_name" form:"class_name"`
	Logs                   []LogResponse `json:"logs"`
}

type LogResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"user name"`
	Log      string `json:"Log"`
	Status   string `json:"status"`
	MenteeID uint   `json:"mentee_id"`
	UserID   uint   `json:"user_id"`
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
		ClassName:              dataCore.Class.Name,
	}
}

func fromCoreList(dataCore []mentee.MenteeCore) []ResponseDataCore {
	var dataResponse []ResponseDataCore
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreDataResponse(v))
	}
	return dataResponse
}

type DetailMenteeResponse struct {
	Name                   string        `json:"name" form:"name"`
	Nickname               string        `json:"nickname" form:"nickname"`
	Email                  string        `json:"email" form:"email"`
	Gender                 string        `json:"gender" form:"gender"`
	PhoneNumber            string        `json:"phone_number" form:"phone_number"`
	TelegramAccount        string        `json:"telegram_account" form:"telegram_account"`
	DiscordAccount         string        `json:"discord_account" form:"discord_account"`
	Address                string        `json:"address" form:"address"`
	HomeAddress            string        `json:"home_address" form:"home_address"`
	Status                 string        `json:"status" form:"status"`
	EducationType          string        `json:"education_type" form:"education_type"`
	EducationMajor         string        `json:"education_major" form:"education_major"`
	Institution            string        `json:"institution" form:"institution"`
	EmergencyContact       string        `json:"emergency_contact" form:"emergency_contact"`
	EmergencyContactName   string        `json:"emergency_contact_name" form:"emergency_contact_name"`
	EmergencyContactStatus string        `json:"emergency_contact_status" form:"emergency_contact_status"`
	ClassName              string        `json:"class_name" form:"class_name"`
	Logs                   []LogResponse `json:"logs"`
}

func menteeDetail(dataCore mentee.MenteeCore) DetailMenteeResponse {

	var ArrLogs []LogResponse

	for _, val := range dataCore.LogStruct {
		ArrLogs = append(ArrLogs, LogResponse{
			ID:       val.ID,
			Username: val.User.Name,
			Log:      val.Notes,
			Status:   val.Status,
			MenteeID: dataCore.ID,
			UserID:   val.User.ID,
		})
	}

	return DetailMenteeResponse{
		Name:                   dataCore.Name,
		Nickname:               dataCore.Nickname,
		Status:                 dataCore.Status,
		Gender:                 dataCore.Gender,
		Address:                dataCore.Address,
		HomeAddress:            dataCore.HomeAddress,
		Email:                  dataCore.Email,
		TelegramAccount:        dataCore.TelegramAccount,
		DiscordAccount:         dataCore.DiscordAccount,
		PhoneNumber:            dataCore.PhoneNumber,
		EmergencyContact:       dataCore.EmergencyContact,
		EmergencyContactStatus: dataCore.EmergencyContactStatus,
		EducationType:          dataCore.EducationType,
		EducationMajor:         dataCore.EducationMajor,
		Institution:            dataCore.Institution,
		ClassName:              dataCore.Class.Name,
		Logs:                   ArrLogs,
	}
}

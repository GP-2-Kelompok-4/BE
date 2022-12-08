package delivery

import (
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log"
)

type LogRequest struct {
	UserID   uint   `json:"user_id" form:"user_id"`
	UserName string `json:"user_name" form:"user_name"`
	MenteeID uint   `json:"mentee_id" form:"mentee_id"`
	Notes    string `json:"notes" form:"notes"`
	Status   string `json:"status" form:"status"`
}

func requestToCore(logInput LogRequest) log.CoreLog {
	return log.CoreLog{
		UserID:   logInput.UserID,
		UserName: logInput.UserName,
		MenteeID: logInput.MenteeID,
		Notes:    logInput.Notes,
		Status:   logInput.Status,
	}
}

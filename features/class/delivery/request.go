package delivery

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
)

type ClassRequest struct {
	Name         string `json:"name" form:"name"`
	UserID       uint   `json:"user_id" form:"user_id"`
	StartDate    string `json:"start_date" form:"start_date"`
	GraduateDate string `json:"graduate_date" form:"graduate_date"`
}

// var layoutFormat = "January"
var layout2 = "2006-01-02"

func toCore(data ClassRequest) class.ClassCore {
	start, _ := time.Parse(layout2, data.StartDate)
	end, _ := time.Parse(layout2, data.GraduateDate)
	return class.ClassCore{
		Name:         data.Name,
		UserID:       uint(data.UserID),
		StartDate:    start,
		GraduateDate: end,
	}
}

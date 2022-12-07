package delivery

import (
	"time"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
)

type ClassResponse struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	StartDate    time.Time `json:"start_date"`
	GraduateDate time.Time `json:"graduate_date"`
	// UserID       uint      `json:"user_id"`
	PIC string `json:"PIC"`
}

func fromCore(dataCore class.ClassCore) ClassResponse {
	return ClassResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		StartDate:    dataCore.StartDate,
		GraduateDate: dataCore.GraduateDate,
		// UserID:       dataCore.UserID,
		PIC: dataCore.User.Name,
	}
}

func fromCoreList(dataCore []class.ClassCore) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

// func fromCoreDetail(dataCore class.ClassDetailCore) ClassDetailResponse {
// 	return ClassDetailResponse{
// 		ID:           dataCore.ID,
// 		Name:         dataCore.Name,
// 		StartDate:    dataCore.StartDate,
// 		GraduateDate: dataCore.GraduateDate,
// 		PIC:          dataCore.PIC,
// 	}
// }

// func fromCoreListDetail(dataCore []class.ClassDetailCore) []ClassDetailResponse {
// 	var dataResponse []ClassDetailResponse
// 	for _, v := range dataCore {
// 		dataResponse = append(dataResponse, fromCoreDetail(v))
// 	}
// 	return dataResponse
// }

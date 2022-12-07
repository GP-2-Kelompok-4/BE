package delivery

import "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"

type UserResponse struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
	Role         string `json:"role" form:"role"`
	Address      string `json:"address" form:"address"`
	Home_Address string `json:"home_address" form:"home_address"`
	Team         string `json:"team" form:"team"`
	Status       string `json:"status" form:"status"`
	Gender       string `json:"gender" form:"gender"`
}

type AddResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func FromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		Email:        dataCore.Email,
		Phone_Number: dataCore.Phone_Number,
		Role:         dataCore.Role,
		Address:      dataCore.Address,
		Home_Address: dataCore.Home_Address,
		Team:         dataCore.Team,
		Status:       dataCore.Status,
		Gender:       dataCore.Gender,
	}
}

func FromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, FromCore(v))
	}
	return dataResponse

}

func AddFromCore(dataCore user.Core) AddResponse {
	return AddResponse{
		ID:    dataCore.ID,
		Name:  dataCore.Name,
		Email: dataCore.Email,
	}
}

// func DataUserRespon(users []user.Core) []UserResponse {
// 	var dataResponse []UserResponse
// 	for _, v := range users {
// 		dataResponse = append(dataResponse, UserResponse{
// 			ID:           v.ID,
// 			Name:         v.Name,
// 			Email:        v.Email,
// 			Phone_Number: v.Phone_Number,
// 			Role:         v.Role,
// 			Address:      v.Address,
// 			Home_Address: v.Home_Address,
// 			Team:         v.Team,
// 			Status:       v.Status,
// 			Gender:       v.Gender,
// 		})
// 	}
// 	return dataResponse
// }

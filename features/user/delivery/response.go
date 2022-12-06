package delivery

import "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"

type UserResponse struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	Phone_Number string
	Role         string
	Address      string
	Home_Address string
	Team         string
	Status       string
	Gender       string
}

func FromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		Email:        dataCore.Email,
		Password:     dataCore.Password,
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

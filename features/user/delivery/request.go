package delivery

import "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"

type UserRequest struct {
	Name         string `json:"name" form:"name"`
	Email        string `json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
	Role         string `json:"role" form:"role"`
	Address      string `json:"address" form:"adress"`
	Home_Address string `json:"home_address" form:"home_address"`
	Team         string `json:"team" form:"team"`
	Status       string `json:"status" form:"status"`
	Gender       string `json:"gender" form:"gender"`
	File         string `json:"file" form:"file"`
}

func requestToCore(userInput UserRequest) user.Core {
	userCoreData := user.Core{
		Name:         userInput.Name,
		Email:        userInput.Email,
		Password:     userInput.Password,
		Phone_Number: userInput.Phone_Number,
		Role:         userInput.Role,
		Address:      userInput.Address,
		Home_Address: userInput.Home_Address,
		Team:         userInput.Team,
		Status:       userInput.Status,
		Gender:       userInput.Gender,
		ImageUrl:     userInput.File,
	}
	return userCoreData
}

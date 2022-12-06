package delivery

import "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

func requestToCore(userInput UserRequest) user.Core {
	userCoreData := user.Core{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: userInput.Password,
		Role:     userInput.Role,
	}
	return userCoreData
}

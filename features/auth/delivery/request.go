package delivery

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// func reqToCore(dataRequest LoginRequest) auth.Core {
// 	return auth.Core{
// 		Email:    dataRequest.Email,
// 		Password: dataRequest.Password,
// 	}
// }

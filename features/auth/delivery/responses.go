package delivery

type UserResponse struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

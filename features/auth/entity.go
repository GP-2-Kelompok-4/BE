package auth

import "time"

type Core struct {
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
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Login struct {
	ID    uint
	Email string
	Name  string
	Role  string
	Token string
}

type ServiceInterface interface {
	Login(email, password string) (LoginData Login, err error)
}

type RepositoryInterface interface {
	Login(email, password string) (loginData Core, err error)
}

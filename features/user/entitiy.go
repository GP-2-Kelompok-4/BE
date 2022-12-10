package user

import (
	"time"

	"github.com/labstack/echo/v4"
)

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
	ImageUrl     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ServiceInterface interface {
	GetAllUser() (data []Core, err error)
	AddUser(input Core, c echo.Context) (err error)
	UpdateUser(input Core, id uint) (err error)
	DeleteUser(id uint) (err error)
	UpdateById(input Core, id uint) (err error)
}

type RepositoryInterface interface {
	GetAllUser() (data []Core, err error)
	AddUser(input Core) (row int, err error)
	UpdateUser(input Core, id uint) (err error)
	DeleteUser(id uint) (err error)
	UpdateById(input Core, id uint) (err error)
}

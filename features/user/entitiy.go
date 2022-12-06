package user

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

type ServiceInterface interface {
	GetAllUser() (data []Core, err error)
	AddUser(input Core) (err error)
	UpdateUser(input Core, id int) (err error)
	DeleteUser(id int) (err error)
	UpdateById(input Core, id int) (err error)
}

type RepositoryInterface interface {
	GetAllUser() (data []Core, err error)
	AddUser(input Core) (row int, err error)
	UpdateUser(input Core, id int) (err error)
	DeleteUser(id int) (err error)
	UpdateById(input Core, id int) (err error)
}

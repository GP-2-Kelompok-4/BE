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
	Home_Adress  string
	Team         string
	Status       string
	Gender       string
	Created_At   time.Time
	Updated_At   time.Time
	Deleted_AT   time.Time
}

type ServiceInterface interface {
	// GetAll() (data []Core, err error)
	AddUser(input Core) (row int, err error)
	// DeleteUser(id int) (err error)
	// UpdateUser(input Core, id int) (err error)
	// UpdateById(id int) (data Core, err error)
}

type RepositoryInterface interface {
	// GetAll() (data []Core, err error)
	AddUser(input Core) (row int, err error)
	// DeleteUser(id int) (err error)
	// UpdateUser(input Core, id int) (err error)
	// UpdateById(id int) (data Core, err error)
}

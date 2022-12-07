package class

import (
	"time"
)

type ClassCore struct {
	ID           uint
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
	UserID       uint
	User         User
}

type User struct {
	ID   uint
	Name string
}

// type ClassDetailCore struct {
// 	ID           uint
// 	Name         string
// 	StartDate    time.Time
// 	GraduateDate time.Time
// 	PIC          string
// }

type ServiceInterface interface {
	CreateClass(data ClassCore) (err error)
	GetAllClassess() (data []ClassCore, err error)
	GetClassById(id uint) (data ClassCore, err error)
	UpdateClass(input ClassCore, id uint) (data ClassCore, err error)
	DeleteClass(id uint) (err error)
}

type RepositoryInterface interface {
	InsertClass(data ClassCore) (row int, err error)
	GetAllClassess() (data []ClassCore, err error)
	GetClassById(id uint) (data ClassCore, err error)
	UpdateClass(input ClassCore, id uint) (data ClassCore, err error)
	DeleteClass(id uint) (err error)
}

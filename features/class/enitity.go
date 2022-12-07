package class

import (
	// "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
	"time"
)

type ClassCore struct {
	ID           uint
	Name         string
	StartDate    time.Time
	GraduateDate time.Time
	UserID       uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type User struct {
	ID   uint
	Name string
}

type Mentee struct {
	ID   uint
	Name string
}

type ServiceInterface interface {
	CreateClass(data ClassCore) (err error)
	GetAllClassess() (data []ClassCore, err error)
	GetClassById(id uint) (data ClassCore, err error)
	UpdateClass(input ClassCore, id uint) (data ClassCore, err error)
}

type RepositoryInterface interface {
	InsertClass(data ClassCore) (row int, err error)
	GetAllClassess() (data []ClassCore, err error)
	GetClassById(id uint) (data ClassCore, err error)
	UpdateClass(input ClassCore, id uint) (data ClassCore, err error)
}

package class

import (
	"time"
)

type Core struct {
	ID            uint
	Name          string
	StartDate     time.Time
	GraduatedDate time.Time
	UserId        uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type ServiceInterface interface {
}

type RepositoryInterface interface {
}

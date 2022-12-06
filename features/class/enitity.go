package class

import (
	// "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
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

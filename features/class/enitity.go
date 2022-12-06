package class

import (
	// "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
	"time"
)

type ClassCore struct {
	ID            uint
	Name          string
	StartDate     time.Time
	GraduatedDate time.Time
	UserID        uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type ServiceInterface interface {
}

type RepositoryInterface interface {
}

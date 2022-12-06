package log

import "time"

type CoreLog struct {
	ID        uint
	UserID    uint
	MenteeID  uint
	Notes     string
	Status    string
	CreatedAt time.Time
}

type ServiceInterface interface {
}

type RepositoryInterface interface {
}

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
	CreateLog(data CoreLog) (res CoreLog, err error)
}

type RepositoryInterface interface {
	CreateLog(data CoreLog) (res CoreLog, err error)
}

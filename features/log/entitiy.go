package log

import "time"

type CoreLog struct {
	ID        uint
	UserID    uint
	MenteeID  uint
	Notes     string
	Status    string
	CreatedAt time.Time
	User      User
	Mentee    Mentee
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
	CreateLog(input CoreLog) (err error)
}

type RepositoryInterface interface {
	CreateLog(input CoreLog) (err error)
}

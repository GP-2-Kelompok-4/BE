package service

import "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log"

type logService struct {
	logRepository log.RepositoryInterface
}

func New(repo log.RepositoryInterface) log.ServiceInterface {
	return &logService{
		logRepository: repo,
	}
}

// CreateLogs implements log.ServiceInterface
func (repo *logService) CreateLog(data log.CoreLog) (err error) {
	panic("unimplemented")
}

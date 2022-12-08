package service

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log"
)

type logService struct {
	logRepository log.RepositoryInterface
}

func New(repo log.RepositoryInterface) log.ServiceInterface {
	return &logService{
		logRepository: repo,
	}
}

// CreateLogs implements log.ServiceInterface
func (repo *logService) CreateLog(input log.CoreLog) (err error) {
	errCreate := repo.logRepository.CreateLog(input)
	if errCreate != nil {
		return errors.New("failed create log, error query")
	}

	return nil

}

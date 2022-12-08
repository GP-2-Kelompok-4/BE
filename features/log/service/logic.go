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
func (repo *logService) CreateLog(data log.CoreLog) (res log.CoreLog, err error) {
	result, errCreate := repo.logRepository.CreateLog(data)
	if errCreate != nil {
		return res, errors.New("failed create log, error query")
	}

	var logData log.CoreLog
	logData.ID = result.ID
	logData.UserID = result.UserID
	logData.MenteeID = result.MenteeID
	logData.CreatedAt = res.CreatedAt
	return logData, nil

}

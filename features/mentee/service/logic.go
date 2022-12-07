package service

import (
	// "errors"

	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee"
)

type menteeService struct {
	menteeRepository mentee.RepositoryInterface
}

func New(repo mentee.RepositoryInterface) mentee.ServiceInterface {
	return &menteeService{
		menteeRepository: repo,
	}
}

// Create implements mentee.ServiceInterface
func (service *menteeService) Create(input mentee.MenteeCore) (err error) {

	_, errCreate := service.menteeRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}
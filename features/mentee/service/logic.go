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

// GetAll implements mentee.ServiceInterface
func (service *menteeService) GetAll(queryClass, queryEducationType, queryStatus string) (data []mentee.MenteeCore, err error) {
	data, err = service.menteeRepository.GetAll(queryClass, queryEducationType, queryStatus)
	if err != nil {
		return nil, err
	}
	return
}

func (service *menteeService) DeleteMentee(id uint) (err error) {
	err = service.menteeRepository.DeleteMentee(id)
	if err != nil {
		return errors.New("failed to delete data, error query")
	}
	return nil
}

// UpdateClass implements mentee.ServiceInterface
func (service *menteeService) UpdateMentee(input mentee.MenteeCore, id uint) (err error) {
	err = service.menteeRepository.UpdateMentee(input, id)
	if err != nil {
		return errors.New("failed update data, error query")

	}
	return nil
}

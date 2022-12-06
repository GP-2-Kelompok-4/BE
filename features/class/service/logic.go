package service

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
)

type classService struct {
	classRepository class.RepositoryInterface
}

func New(repo class.RepositoryInterface) class.ServiceInterface {
	return &classService{
		classRepository: repo,
	}
}

// CreateClass implements class.ServiceInterface
func (service *classService) CreateClass(data class.ClassCore) (err error) {
	_, errCreate := service.classRepository.InsertClass(data)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAllClassess implements class.ServiceInterface
func (service *classService) GetAllClassess() (data []class.ClassCore, err error) {
	data, err = service.classRepository.GetAllClassess()
	return
}

// GetClassById implements class.ServiceInterface
func (service *classService) GetClassById(id uint) (data class.ClassCore, err error) {
	data, err = service.classRepository.GetClassById(id)
	return
}

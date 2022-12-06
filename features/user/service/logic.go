package service

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"
)

type userService struct {
	userRepository user.RepositoryInterface
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) AddUser(input user.Core) (err error) {
	_, errCreate := service.userRepository.AddUser(input)
	if errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

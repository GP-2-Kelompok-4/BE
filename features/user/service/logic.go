package service

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/utils/helper"
)

type userService struct {
	userRepository user.RepositoryInterface
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) GetAllUser() (data []user.Core, err error) {
	data, err = service.userRepository.GetAllUser()
	return

}

func (service *userService) AddUser(input user.Core) (err error) {
	input.Gender = "Male"
	input.Team = "Academic"
	input.Status = "Active"

	hash_pass, errHash := helper.HashPassword(input.Password)
	if errHash != nil {
		return errHash
	}
	input.Password = hash_pass

	_, errCreate := service.userRepository.AddUser(input)
	if errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

func (service *userService) UpdateUser(dataCore user.Core, id uint) (err error) {
	dataCore.Role = "User"
	errUpdate := service.userRepository.UpdateUser(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

func (service *userService) DeleteUser(id uint) (err error) {
	errDel := service.userRepository.DeleteUser(id)
	if errDel != nil {
		return errors.New("failed delete user, error query")
	}
	return nil
}

func (service *userService) UpdateById(dataCore user.Core, id uint) (err error) {
	errUpdate := service.userRepository.UpdateById(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

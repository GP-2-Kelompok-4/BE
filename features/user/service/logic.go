package service

import (
	"errors"
	"log"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/utils/helper"
	"github.com/labstack/echo/v4"
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

func (service *userService) AddUser(input user.Core, c echo.Context) (err error) {
	input.Gender = "Not-Assigned"
	input.Team = "Not-Assigned"
	input.Status = "Active"

	hash_pass, errHash := helper.HashPassword(input.Password)
	if errHash != nil {
		return errHash
	}
	input.Password = hash_pass

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := helper.UploadProfile(c)
		if err != nil {
			return errors.New("Registration Failed. Cannot Upload Data.")
		}
		log.Print(res)
		input.ImageUrl = res
	} else {
		input.ImageUrl = "https://project3bucker.s3.ap-southeast-1.amazonaws.com/dummy-profile-pic.png"
	}

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

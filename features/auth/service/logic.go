package service

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/middlewares"
)

type authService struct {
	authRepo auth.RepositoryInterface
}

func New(data auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authRepo: data,
	}
}

func (service *authService) Login(email, password string) (LoginData auth.Login, err error) {

	if email == "" || password == "" {
		return LoginData, errors.New("email and password must be filled")
	}

	result, err := service.authRepo.Login(email, password)

	if err != nil {
		return LoginData, err
	}

	// cekPass := helper.CheckPasswordHash(password, result.Password)

	// if !cekPass {
	// 	return LoginData, errors.New("login failed")
	// }

	token, errToken := middlewares.CreateToken(int(result.ID), result.Role)
	if errToken != nil {
		return LoginData, errToken
	}

	var DataLogin auth.Login
	DataLogin.ID = result.ID
	DataLogin.Email = result.Email
	DataLogin.Role = result.Role
	DataLogin.Token = token

	return DataLogin, err
}

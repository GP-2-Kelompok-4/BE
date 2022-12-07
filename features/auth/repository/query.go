package repository

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth"
	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authRepo{
		db: db,
	}
}

func (repo *authRepo) Login(email, password string) (loginData auth.Core, err error) {

	var userData = User{}
	tx := repo.db.Where("email = ? AND password = ?", email, password).First(&userData)
	if tx.Error != nil {
		return loginData, tx.Error
	}
	if tx.RowsAffected == 0 {
		return loginData, errors.New("login failed")
	}
	loginData = ToCore(userData)

	return loginData, nil
}

// func Login(email, password string) (string, error) {
// 	var userData models.User
// 	tx := config.DB.Where("email = ? AND password = ?", email, password).First(&userData)
// 	if tx.Error != nil {
// 		return "", tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return "", errors.New("login failed")
// 	}

// 	token, errToken := middlewares.CreateToken(int(userData.ID), userData.Role)
// 	if errToken != nil {
// 		return "", errToken
// 	}

// 	return token, nil
// }

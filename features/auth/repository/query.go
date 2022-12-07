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
	userModel := User{}
	loginData = auth.Core{}
	tx := repo.db.Where("email = ? ", email).First(&userModel)
	if tx.Error != nil {
		return loginData, tx.Error
	}
	// check_result := helper.CheckPasswordHash(password, userModel.Password)

	// if !check_result {
	// 	return loginData, errors.New("password salah")
	// }

	if tx.RowsAffected == 0 {
		return loginData, errors.New("login failed")
	}

	loginData = ToCore(userModel)
	return loginData, nil

}

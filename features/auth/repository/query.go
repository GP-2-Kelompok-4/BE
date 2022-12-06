package repository

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth"
	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) Login(email, password string) (loginData auth.Core, err error) {

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

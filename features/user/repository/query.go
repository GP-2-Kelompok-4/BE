package repository

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// dependency injection
func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) AddUser(input user.Core) (row int, err error) {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

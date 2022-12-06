package repository

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// InsertClass implements class.RepositoryInterface
func (repo *userRepository) InsertClass(data class.ClassCore) (row int, err error) {
	classGorm := fromCore(data)
	tx := repo.db.Create(&classGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

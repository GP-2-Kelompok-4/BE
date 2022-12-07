package repository

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee"
	"gorm.io/gorm"
)

type menteeRepository struct {
	db *gorm.DB
}

// Create implements mentee.RepositoryInterface

func New(db *gorm.DB) mentee.RepositoryInterface {
	return &menteeRepository{
		db: db,
	}
}

func (repo *menteeRepository) Create(input mentee.MenteeCore) (row int, err error) {
	menteeGorm := fromCore(input)
	tx := repo.db.Create(&menteeGorm)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

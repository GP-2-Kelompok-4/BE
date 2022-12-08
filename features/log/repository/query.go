package repository

import (
	"errors"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log"
	"gorm.io/gorm"
)

type logRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) log.RepositoryInterface {
	return &logRepository{
		db: db,
	}
}

// CreateLog implements log.RepositoryInterface
func (repo *logRepository) CreateLog(data log.CoreLog) (err error) {
	var mentee Mentee
	classGorm := fromCore(data)
	tx := repo.db.Create(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}
	tx2 := repo.db.Model(&mentee).Where("id = ?", data.MenteeID).Update("status", data.Status)
	if tx2.Error != nil {
		return tx.Error
	}
	return nil
}

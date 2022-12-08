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
func (repo *logRepository) CreateLog(data log.CoreLog) (res log.CoreLog, err error) {
	logModel := Log{}
	dataUpdate := log.CoreLog{}
	classGorm := fromCore(data)
	tx := repo.db.Create(&classGorm)
	if tx.Error != nil {
		return dataUpdate, tx.Error
	}
	if tx.RowsAffected == 0 {
		return dataUpdate, errors.New("insert failed")
	}

	tx1 := repo.db.Where("id = ? ", data.ID).First(&logModel)
	if tx1.Error != nil {
		return dataUpdate, tx.Error
	}

	tx2 := repo.db.Joins("join mentee on log.mentee_id = mentee.id").Joins("join user on log.user_id = user.id").Where("id = ?", logModel.ID).Find(&logModel)
	if tx2.Error != nil {
		return dataUpdate, tx.Error
	}

	dataUpdate = toCore(logModel)
	return dataUpdate, nil
}

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
	// logModel := Log{}
	classGorm := fromCore(data)
	tx := repo.db.Create(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed")
	}

	// tx1 := repo.db.Where("id= ? ", classGorm.ID).First(&logModel)
	// if tx1.Error != nil {
	// 	return tx.Error
	// }

	// tx2 := repo.db.Preload("User").Preload("Mentee").Find(&logModel, logModel.ID)
	// if tx2.Error != nil {
	// 	return tx.Error
	// }
	return nil
}

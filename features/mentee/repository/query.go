package repository

import (
	"errors"
	"fmt"

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
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm)
	fmt.Println(tx)

	if tx.Error != nil {
		return -1, tx.Error
	}

	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil

}

// DeleteMentee implements mentee.RepositoryInterface
func (repo *menteeRepository) DeleteMentee(id uint) (err error) {
	var mentee Mentee

	tx := repo.db.Delete(&mentee, id)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("delete failed")
	}
	return nil
}

// UpdateClass implements mentee.RepositoryInterface
func (repo *menteeRepository) UpdateMentee(input mentee.MenteeCore, id uint) (data mentee.MenteeCore, err error) {
	var mentee Mentee

	inputData := fromCore(input)
	tx := repo.db.Model(&mentee).Where("id = ?", id).Updates(inputData)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = mentee.toCore()
	return data, nil
}

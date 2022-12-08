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


func New(db *gorm.DB) mentee.RepositoryInterface {
	return &menteeRepository{
		db: db,
	}
}

// Create implements mentee.RepositoryInterface
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


// GetAll implements mentee.RepositoryInterface
func (repo *menteeRepository) GetAll(queryClass, queryEducationType, queryStatus string) (data []mentee.MenteeCore, err error) {
	var mentees []Mentee
	tx := repo.db.Preload("Class").Find(&mentees)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(mentees)
	return dataCore, nil
}

// GetAllFiltering implements mentee.RepositoryInterface
// func (repo *menteeRepository) GetAllFiltering(queryClass string, queryEducationType string, queryStatus string) (data []mentee.MenteeCore, err error) {
// 	var mentees []Mentee
// 	// tx := repo.db.Preload("Class").Where(&Mentee{EducationType: queryEducationType, Status: queryEducationType}).Find(&mentees)
// 	tx := repo.db.Preload("Class").Where("class.name = ? AND education_type = ? AND status = ?", queryClass, queryEducationType, queryStatus).Find(&mentees)
// 	if tx.Error != nil {
// 		return nil, tx.Error
// 	}

// 	dataCore := toCoreList(mentees)
// 	return dataCore, nil
// }

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


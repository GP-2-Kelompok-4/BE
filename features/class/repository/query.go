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

// GetAllClassess implements class.RepositoryInterface
func (repo *userRepository) GetAllClassess() (data []class.ClassCore, err error) {
	// db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
	var classes []Class
	tx := repo.db.Model(&Class{}).Select("classes.ID, classes.Name, classes.StartDate, classes.GraduateDate, users.name as PIC").Joins("left join classes on classes.user_id = users.id").Find(&classes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(classes)
	return dataCore, nil
}

// GetClassById implements class.RepositoryInterface
func (repo *userRepository) GetClassById(id uint) (data class.ClassCore, err error) {
	var class Class

	tx := repo.db.First(&class, id)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = class.toCore()
	return data, nil
}

// UpdateClass implements class.RepositoryInterface
func (repo *userRepository) UpdateClass(input class.ClassCore, id uint) (data class.ClassCore, err error) {
	var class Class

	inputData := fromCore(input)
	tx := repo.db.Model(&class).Where("id = ?", id).Updates(inputData)
	if tx.Error != nil {
		return data, tx.Error
	}
	data = class.toCore()
	return data, nil
}

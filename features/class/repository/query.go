package repository

import (
	// "errors"

	// "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// func New(db *gorm.DB) Class.RepositoryInterface {
// 	return &userRepository{
// 		db: db,
// 	}
// }

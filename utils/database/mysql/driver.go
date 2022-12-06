package mysql

import (
	"fmt"
	"log"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/config"
	classRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class/repository"
	logRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log/repository"
	menteeRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee/repository"
	userRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&userRepo.User{})
	db.AutoMigrate(&classRepo.Class{})
	db.AutoMigrate(&menteeRepo.Mentee{})
	db.AutoMigrate(&logRepo.Log{})
}

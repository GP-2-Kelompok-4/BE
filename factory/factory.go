package factory

import (
	userDelivery "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/delivery"
	userRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/repository"
	userService "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

}

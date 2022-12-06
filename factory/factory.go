package factory

import (
	userDelivery "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/delivery"
	userRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/repository"
	userService "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/service"

	classDelivery "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class/delivery"
	classRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class/repository"
	classService "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	classRepoFactory := classRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	classServiceFactory := classService.New(classRepoFactory)
	userDelivery.New(userServiceFactory, e)
	classDelivery.New(classServiceFactory, e)

}

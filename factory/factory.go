package factory

import (
	userDelivery "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/delivery"
	userRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/repository"
	userService "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user/service"

	classDelivery "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class/delivery"
	classRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class/repository"
	classService "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class/service"

	authDelivery "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth/delivery"
	authRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth/repository"
	authService "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/auth/service"

	menteeDelivery "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee/delivery"
	menteeRepo "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee/repository"
	menteeService "github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	classRepoFactory := classRepo.New(db)
	classServiceFactory := classService.New(classRepoFactory)
	classDelivery.New(classServiceFactory, e)

	menteeRepoFactory := menteeRepo.New(db)
	menteeServiceFactory := menteeService.New(menteeRepoFactory)
	menteeDelivery.New(menteeServiceFactory, e)

}

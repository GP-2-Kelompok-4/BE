package delivery

// import (
// 	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log"
// 	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/middlewares"
// )

// type LogDelivery struct {
// 	logService log.ServiceInterface
// }

// func New(service log.ServiceInterface, e *echo.Echo) {
// 	handler := &LogDelivery{
// 		logService: service,
// 	}
// 	e.POST("/logs", handler.CreateLog, middlewares.JWTMiddleware())
// }

// func (delivery *LogDelivery) CreateLog(c echo.Context) error {
// 	// userLogin := middlewares.ExtractTokenUserId(c)
// 	// if userLogin == 0 {
// 	// 	return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"))
// 	// }
// 	// logInput := LogRequest{}
// 	// errBind := c.Bind(&logInput)
// 	// if errBind != nil {
// 	// 	return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
// 	// }
// }

package delivery

import (
	"net/http"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/log"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/middlewares"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type LogDelivery struct {
	logService log.ServiceInterface
}

func New(service log.ServiceInterface, e *echo.Echo) {
	handler := &LogDelivery{
		logService: service,
	}
	e.POST("/logs", handler.CreateLog, middlewares.JWTMiddleware())
}

func (delivery *LogDelivery) CreateLog(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	name := middlewares.ExtractTokenUserName(c)
	helper.LogDebug("\n extracttokenname= ", name)

	if userId == 0 {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"))
	}

	logInput := LogRequest{}
	logInput.UserName = name
	logInput.UserID = uint(userId)

	helper.LogDebug("\n extracttokenname= ", logInput.UserName)

	errBind := c.Bind(&logInput)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found "+errBind.Error()))
	}

	logInput.UserName = name
	logInput.UserID = uint(userId)

	helper.LogDebug("\n extracttokenname= ", logInput.UserName)

	logCore := requestToCore(logInput)

	logCore.UserID = uint(userId)

	logCore.UserName = name
	helper.LogDebug("\n extracttokenname= ", logCore.UserName)

	err := delivery.logService.CreateLog(logCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error "+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success add new log"))
}

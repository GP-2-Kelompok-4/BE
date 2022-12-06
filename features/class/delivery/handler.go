package delivery

import (
	// "strconv"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/class"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/utils/helper"

	"net/http"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/middlewares"
	"github.com/labstack/echo/v4"
)

type classDelivery struct {
	classService class.ServiceInterface
}

func New(service class.ServiceInterface, e *echo.Echo) {
	handler := &classDelivery{
		classService: service,
	}

	e.POST("/classes", handler.Create, middlewares.JWTMiddleware())

}

func (delivery *classDelivery) Create(c echo.Context) error {
	userLogin := middlewares.ExtractTokenUserId(c)
	if userLogin == 0 {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"))
	}
	classInput := ClassRequest{}
	errBind := c.Bind(&classInput)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
	}

	dataCore := toCore(classInput)
	err := delivery.classService.CreateClass(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success add new class", classInput))
}

func (delivery *classDelivery) GetAll(c echo.Context) error {
	results, err := delivery.classService.GetAllClassess()
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get all classes", dataResponse))
}

package delivery

import (
	"strconv"

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
	e.GET("/classes", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/classes/:id", handler.GetClassById, middlewares.JWTMiddleware())
	e.PUT("classes/:id", handler.UpdateClass, middlewares.JWTMiddleware())
	e.DELETE("classes/:id", handler.DeleteClass, middlewares.JWTMiddleware())

}

func (delivery *classDelivery) Create(c echo.Context) error {
	// userLogin := middlewares.ExtractTokenUserId(c)
	// if userLogin == 0 {
	// 	return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"))
	// }
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

func (delivery *classDelivery) GetClassById(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))
	result, err := delivery.classService.GetClassById(uint(idParam))
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"))

	}
	dataResponse := fromCore(result)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get spesific class", dataResponse))
}

func (delivery *classDelivery) UpdateClass(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))
	inputData := ClassRequest{}
	errBind := c.Bind(&inputData)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
	}

	dataUpdateCore := toCore(inputData)
	_, err := delivery.classService.UpdateClass(dataUpdateCore, uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success update class", dataUpdateCore))
}

func (delivery *classDelivery) DeleteClass(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	err := delivery.classService.DeleteClass(uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]any{
		"message": "success delete class",
	})
}

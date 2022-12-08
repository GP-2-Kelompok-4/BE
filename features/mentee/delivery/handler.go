package delivery

import (
	// "strconv"

	"strconv"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/mentee"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/utils/helper"

	"net/http"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/middlewares"

	"github.com/labstack/echo/v4"
)

type MenteeDelivery struct {
	menteeService mentee.ServiceInterface
}

func New(service mentee.ServiceInterface, e *echo.Echo) {
	handler := &MenteeDelivery{
		menteeService: service,
	}

	e.POST("/mentees", handler.AddMentee, middlewares.JWTMiddleware())
	e.GET("/mentees", handler.GetAll, middlewares.JWTMiddleware())
	e.GET("/mentees/:id", handler.GetMenteeById, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", handler.DeleteMentee, middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", handler.UpdateMentee, middlewares.JWTMiddleware())

}

func (delivery *MenteeDelivery) AddMentee(c echo.Context) error {
	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Add new mentee, semua field harus diisi"+errBind.Error()))
	}

	dataCore := toCore(menteeInput)
	dataResponse := fromCoreDataResponse(dataCore)
	err := delivery.menteeService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success add new mentee", dataResponse))
}

func (delivery *MenteeDelivery) GetAll(c echo.Context) error {
	queryClass := c.QueryParam("class")
	queryEducationType := c.QueryParam("education_type")
	queryStatus := c.QueryParam("status")
	results, err := delivery.menteeService.GetAll(queryClass, queryEducationType, queryStatus)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	dataResponse := fromCoreList(results)

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all data users", dataResponse))
}
func (delivery *MenteeDelivery) DeleteMentee(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	errDel := delivery.menteeService.DeleteMentee(uint(id))
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete mentee "+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete mentee "))
}

func (delivery *MenteeDelivery) UpdateMentee(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))
	inputData := MenteeRequest{}
	errBind := c.Bind(&inputData)
	if errBind != nil {
		return c.JSON(http.StatusNotFound, helper.FailedResponse("requested resource was not found"+errBind.Error()))
	}

	dataUpdateCore := toCore(inputData)
	err := delivery.menteeService.UpdateMentee(dataUpdateCore, uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success update class"))

}

func (delivery *MenteeDelivery) GetMenteeById(c echo.Context) error {
	id := c.Param("id")
	id_conv, err_conv := strconv.Atoi(id)
	if err_conv != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err_conv.Error())
	}

	result, err := delivery.menteeService.GetMentee(uint(id_conv))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("fail get data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get data", menteeDetail(result)))

}

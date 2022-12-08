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

func (delivery *MenteeDelivery) DeleteUser(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	roleToken := middlewares.ExtractTokenUserRole(c)
	if roleToken != "Admin" {
		return c.JSON(http.StatusUnauthorized, helper.FailedResponse("Data only can be deleted by admin"))
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
	_, err := delivery.menteeService.UpdateMentee(dataUpdateCore, uint(idParam))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success update class", dataUpdateCore))
}

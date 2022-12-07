package delivery

import (
	// "strconv"

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
	err := delivery.menteeService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("internal server error"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success add new mentee", dataCore))
}

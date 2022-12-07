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

	e.POST("/mentees", handler.Create, middlewares.JWTMiddleware())
}

func (delivery *MenteeDelivery) Create(c echo.Context) error {
	menteeInput := MenteeRequest{}
	errBind := c.Bind(&menteeInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Failed insert data"+errBind.Error()))
	}

	dataCore := toCore(menteeInput)
	err := delivery.menteeService.Create(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("Failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusCreated, helper.SuccessResponse("success create data"))
}

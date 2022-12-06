package delivery

import (
	"net/http"
	"strconv"

	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/features/user"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/middlewares"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/utils/helper"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, e *echo.Echo) {
	handler := &UserDelivery{
		userService: service,
	}

	e.GET("/users", handler.GetAllUser, middlewares.JWTMiddleware())
	e.POST("/users", handler.AddUser)
}

func (delivery *UserDelivery) GetAllUser(c echo.Context) error {
	results, err := delivery.userService.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read all users", FromCoreList(results)))
}

func (delivery *UserDelivery) AddUser(c echo.Context) error {
	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	dataCore := requestToCore(userInput)
	err := delivery.userService.AddUser(dataCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponse("failed insert data"+err.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success create new users"))
}

func (delivery *UserDelivery) UpdateData(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error binding data "+errBind.Error()))
	}

	dataCore := requestToCore(userInput)
	errUpt := delivery.userService.UpdateUser(dataCore, id)
	if errUpt != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("Error Db update "+errUpt.Error()))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (delivery *UserDelivery) DeleteUser(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.BadRequest(errConv.Error()))
	}

	errDel := delivery.userService.DeleteUser(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete user"+errDel.Error()))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))

}

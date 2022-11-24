package controllers

import (
	_helper "api/mvc/helper"
	_models "api/mvc/models"
	_repositories "api/mvc/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {
	result, err := _repositories.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("error read data"))
	}

	return c.JSON(http.StatusOK, _helper.SuccessWithDataResponse("success read all users", result))
}

func AddUserController(c echo.Context) error {
	userInput := _models.User{}
	errBind := c.Bind(&userInput) // menangkap data yg dikirim dari req body dan disimpan ke variabel
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("Error binding data "+errBind.Error()))
	}
	errInsert := _repositories.InsertUser(userInput)
	if errInsert != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("failed insert data"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessResponse("success create user"))
}

func GetIDUserController(c echo.Context) error {
	id, errconv := strconv.Atoi(c.Param("id"))
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("Error read data "))
	}
	idUser, err := _repositories.GetUserbyID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("failed read data"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessWithDataResponse("success read user", idUser))
}

func DeleteUserController(c echo.Context) error {
	id, errconv := strconv.Atoi(c.Param("id"))
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("Error read data "))
	}

	errDel := _repositories.DeleteUser(id)
	if errDel != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("failed delete data"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessResponse("delete user success"))
}

func UpdateUserController(c echo.Context) error {
	id, errconv := strconv.Atoi(c.Param("id"))
	if errconv != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("Error read data "))
	}

	updateReq := _models.User{}
	errBind := c.Bind(&updateReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("Error read data "))
	}

	errUpdate := _repositories.UpdateUserbyID(id, updateReq)
	if errUpdate != nil {
		return c.JSON(http.StatusBadRequest, _helper.FailedResponse("failed delete data"))
	}
	return c.JSON(http.StatusOK, _helper.SuccessResponse("delete user success"))
}

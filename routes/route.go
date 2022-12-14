package routes

import (
	"api/mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetAllUserController)
	e.POST("/users", controllers.AddUserController)
	e.GET("/users/:id", controllers.GetIDUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}

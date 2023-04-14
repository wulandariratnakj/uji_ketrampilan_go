package route

import (
	"prakerja2/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.POST("/student", controller.AddStudents)
	e.GET("/student", controller.GetStudent)
	e.GET("/student/:id", controller.GetDetailStudent)
	e.DELETE("/student/:id", controller.DeleteStudent)

	e.POST("/login", controller.Login)
	return e
}

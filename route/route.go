package route

import (
	"prakerja2/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.POST("/students", controller.AddStudents)
	e.GET("/students", controller.GetStudent)
	e.GET("/students/:id", controller.GetDetailStudent)
	e.DELETE("/students/:id", controller.DeleteStudent)

	e.POST("/login", controller.Login)
	return e
}

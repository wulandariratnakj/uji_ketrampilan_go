package controller

import (
	"net/http"
	"prakerja2/model"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {

	var loginRequest model.LoginRequest
	c.Bind(&loginRequest)

	if loginRequest.Username == "" || loginRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, model.Response{
			"error", "Username or password is empty",
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		"success", loginRequest,
	})
}

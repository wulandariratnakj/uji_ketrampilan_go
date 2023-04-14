package controller

import (
	"fmt"
	"net/http"
	"prakerja2/config"
	"prakerja2/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddStudents(c echo.Context) error {
	var student model.Student
	c.Bind(&student)

	result := config.DB.Create(&student)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			"failed", nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		"success", student,
	})
}

func GetStudent(c echo.Context) error {
	var students []model.Student

	result := config.DB.Find(&students)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			"failed get data", nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		"get data sucess", students,
	})
}

func getStudentById(id int) (model.Student, error) {
	var student model.Student
	result := config.DB.First(&student, id)
	if result.Error != nil {
		return student, result.Error
	}
	return student, nil
}

func GetDetailStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	student, err := getStudentById(id)
	if err != nil {

		result := config.DB.Find(&model.Student{}, id)
		if result.RowsAffected == 0 {
			return c.JSON(http.StatusNotFound, model.Response{
				Message: fmt.Sprintf("Student with ID %d not found", id),
				Data:    nil,
			})
		} else if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Message: "Failed to delete student with ID %d",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "get data student success", Data: student,
	})
}

func DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid ID")
	}

	result := config.DB.Delete(&model.Student{}, id)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, model.Response{
			Message: fmt.Sprintf("User with ID %d not found", id),
			Data:    nil,
		})
	} else if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "Failed to delete user",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: fmt.Sprintf("User with ID %d deleted successfully", id),
		Data:    nil,
	})
}

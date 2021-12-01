package controller

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/model"
	"ATIKAH-PERCA/respon"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create Tutorial Perca
func CreateTutorial(c echo.Context) error {
	var tutorial model.TutorialPerca
	c.Bind(&tutorial)

	result := config.DB.Create(&tutorial)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Error save data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, respon.BaseRespon{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &tutorial,
	})
}

// Get All Tutorial Perca
func GetAllTutorial(c echo.Context) error {
	var tutorial = []model.TutorialPerca{}

	result := config.DB.Find(&tutorial)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &tutorial,
	})
}

// Get Tutorial Perca By Id
func GetTutorialById(c echo.Context) error {
	var tutorial model.TutorialPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&tutorial, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &tutorial,
	})
}

// Update Tutorial Perca
func UpdateTutorial(c echo.Context) error {
	var tutorial model.TutorialPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&tutorial, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&tutorial)
	result = config.DB.Save(&tutorial)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &tutorial,
	})
}

// Delete Tutorial Perca
func DeleteTutorial(c echo.Context) error {
	var tutorial model.TutorialPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&tutorial, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = config.DB.Where("id = ?", id).Unscoped().Delete(&tutorial)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Delete data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful Delete data",
		Data:    &tutorial,
	})
}

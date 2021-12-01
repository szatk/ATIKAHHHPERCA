package controller

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/model"
	"ATIKAH-PERCA/respon"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create Jenis Perca
func CreateJenisPerca(c echo.Context) error {
	var jenisperca model.JenisPerca
	c.Bind(&jenisperca)

	result := config.DB.Create(&jenisperca)
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
		Data:    &jenisperca,
	})
}

// Get All Jenis Perca
func GetAllJenisPerca(c echo.Context) error {
	var jenisperca = []model.JenisPerca{}

	result := config.DB.Find(&jenisperca)
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
		Data:    &jenisperca,
	})
}

// Get Jenis Perca By Id
func GetJenisPercaById(c echo.Context) error {
	var jenisperca model.JenisPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&jenisperca, id)

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
		Data:    &jenisperca,
	})
}

// Update Jenis Perca
func UpdateJenisPerca(c echo.Context) error {
	var jenisperca model.JenisPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&jenisperca, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&jenisperca)
	result = config.DB.Save(&jenisperca)
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
		Data:    &jenisperca,
	})
}

// Delete Jenis Perca
func DeleteJenisPerca(c echo.Context) error {
	var jenisperca model.JenisPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&jenisperca, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = config.DB.Where("id = ?", id).Unscoped().Delete(&jenisperca)
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
		Data:    &jenisperca,
	})
}

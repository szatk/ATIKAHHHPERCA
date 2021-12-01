package controller

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/model"
	"ATIKAH-PERCA/respon"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Register PekerjaPerca
func PekerjaPercaRegister(c echo.Context) error {
	var pekerja model.PekerjaPerca
	c.Bind(&pekerja)

	result := config.DB.Create(&pekerja)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Error save data to database",
			Data:    nil,
		})
	}

	if res := config.DB.Preload("HousePerca").Find(&pekerja); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Error retrieve data after saved",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, respon.BaseRespon{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &pekerja,
	})
}

// Get All Pekerja Perca Data
func GetAllPekerjaPerca(c echo.Context) error {
	var pekerja = []model.PekerjaPerca{}

	result := config.DB.Preload("HousePerca").Find(&pekerja)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &pekerja,
	})
}

// Get Pekerja Perca by ID
func GetPekerjaPecraById(c echo.Context) error {
	var pekerja model.PekerjaPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.Preload("HousePerca").First(&pekerja, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Cannot retrieve data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, respon.BaseRespon{
		Code:    http.StatusOK,
		Message: "Successful retrieve data",
		Data:    &pekerja,
	})
}

// Update Pekerja Perca
func UpdatePekerjaPerca(c echo.Context) error {
	var pekerja model.PekerjaPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&pekerja, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&pekerja)
	config.DB.Model(&pekerja).Update("HousePercaId", pekerja.HousePerca)
	result = config.DB.Save(&pekerja)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	config.DB.Preload("HousePerca").First(&pekerja, id)
	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &pekerja,
	})
}

// Delete Pekerja Perca
func DeletePekerjaPerca(c echo.Context) error {
	var pekerja model.PekerjaPerca

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&pekerja, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = config.DB.Where("id = ?", id).Unscoped().Delete(&pekerja)
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
		Data:    &pekerja,
	})
}

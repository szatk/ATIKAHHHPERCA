package controller

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/model"
	"ATIKAH-PERCA/respon"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create Perca Sepatu
func CreatePercaSepatu(c echo.Context) error {
	var percasepatu model.PercaSepatu
	c.Bind(&percasepatu)

	result := config.DB.Create(&percasepatu)
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
		Data:    &percasepatu,
	})
}

// Get All Perca Sepatu
func GetAllPercaSepatu(c echo.Context) error {
	var percasepatu = []model.PercaSepatu{}

	result := config.DB.Find(&percasepatu)
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
		Data:    &percasepatu,
	})
}

// Get Perca Sepatu by ID
func GetPercaSepatuById(c echo.Context) error {
	var percasepatu model.PercaSepatu

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&percasepatu, id)

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
		Data:    &percasepatu,
	})
}

// Update Perca Sepatu
func UpdatePercaSepatu(c echo.Context) error {
	var percasepatu model.PercaSepatu

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&percasepatu, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&percasepatu)
	result = config.DB.Save(&percasepatu)
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
		Data:    &percasepatu,
	})
}

// Delete Perca Sepatu
func DeletePercaSepatu(c echo.Context) error {
	var percasepatu model.PercaSepatu

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&percasepatu, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotAcceptable, respon.BaseRespon{
			Code:    http.StatusNotAcceptable,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = config.DB.Where("id = ?", id).Unscoped().Delete(&percasepatu)
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
		Data:    &percasepatu,
	})
}

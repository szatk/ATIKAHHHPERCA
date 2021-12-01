package controller

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/model"
	"ATIKAH-PERCA/respon"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create Info House
func CreateInfoHouse(c echo.Context) error {
	var infohouse model.InfoHouse
	c.Bind(&infohouse)

	result := config.DB.Create(&infohouse)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Error save data to database",
			Data:    nil,
		})
	}

	if res := config.DB.Preload("HousePerca").Find(&infohouse); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Error retrieve data after saved",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, respon.BaseRespon{
		Code:    http.StatusCreated,
		Message: "Successful create data",
		Data:    &infohouse,
	})
}

// Get All Info House
func GetAllInfoHouse(c echo.Context) error {
	var infohouse = []model.InfoHouse{}

	result := config.DB.Preload("HousePerca").Find(&infohouse)
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
		Data:    &infohouse,
	})
}

// Get Info By Id
func GetInfoHouseById(c echo.Context) error {
	var infohouse model.InfoHouse

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.Preload("HousePerca").First(&infohouse, id)
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
		Data:    &infohouse,
	})
}

// Update Info House
func UpdateInfoHouse(c echo.Context) error {
	var infohouse model.InfoHouse

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&infohouse, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	c.Bind(&infohouse)
	config.DB.Model(&infohouse).Update("HousePercaId", infohouse.HousePerca)
	result = config.DB.Save(&infohouse)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, respon.BaseRespon{
			Code:    http.StatusInternalServerError,
			Message: "Cannot Update data to database",
			Data:    nil,
		})
	}

	config.DB.Preload("HousePerca").First(&infohouse, id)
	return c.JSON(http.StatusAccepted, respon.BaseRespon{
		Code:    http.StatusAccepted,
		Message: "Successful update data",
		Data:    &infohouse,
	})
}

// Delete Info House Perca
func DeleteInfoHouse(c echo.Context) error {
	var infohouse model.InfoHouse

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, respon.BaseRespon{
			Code:    http.StatusUnprocessableEntity,
			Message: "Path parameter invalid",
			Data:    nil,
		})
	}

	result := config.DB.First(&infohouse, id)
	if result.Error != nil {
		return c.JSON(http.StatusGone, respon.BaseRespon{
			Code:    http.StatusGone,
			Message: "Data not Found",
			Data:    nil,
		})
	}

	result = config.DB.Where("id = ?", id).Unscoped().Delete(&infohouse)
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
		Data:    &infohouse,
	})
}

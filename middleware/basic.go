package middleware

import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/model"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(adminname, password string, c echo.Context) (bool, error) {
	var admin model.Admins

	err := config.DB.Where("email = ? AND password = ?", adminname, password).First(&admin).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

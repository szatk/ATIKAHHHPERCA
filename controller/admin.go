package controller

//import packages
import (
	"ATIKAH-PERCA/config"
	"ATIKAH-PERCA/middleware"
	"ATIKAH-PERCA/model"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

//fungsi memanggil seluruh data admin
func GetAdminsController(c echo.Context) error {
	var admins []model.Admins

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all admins",
		"admins":  admins,
	})
}

//Fungsi get admin by ID
func GetAdminsById(c echo.Context) error {
	var admins model.Admins
	//  Fungsi strconv.Atoi() Fungsi ini digunakan untuk konversi data dari tipe string ke int. strconv.Atoi()
	//menghasilkan 2 buah nilai , yaitu hasil konversi dan error (jika konversi sukses, maka error berisi nil)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := config.DB.Where("id= ?", id).Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get admins ",
		"admins":  admins,
	})
}

//fungsi create new admins
func CreateAdminController(e echo.Context) error {
	admin := model.Admins{}
	e.Bind(&admin)

	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success menambahkan Admins",
		"admin":   admin,
	})
}

func UpdateAdmin(e echo.Context) error {
	admin := model.Admins{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&admin)
	admin.ID = id

	if err := config.DB.Where("id= ?", admin.ID).Updates(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Admins",
		"admin":   admin,
	})
}

//Fungsi hapus data admin
func DeleteAdmin(e echo.Context) error {
	var admin model.Admins
	id, _ := strconv.Atoi(e.Param("id"))
	config.DB.Where("id = ?", id).Delete(&admin)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data Berhasil Dihapus",
	})
}

func LoginAdminController(e echo.Context) error {
	admin := model.Admins{}
	e.Bind(&admin)

	err := config.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}
	token, err := middleware.CreateToken(admin.ID, admin.Name)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	//adminRes := db.AdminRes{admin.ID, admin.Name, admin.Email, token}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"token":   token,
		//"user":    adminRes,
	})
}

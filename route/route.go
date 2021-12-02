package route

import (
	"ATIKAH-PERCA/constants"
	"ATIKAH-PERCA/controller"
	"ATIKAH-PERCA/middleware"
	m "ATIKAH-PERCA/middleware"

	"github.com/labstack/echo/v4"
	midd "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	//CRUD ADMIN
	e.GET("/admins", controller.GetAdminsController)
	e.GET("/admins/:id", controller.GetAdminsById)
	m.LogMiddleware(e)
	e.POST("/admins", controller.CreateAdminController)
	e.DELETE("/admins/:id", controller.DeleteAdmin)
	e.PUT("/admins/:id", controller.UpdateAdmin)
	e.POST("/login", controller.LoginAdminController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(midd.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.GET("/admins", controller.GetAdminsController)

	eJwt := e.Group("/jwt")
	eJwt.Use(midd.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/admins", controller.GetAdminsController)

	//CRUD HOUSEPERCA
	e.GET("/houseperca", controller.GetAllHousePerca)
	e.GET("/houseperca/:id", controller.GetHousePercaById)
	e.POST("/houseperca", controller.CreateHousePerca)
	e.DELETE("/houseperca/:id", controller.DeleteHousePerca)
	e.PUT("/houseperca/:id", controller.UpdateHousePerca)

	//CRUD INFO HOUSE
	e.GET("/info", controller.GetAllInfoHouse)
	e.GET("/info/:id", controller.GetInfoHouseById)
	e.POST("/info", controller.CreateInfoHouse)
	e.DELETE("/info/:id", controller.DeleteInfoHouse)
	e.PUT("/info/:id", controller.UpdateInfoHouse)

	//CRUD JENIS PERCA
	e.GET("/jenisperca", controller.GetAllJenisPerca)
	e.GET("/jenisperca/:id", controller.GetJenisPercaById)
	e.POST("/jenisperca", controller.CreateJenisPerca)
	e.DELETE("/jenisperca/:id", controller.DeleteJenisPerca)
	e.PUT("/jenisperca/:id", controller.UpdateJenisPerca)

	//CRUD PEKERJA PERCA
	e.GET("/pekerjaperca", controller.GetAllPekerjaPerca)
	e.GET("/pekerjaperca/:id", controller.GetPekerjaPecraById)
	e.POST("/pekerjaperca", controller.CreatePekerjaPerca)
	e.DELETE("/pekerjaperca/:id", controller.DeletePekerjaPerca)
	e.PUT("/pekerjaperca/:id", controller.UpdatePekerjaPerca)

	//CRUD PERCA BAJU
	e.GET("/percabaju", controller.GetAllPercaBaju)
	e.GET("/percabaju/:id", controller.GetPercaBajuById)
	e.POST("/percabaju", controller.CreatePercaBaju)
	e.DELETE("/percabaju/:id", controller.DeletePercaBaju)
	e.PUT("/percabaju/:id", controller.UpdatePercaBaju)

	//CRUD PERCA SEPATU
	e.GET("/percasepatu", controller.GetAllPercaSepatu)
	e.GET("/percasepatu/:id", controller.GetPercaSepatuById)
	e.POST("/percasepatu", controller.CreatePercaSepatu)
	e.DELETE("/percasepatu/:id", controller.DeletePercaSepatu)
	e.PUT("/percasepatu/:id", controller.UpdatePercaSepatu)

	//CRUD TUTORIAL PERCA
	e.GET("/tutor", controller.GetAllTutorial)
	e.GET("/tutor/:id", controller.GetTutorialById)
	e.POST("/tutor", controller.CreateTutorial)
	e.DELETE("/tutor/:id", controller.DeleteTutorial)
	e.PUT("/tutor/:id", controller.UpdateTutorial)

	return e
}

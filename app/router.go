package app

import (
	"alurkerjaService/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(jenisDataController controller.JenisDataController, dataKinerjaPemdaController controller.DataKinerjaPemdaController, dataKinerjaOpdController controller.DataKinerjaOpdController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/jenisdata", jenisDataController.Create)
	e.PUT("/jenisdata/:id", jenisDataController.Update)
	e.DELETE("/jenisdata/:id", jenisDataController.Delete)
	e.GET("/jenisdata/:id", jenisDataController.FindById)
	e.GET("/jenisdata", jenisDataController.FindAll)

	e.POST("/datakinerjapemda", dataKinerjaPemdaController.Create)
	e.PUT("/datakinerjapemda/:id", dataKinerjaPemdaController.Update)
	e.DELETE("/datakinerjapemda/:id", dataKinerjaPemdaController.Delete)
	e.GET("/datakinerjapemda/detail/:id", dataKinerjaPemdaController.FindById)
	e.GET("/datakinerjapemda/list/", dataKinerjaPemdaController.FindAll)

	e.POST("/datakinerjaopd", dataKinerjaOpdController.Create)
	e.PUT("/datakinerjaopd/:id", dataKinerjaOpdController.Update)
	e.DELETE("/datakinerjaopd/:id", dataKinerjaOpdController.Delete)
	e.GET("/datakinerjaopd/detail/:id", dataKinerjaOpdController.FindById)
	e.GET("/datakinerjaopd/list/:kode_opd", dataKinerjaOpdController.FindAll)

	e.POST("/jenisdataopd", jenisDataController.CreateOpd)
	e.PUT("/jenisdataopd/:id", jenisDataController.UpdateOpd)
	e.DELETE("/jenisdataopd/:id", jenisDataController.DeleteOpd)
	e.GET("/jenisdataopd/detail/:id", jenisDataController.FindByIdOpd)
	e.GET("/jenisdataopd/list/:kode_opd", jenisDataController.FindAllOpd)

	return e
}

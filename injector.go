//go:build wireinject
// +build wireinject

package main

import (
	"alurkerjaService/app"

	"alurkerjaService/controller"
	"alurkerjaService/repository"
	"alurkerjaService/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var jenisDataSet = wire.NewSet(
	repository.NewJenisDataRepositoryImpl,
	wire.Bind(new(repository.JenisDataRepository), new(*repository.JenisDataRepositoryImpl)),
	service.NewJenisDataServiceImpl,
	wire.Bind(new(service.JenisDataService), new(*service.JenisDataServiceImpl)),
	controller.NewJenisDataControllerImpl,
	wire.Bind(new(controller.JenisDataController), new(*controller.JenisDataControllerImpl)),
)

var dataKinerjaPemdaSet = wire.NewSet(
	repository.NewDataKinerjaPemdaRepositoryImpl,
	wire.Bind(new(repository.DataKinerjaPemdaRepository), new(*repository.DataKinerjaPemdaRepositoryImpl)),
	service.NewDataKinerjaPemdaServiceImpl,
	wire.Bind(new(service.DataKinerjaPemdaService), new(*service.DataKinerjaPemdaServiceImpl)),
	controller.NewDataKinerjaPemdaControllerImpl,
	wire.Bind(new(controller.DataKinerjaPemdaController), new(*controller.DataKinerjaPemdaControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		jenisDataSet,
		dataKinerjaPemdaSet,
		app.NewRouter,
	)
	return nil
}

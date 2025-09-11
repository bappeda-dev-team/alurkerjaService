package controller

import (
	"alurkerjaService/model/web"
	"alurkerjaService/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DataKinerjaPemdaControllerImpl struct {
	DataKinerjaPemdaService service.DataKinerjaPemdaService
}

func NewDataKinerjaPemdaControllerImpl(dataKinerjaPemdaService service.DataKinerjaPemdaService) *DataKinerjaPemdaControllerImpl {
	return &DataKinerjaPemdaControllerImpl{
		DataKinerjaPemdaService: dataKinerjaPemdaService,
	}
}

// @Summary Create Data Kinerja Pemda
// @Description Create new Data Kinerja Pemda with targets
// @Tags Data Kinerja Pemda
// @Accept json
// @Produce json
// @Param data body web.DataKinerjaPemdaCreateRequest true "Data Kinerja Pemda Create Request"
// @Success 201 {object} web.WebResponse{data=web.DataKinerjaPemdaResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Router /datakinerjapemda [post]
func (controller *DataKinerjaPemdaControllerImpl) Create(c echo.Context) error {
	dataKinerjaPemdaCreateRequest := web.DataKinerjaPemdaCreateRequest{}
	err := c.Bind(&dataKinerjaPemdaCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaPemdaResponse := controller.DataKinerjaPemdaService.Create(c.Request().Context(), dataKinerjaPemdaCreateRequest)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   dataKinerjaPemdaResponse,
	})
}

// @Summary Update Data Kinerja Pemda
// @Description Update existing Data Kinerja Pemda by ID
// @Tags Data Kinerja Pemda
// @Accept json
// @Produce json
// @Param id path int true "Data Kinerja Pemda ID"
// @Param data body web.DataKinerjaPemdaUpdateRequest true "Data Kinerja Pemda Update Request"
// @Success 200 {object} web.WebResponse{data=web.DataKinerjaPemdaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Router /datakinerjapemda/{id} [put]
func (controller *DataKinerjaPemdaControllerImpl) Update(c echo.Context) error {
	dataKinerjaPemdaUpdateRequest := web.DataKinerjaPemdaUpdateRequest{}
	err := c.Bind(&dataKinerjaPemdaUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaPemdaUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaPemdaResponse := controller.DataKinerjaPemdaService.Update(c.Request().Context(), dataKinerjaPemdaUpdateRequest)

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   dataKinerjaPemdaResponse,
	})
}

// @Summary Delete Data Kinerja Pemda
// @Description Delete Data Kinerja Pemda by ID
// @Tags Data Kinerja Pemda
// @Accept json
// @Produce json
// @Param id path int true "Data Kinerja Pemda ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Router /datakinerjapemda/{id} [delete]
func (controller *DataKinerjaPemdaControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	controller.DataKinerjaPemdaService.Delete(c.Request().Context(), id)

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

// @Summary Get Data Kinerja Pemda by ID
// @Description Get Data Kinerja Pemda detail by ID including targets
// @Tags Data Kinerja Pemda
// @Accept json
// @Produce json
// @Param id path int true "Data Kinerja Pemda ID"
// @Success 200 {object} web.WebResponse{data=web.DataKinerjaPemdaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Router /datakinerjapemda/detail/{id} [get]
func (controller *DataKinerjaPemdaControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaPemdaResponse := controller.DataKinerjaPemdaService.FindById(c.Request().Context(), id)

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   dataKinerjaPemdaResponse,
	})
}

// @Summary List Data Kinerja Pemda
// @Description Get list of Data Kinerja Pemda filtered by jenis data
// @Tags Data Kinerja Pemda
// @Accept json
// @Produce json
// @Param jenis_data_id path int true "Jenis Data ID"
// @Success 200 {object} web.WebResponse{data=[]web.DataKinerjaPemdaResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Router /datakinerjapemda/list/{jenis_data_id} [get]
func (controller *DataKinerjaPemdaControllerImpl) FindAll(c echo.Context) error {
	jenisDataId, err := strconv.Atoi(c.Param("jenis_data_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   "Jenis data id tidak boleh kosong",
		})
	}
	dataKinerjaPemdaResponses := controller.DataKinerjaPemdaService.FindAll(c.Request().Context(), jenisDataId)

	if len(dataKinerjaPemdaResponses) == 0 {
		return c.JSON(http.StatusOK, web.WebResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   []web.DataKinerjaPemdaResponse{},
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   dataKinerjaPemdaResponses,
	})
}

package controller

import (
	"alurkerjaService/model/web"
	"alurkerjaService/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DataKinerjaOpdControllerImpl struct {
	DataKinerjaOpdService service.DataKinerjaOpdService
}

func NewDataKinerjaOpdControllerImpl(dataKinerjaOpdService service.DataKinerjaOpdService) *DataKinerjaOpdControllerImpl {
	return &DataKinerjaOpdControllerImpl{
		DataKinerjaOpdService: dataKinerjaOpdService,
	}
}

// @Summary Create Data Kinerja OPD
// @Description Create new Data Kinerja OPD with targets
// @Tags Data Kinerja OPD
// @Accept json
// @Produce json
// @Param data body web.DataKinerjaOpdCreateRequest true "Data Kinerja OPD Create Request"
// @Success 201 {object} web.WebResponse{data=web.DataKinerjaOpdResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Router /datakinerjaopd [post]
func (controller *DataKinerjaOpdControllerImpl) Create(c echo.Context) error {
	dataKinerjaOpdCreateRequest := web.DataKinerjaOpdCreateRequest{}
	err := c.Bind(&dataKinerjaOpdCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaOpdResponse := controller.DataKinerjaOpdService.Create(c.Request().Context(), dataKinerjaOpdCreateRequest)

	return c.JSON(http.StatusCreated, web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   dataKinerjaOpdResponse,
	})
}

// @Summary Update Data Kinerja OPD
// @Description Update existing Data Kinerja OPD by ID
// @Tags Data Kinerja OPD
// @Accept json
// @Produce json
// @Param id path int true "Data Kinerja OPD ID"
// @Param data body web.DataKinerjaOpdUpdateRequest true "Data Kinerja OPD Update Request"
// @Success 200 {object} web.WebResponse{data=web.DataKinerjaOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Router /datakinerjaopd/{id} [put]
func (controller *DataKinerjaOpdControllerImpl) Update(c echo.Context) error {
	dataKinerjaOpdUpdateRequest := web.DataKinerjaOpdUpdateRequest{}
	err := c.Bind(&dataKinerjaOpdUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaOpdUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaOpdResponse := controller.DataKinerjaOpdService.Update(c.Request().Context(), dataKinerjaOpdUpdateRequest)

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   dataKinerjaOpdResponse,
	})
}

// @Summary Delete Data Kinerja OPD
// @Description Delete Data Kinerja OPD by ID
// @Tags Data Kinerja OPD
// @Accept json
// @Produce json
// @Param id path int true "Data Kinerja OPD ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Router /datakinerjaopd/{id} [delete]
func (controller *DataKinerjaOpdControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	controller.DataKinerjaOpdService.Delete(c.Request().Context(), id)

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

// @Summary Get Data Kinerja OPD by ID
// @Description Get Data Kinerja OPD detail by ID including targets
// @Tags Data Kinerja OPD
// @Accept json
// @Produce json
// @Param id path int true "Data Kinerja OPD ID"
// @Success 200 {object} web.WebResponse{data=web.DataKinerjaOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Router /datakinerjaopd/detail/{id} [get]
func (controller *DataKinerjaOpdControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaOpdResponse, err := controller.DataKinerjaOpdService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   dataKinerjaOpdResponse,
	})
}

// @Summary List Data Kinerja OPD
// @Description Get list of Data Kinerja OPD filtered by jenis data
// @Tags Data Kinerja OPD
// @Accept json
// @Produce json
// @Param kode_opd path string true "Kode OPD"
// @Param jenis_data_id path int true "Jenis Data ID"
// @Success 200 {object} web.WebResponse{data=[]web.DataKinerjaOpdResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Router /datakinerjaopd/list/{kode_opd}/{jenis_data_id} [get]
func (controller *DataKinerjaOpdControllerImpl) FindAll(c echo.Context) error {
	kodeOpd := c.Param("kode_opd")
	if kodeOpd == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   "Kode OPD tidak boleh kosong",
		})
	}
	jenisDataId, err := strconv.Atoi(c.Param("jenis_data_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	dataKinerjaOpdResponses := controller.DataKinerjaOpdService.FindAll(c.Request().Context(), kodeOpd, jenisDataId)

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   dataKinerjaOpdResponses,
	})
}

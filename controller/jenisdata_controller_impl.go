package controller

import (
	"alurkerjaService/model/web"
	"alurkerjaService/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JenisDataControllerImpl struct {
	JenisDataService service.JenisDataService
}

func NewJenisDataControllerImpl(jenisDataService service.JenisDataService) *JenisDataControllerImpl {
	return &JenisDataControllerImpl{
		JenisDataService: jenisDataService,
	}
}

// @Summary Create Jenis Data
// @Description Create new Jenis Data
// @Tags Jenis Data
// @Accept json
// @Produce json
// @Param data body web.JenisDataCreateRequest true "Jenis Data Create Request"
// @Success 201 {object} web.WebResponse{data=web.JenisDataResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jenisdata [post]
func (controller *JenisDataControllerImpl) Create(c echo.Context) error {
	jenisDataCreateRequest := web.JenisDataCreateRequest{}
	err := c.Bind(&jenisDataCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jenisDataResponse, err := controller.JenisDataService.Create(c.Request().Context(), jenisDataCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   jenisDataResponse,
	})
}

// @Summary Update Jenis Data
// @Description Update existing Jenis Data by ID
// @Tags Jenis Data
// @Accept json
// @Produce json
// @Param id path int true "Jenis Data ID"
// @Param data body web.JenisDataUpdateRequest true "Jenis Data Update Request"
// @Success 200 {object} web.WebResponse{data=web.JenisDataResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jenisdata/{id} [put]
func (controller *JenisDataControllerImpl) Update(c echo.Context) error {
	jenisDataUpdateRequest := web.JenisDataUpdateRequest{}
	err := c.Bind(&jenisDataUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jenisDataUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jenisDataResponse, err := controller.JenisDataService.Update(c.Request().Context(), jenisDataUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   jenisDataResponse,
	})
}

// @Summary Update Jenis Data
// @Description Update existing Jenis Data by ID
// @Tags Jenis Data
// @Accept json
// @Produce json
// @Param id path int true "Jenis Data ID"
// @Param data body web.JenisDataUpdateRequest true "Jenis Data Update Request"
// @Success 200 {object} web.WebResponse{data=web.JenisDataResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jenisdata/{id} [put]
func (controller *JenisDataControllerImpl) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	err = controller.JenisDataService.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	})
}

// @Summary Get Jenis Data by ID
// @Description Get Jenis Data detail by ID
// @Tags Jenis Data
// @Accept json
// @Produce json
// @Param id path int true "Jenis Data ID"
// @Success 200 {object} web.WebResponse{data=web.JenisDataResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jenisdata/{id} [get]
func (controller *JenisDataControllerImpl) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
		})
	}

	jenisDataResponse, err := controller.JenisDataService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   jenisDataResponse,
	})
}

// @Summary List All Jenis Data
// @Description Get list of all Jenis Data
// @Tags Jenis Data
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.JenisDataResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /jenisdata [get]
func (controller *JenisDataControllerImpl) FindAll(c echo.Context) error {
	jenisDataResponses, err := controller.JenisDataService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   jenisDataResponses,
	})
}

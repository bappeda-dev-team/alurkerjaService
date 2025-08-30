package web

type JenisDataCreateRequest struct {
	JenisData string `json:"jenis_data" validate:"required"`
}

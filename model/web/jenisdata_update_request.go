package web

type JenisDataUpdateRequest struct {
	Id        int    `json:"id" validate:"required"`
	JenisData string `json:"jenis_data" validate:"required"`
}

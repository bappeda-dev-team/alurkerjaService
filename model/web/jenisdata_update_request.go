package web

type JenisDataUpdateRequest struct {
	Id        int    `json:"id" validate:"required"`
	JenisData string `json:"jenis_data" validate:"required"`
}

type JenisDataOpdUpdateRequest struct {
	Id        int    `json:"id" validate:"required"`
	KodeOpd   string `json:"kode_opd" validate:"required"`
	NamaOpd   string `json:"nama_opd" validate:"required"`
	JenisData string `json:"jenis_data" validate:"required"`
}

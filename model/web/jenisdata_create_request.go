package web

type JenisDataCreateRequest struct {
	JenisData string `json:"jenis_data" validate:"required"`
}

type JenisDataOpdCreateRequest struct {
	KodeOpd   string `json:"kode_opd" validate:"required"`
	NamaOpd   string `json:"nama_opd" validate:"required"`
	JenisData string `json:"jenis_data" validate:"required"`
}

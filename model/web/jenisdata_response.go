package web

type JenisDataResponse struct {
	Id        int    `json:"id,omitempty"`
	JenisData string `json:"jenis_data"`
}

type JenisDataOpdResponse struct {
	Id        int    `json:"id,omitempty"`
	KodeOpd   string `json:"kode_opd"`
	NamaOpd   string `json:"nama_opd"`
	JenisData string `json:"jenis_data"`
}

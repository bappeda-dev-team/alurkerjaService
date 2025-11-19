package web

type JenisDataPemdaResponse struct {
	Id               int                        `json:"id"`
	JenisData        string                     `json:"jenis_data"`
	DataKinerjaPemda []DataKinerjaPemdaResponse `json:"data_kinerja"`
}

type DataKinerjaPemdaResponse struct {
	Id                   int              `json:"id"`
	JenisDataId          int              `json:"jenis_data_id"`
	JenisData            string           `json:"jenis_data"`
	NamaData             string           `json:"nama_data"`
	RumusPerhitungan     string           `json:"rumus_perhitungan"`
	SumberData           string           `json:"sumber_data"`
	InstansiProdusenData string           `json:"instansi_produsen_data"`
	Target               []TargetResponse `json:"target"`
	Keterangan           string           `json:"keterangan"`
}

type TargetResponse struct {
	Id     int    `json:"id"`
	Target string `json:"target"`
	Satuan string `json:"satuan"`
	Tahun  string `json:"tahun"`
}

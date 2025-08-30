package web

type DataKinerjaPemdaUpdateRequest struct {
	Id                   int                   `validate:"required" json:"id"`
	JenisDataId          int                   `validate:"required" json:"jenis_data_id"`
	NamaData             string                `validate:"required" json:"nama_data"`
	RumusPerhitungan     string                `validate:"required" json:"rumus_perhitungan"`
	SumberData           string                `validate:"required" json:"sumber_data"`
	InstansiProdusenData string                `validate:"required" json:"instansi_produsen_data"`
	Target               []TargetUpdateRequest `validate:"required" json:"target"`
	Keterangan           string                `json:"keterangan"`
}

type TargetUpdateRequest struct {
	Id     int    `json:"id"`
	Target string `validate:"required" json:"target"`
	Satuan string `validate:"required" json:"satuan"`
	Tahun  string `validate:"required" json:"tahun"`
}

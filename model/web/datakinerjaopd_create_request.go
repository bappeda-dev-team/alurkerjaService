package web

type DataKinerjaOpdCreateRequest struct {
	JenisDataId          int                   `validate:"required" json:"jenis_data_id"`
	KodeOpd              string                `validate:"required" json:"kode_opd"`
	NamaData             string                `validate:"required" json:"nama_data"`
	RumusPerhitungan     string                `validate:"required" json:"rumus_perhitungan"`
	SumberData           string                `validate:"required" json:"sumber_data"`
	InstansiProdusenData string                `validate:"required" json:"instansi_produsen_data"`
	Target               []TargetCreateRequest `validate:"required" json:"target"`
	Keterangan           string                `json:"keterangan"`
}

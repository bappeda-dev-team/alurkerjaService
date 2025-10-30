package web

type DataKinerjaOpdResponse struct {
	Id                   int              `json:"id"`
	JenisDataId          int              `json:"jenis_data_id"`
	JenisData            string           `json:"jenis_data"`
	KodeOpd              string           `json:"kode_opd"`
	NamaOpd              string           `json:"nama_opd"`
	NamaData             string           `json:"nama_data"`
	RumusPerhitungan     string           `json:"rumus_perhitungan"`
	SumberData           string           `json:"sumber_data"`
	InstansiProdusenData string           `json:"instansi_produsen_data"`
	Target               []TargetResponse `json:"target"`
	Keterangan           string           `json:"keterangan"`
}

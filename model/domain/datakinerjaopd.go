package domain

type DataKinerjaOpd struct {
	Id                   int
	JenisDataId          int
	JenisData            string
	KodeOpd              string
	NamaOpd              string
	NamaData             string
	RumusPerhitungan     string
	SumberData           string
	InstansiProdusenData string
	Target               []TargetOpd
	Keterangan           string
}

type TargetOpd struct {
	Id            int
	DataKinerjaId int
	Target        string
	Satuan        string
	Tahun         string
}

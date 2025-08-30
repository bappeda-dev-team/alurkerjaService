package domain

type DataKinerjaPemda struct {
	Id                   int
	JenisDataId          int
	JenisData            string
	NamaData             string
	RumusPerhitungan     string
	SumberData           string
	InstansiProdusenData string
	Target               []Target
	Keterangan           string
}

type Target struct {
	Id            int
	DataKinerjaId int
	Target        string
	Satuan        string
	Tahun         string
}

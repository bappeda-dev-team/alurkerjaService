package service

import (
	"alurkerjaService/helper"
	"alurkerjaService/model/domain"
	"alurkerjaService/model/web"
	"alurkerjaService/repository"
	"context"
	"database/sql"
	"log"

	"github.com/go-playground/validator/v10"
)

type DataKinerjaPemdaServiceImpl struct {
	Repository          repository.DataKinerjaPemdaRepository
	DB                  *sql.DB
	Validate            *validator.Validate
	JenisDataRepository repository.JenisDataRepository
}

func NewDataKinerjaPemdaServiceImpl(repository repository.DataKinerjaPemdaRepository, db *sql.DB, validate *validator.Validate, jenisDataRepository repository.JenisDataRepository) *DataKinerjaPemdaServiceImpl {
	return &DataKinerjaPemdaServiceImpl{
		Repository:          repository,
		DB:                  db,
		Validate:            validate,
		JenisDataRepository: jenisDataRepository,
	}
}

func (service *DataKinerjaPemdaServiceImpl) Create(ctx context.Context, request web.DataKinerjaPemdaCreateRequest) web.DataKinerjaPemdaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Log request data untuk debugging
	log.Printf("Creating Data Kinerja with JenisDataId: %d", request.JenisDataId)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Validasi jenis_data_id
	jenisData, err := service.JenisDataRepository.FindById(ctx, tx, request.JenisDataId)
	if err != nil {
		log.Printf("Error finding JenisData with ID %d: %v", request.JenisDataId, err)
		return web.DataKinerjaPemdaResponse{} // Return empty response jika data tidak ditemukan
	}

	dataKinerjaPemda := domain.DataKinerjaPemda{
		JenisDataId:          request.JenisDataId,
		JenisData:            jenisData.JenisData,
		NamaData:             request.NamaData,
		RumusPerhitungan:     helper.EmptyStringIfNull(request.RumusPerhitungan),
		SumberData:           helper.EmptyStringIfNull(request.SumberData),
		InstansiProdusenData: helper.EmptyStringIfNull(request.InstansiProdusenData),
		Keterangan:           helper.EmptyStringIfNull(request.Keterangan),
		Target:               make([]domain.Target, len(request.Target)),
	}

	// Convert target requests to domain
	for i, targetReq := range request.Target {
		dataKinerjaPemda.Target[i] = domain.Target{
			Target: targetReq.Target,
			Satuan: targetReq.Satuan,
			Tahun:  targetReq.Tahun,
		}
	}

	result, err := service.Repository.Create(ctx, tx, dataKinerjaPemda)
	helper.PanicIfError(err)

	return helper.ToDataKinerjaPemdaResponse(result)
}

func (service *DataKinerjaPemdaServiceImpl) Update(ctx context.Context, request web.DataKinerjaPemdaUpdateRequest) web.DataKinerjaPemdaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Cek data exists terlebih dahulu
	existing, err := service.Repository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	jenisData, err := service.JenisDataRepository.FindById(ctx, tx, existing.JenisDataId)
	helper.PanicIfError(err)

	// Log data existing
	log.Printf("Updating Data Kinerja - ID: %d", existing.Id)
	log.Printf("Current Data: %+v", existing)

	dataKinerjaPemda := domain.DataKinerjaPemda{
		Id:                   request.Id,
		JenisDataId:          existing.JenisDataId,
		JenisData:            jenisData.JenisData,
		NamaData:             request.NamaData,
		RumusPerhitungan:     helper.EmptyStringIfNull(request.RumusPerhitungan),
		SumberData:           helper.EmptyStringIfNull(request.SumberData),
		InstansiProdusenData: helper.EmptyStringIfNull(request.InstansiProdusenData),
		Keterangan:           helper.EmptyStringIfNull(request.Keterangan),
		Target:               make([]domain.Target, len(request.Target)),
	}

	// Convert target requests to domain
	for i, targetReq := range request.Target {
		dataKinerjaPemda.Target[i] = domain.Target{
			Id:            targetReq.Id,
			DataKinerjaId: request.Id,
			Target:        targetReq.Target,
			Satuan:        targetReq.Satuan,
			Tahun:         targetReq.Tahun,
		}
	}

	// Log data yang akan diupdate
	log.Printf("New Data: %+v", dataKinerjaPemda)

	result, err := service.Repository.Update(ctx, tx, dataKinerjaPemda)
	helper.PanicIfError(err)

	return helper.ToDataKinerjaPemdaResponse(result)
}

func (service *DataKinerjaPemdaServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = service.Repository.Delete(ctx, tx, id)
	helper.PanicIfError(err)
}

func (service *DataKinerjaPemdaServiceImpl) FindById(ctx context.Context, id int) web.DataKinerjaPemdaResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := service.Repository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	// Debug log
	log.Printf("Found Data Kinerja ID: %d", result.Id)
	log.Printf("- Nama Data: %s", result.NamaData)
	log.Printf("- Jumlah Target: %d", len(result.Target))
	for _, t := range result.Target {
		log.Printf("  * Target: %s, Tahun: %s, Satuan: %s", t.Target, t.Tahun, t.Satuan)
	}

	return helper.ToDataKinerjaPemdaResponse(result)
}

func (service *DataKinerjaPemdaServiceImpl) FindAll(ctx context.Context, jenisDataId int) []web.JenisDataPemdaResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	results, err := service.Repository.FindAll(ctx, tx, jenisDataId)
	helper.PanicIfError(err)

	// Jika tidak ada data sama sekali, kembalikan array kosong
	if len(results) == 0 {
		return []web.JenisDataPemdaResponse{}
	}

	// Grouping berdasarkan jenis_data_id
	jenisDataMap := make(map[int]*web.JenisDataPemdaResponse)

	for _, result := range results {
		// Jika jenis data belum ada di map, buat yang baru
		if _, exists := jenisDataMap[result.JenisDataId]; !exists {
			jenisDataMap[result.JenisDataId] = &web.JenisDataPemdaResponse{
				Id:               result.JenisDataId,
				JenisData:        result.JenisData,
				DataKinerjaPemda: []web.DataKinerjaPemdaResponse{},
			}
		}

		// Tambahkan data kinerja HANYA jika Id > 0 (ada data kinerja)
		if result.Id > 0 {
			jenisDataMap[result.JenisDataId].DataKinerjaPemda = append(
				jenisDataMap[result.JenisDataId].DataKinerjaPemda,
				helper.ToDataKinerjaPemdaResponse(result),
			)
		}
	}

	// Convert map ke slice
	var responses []web.JenisDataPemdaResponse
	for _, jenisDataResp := range jenisDataMap {
		responses = append(responses, *jenisDataResp)
	}

	// Log untuk debugging
	for _, resp := range responses {
		log.Printf("Jenis Data ID: %d, Nama: %s", resp.Id, resp.JenisData)
		log.Printf("Jumlah Data Kinerja: %d", len(resp.DataKinerjaPemda))
		if len(resp.DataKinerjaPemda) == 0 {
			log.Printf("  (Tidak ada data kinerja untuk jenis data ini)")
		}
		for _, dk := range resp.DataKinerjaPemda {
			log.Printf("  - Data Kinerja ID: %d, Nama: %s", dk.Id, dk.NamaData)
			log.Printf("    Jumlah Target: %d", len(dk.Target))
		}
	}

	return responses
}

package service

import (
	"alurkerjaService/helper"
	"alurkerjaService/model/domain"
	"alurkerjaService/model/web"
	"alurkerjaService/repository"
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

type DataKinerjaOpdServiceImpl struct {
	Repository          repository.DataKinerjaOpdRepository
	DB                  *sql.DB
	Validate            *validator.Validate
	JenisDataRepository repository.JenisDataRepository
}

func NewDataKinerjaOpdServiceImpl(
	repo repository.DataKinerjaOpdRepository,
	db *sql.DB,
	validate *validator.Validate,
	jenisDataRepository repository.JenisDataRepository,
) *DataKinerjaOpdServiceImpl {
	return &DataKinerjaOpdServiceImpl{
		Repository:          repo,
		DB:                  db,
		Validate:            validate,
		JenisDataRepository: jenisDataRepository,
	}
}

func (s *DataKinerjaOpdServiceImpl) Create(ctx context.Context, request web.DataKinerjaOpdCreateRequest) (web.DataKinerjaOpdResponse, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		log.Printf("Validation error: %v", err)
		return web.DataKinerjaOpdResponse{}, errors.New("validasi gagal: " + err.Error())
	}

	log.Printf("Creating Data Kinerja OPD with JenisDataId: %d, KodeOpd: %s", request.JenisDataId, request.KodeOpd)

	tx, err := s.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return web.DataKinerjaOpdResponse{}, errors.New("gagal memulai transaksi database")
	}
	defer helper.CommitOrRollback(tx)

	// Cari jenis data OPD berdasarkan ID
	jenisDataOpd, err := s.JenisDataRepository.FindByIdOpd(ctx, tx, request.JenisDataId)
	if err != nil {
		log.Printf("Error finding JenisDataOpd with ID %d: %v", request.JenisDataId, err)
		return web.DataKinerjaOpdResponse{}, errors.New("jenis data OPD tidak ditemukan")
	}

	// Validasi: Kode OPD harus sama dengan yang ada di jenis data OPD
	if jenisDataOpd.KodeOpd != request.KodeOpd {
		log.Printf("Validasi Gagal - KodeOpd tidak sesuai!")
		log.Printf("- KodeOpd di Request: %s", request.KodeOpd)
		log.Printf("- KodeOpd di JenisDataOpd (ID:%d): %s", jenisDataOpd.Id, jenisDataOpd.KodeOpd)
		return web.DataKinerjaOpdResponse{}, errors.New("kode OPD tidak sesuai dengan jenis data OPD yang dipilih")
	}

	log.Printf("Validasi Berhasil - KodeOpd sesuai: %s", request.KodeOpd)

	data := domain.DataKinerjaOpd{
		JenisDataId:          request.JenisDataId,
		JenisData:            jenisDataOpd.JenisData,
		KodeOpd:              request.KodeOpd,
		NamaOpd:              jenisDataOpd.NamaOpd,
		NamaData:             request.NamaData,
		RumusPerhitungan:     helper.EmptyStringIfNull(request.RumusPerhitungan),
		SumberData:           helper.EmptyStringIfNull(request.SumberData),
		InstansiProdusenData: helper.EmptyStringIfNull(request.InstansiProdusenData),
		Keterangan:           helper.EmptyStringIfNull(request.Keterangan),
		Target:               make([]domain.TargetOpd, len(request.Target)),
	}

	for i, t := range request.Target {
		data.Target[i] = domain.TargetOpd{
			Target: t.Target,
			Satuan: t.Satuan,
			Tahun:  t.Tahun,
		}
	}

	result, err := s.Repository.Create(ctx, tx, data)
	if err != nil {
		log.Printf("Error creating data: %v", err)
		return web.DataKinerjaOpdResponse{}, errors.New("gagal menyimpan data kinerja OPD")
	}

	log.Printf("Data Kinerja OPD berhasil dibuat dengan ID: %d", result.Id)

	return helper.ToDataKinerjaOpdResponse(result), nil
}

func (s *DataKinerjaOpdServiceImpl) Update(ctx context.Context, request web.DataKinerjaOpdUpdateRequest) (web.DataKinerjaOpdResponse, error) {
	err := s.Validate.Struct(request)
	if err != nil {
		log.Printf("Validation error: %v", err)
		return web.DataKinerjaOpdResponse{}, errors.New("validasi gagal: " + err.Error())
	}

	tx, err := s.DB.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return web.DataKinerjaOpdResponse{}, errors.New("gagal memulai transaksi database")
	}
	defer helper.CommitOrRollback(tx)

	existing, err := s.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		log.Printf("Error finding existing data with ID %d: %v", request.Id, err)
		return web.DataKinerjaOpdResponse{}, errors.New("data kinerja OPD tidak ditemukan")
	}

	log.Printf("Updating Data Kinerja OPD - ID: %d", existing.Id)

	// Cari jenis data OPD berdasarkan ID yang baru
	jenisDataOpd, err := s.JenisDataRepository.FindByIdOpd(ctx, tx, request.JenisDataId)
	if err != nil {
		log.Printf("Error finding JenisDataOpd with ID %d: %v", request.JenisDataId, err)
		return web.DataKinerjaOpdResponse{}, errors.New("jenis data OPD tidak ditemukan")
	}

	// Validasi: Kode OPD harus sama dengan yang ada di jenis data OPD
	if jenisDataOpd.KodeOpd != request.KodeOpd {
		log.Printf("Validasi Gagal - KodeOpd tidak sesuai!")
		log.Printf("- KodeOpd di Request: %s", request.KodeOpd)
		log.Printf("- KodeOpd di JenisDataOpd (ID:%d): %s", jenisDataOpd.Id, jenisDataOpd.KodeOpd)
		return web.DataKinerjaOpdResponse{}, errors.New("kode OPD tidak sesuai dengan jenis data OPD yang dipilih")
	}

	log.Printf("Validasi Berhasil - KodeOpd sesuai: %s", request.KodeOpd)
	log.Printf("Current Data: %+v", existing)

	data := domain.DataKinerjaOpd{
		Id:                   request.Id,
		JenisDataId:          request.JenisDataId,
		JenisData:            jenisDataOpd.JenisData,
		KodeOpd:              request.KodeOpd,
		NamaOpd:              jenisDataOpd.NamaOpd,
		NamaData:             request.NamaData,
		RumusPerhitungan:     helper.EmptyStringIfNull(request.RumusPerhitungan),
		SumberData:           helper.EmptyStringIfNull(request.SumberData),
		InstansiProdusenData: helper.EmptyStringIfNull(request.InstansiProdusenData),
		Keterangan:           helper.EmptyStringIfNull(request.Keterangan),
		Target:               make([]domain.TargetOpd, len(request.Target)),
	}

	for i, t := range request.Target {
		data.Target[i] = domain.TargetOpd{
			Id:            t.Id,
			DataKinerjaId: request.Id,
			Target:        t.Target,
			Satuan:        t.Satuan,
			Tahun:         t.Tahun,
		}
	}

	log.Printf("New Data: %+v", data)

	result, err := s.Repository.Update(ctx, tx, data)
	if err != nil {
		log.Printf("Error updating data: %v", err)
		return web.DataKinerjaOpdResponse{}, errors.New("gagal mengupdate data kinerja OPD")
	}

	log.Printf("Data Kinerja OPD berhasil diupdate dengan ID: %d", result.Id)

	return helper.ToDataKinerjaOpdResponse(result), nil
}

func (s *DataKinerjaOpdServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	err = s.Repository.Delete(ctx, tx, id)
	helper.PanicIfError(err)
}

func (s *DataKinerjaOpdServiceImpl) FindById(ctx context.Context, id int) (web.DataKinerjaOpdResponse, error) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	result, err := s.Repository.FindById(ctx, tx, id)
	if err != nil {
		return web.DataKinerjaOpdResponse{}, errors.New("id tidak ditemukan")
	}

	log.Printf("Found Data Kinerja OPD ID: %d", result.Id)
	log.Printf("- Nama Data: %s", result.NamaData)
	log.Printf("- Jumlah Target: %d", len(result.Target))

	return helper.ToDataKinerjaOpdResponse(result), nil
}

func (s *DataKinerjaOpdServiceImpl) FindAll(ctx context.Context, kodeOpd string, jenisDataId int) []web.DataKinerjaOpdResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	results, err := s.Repository.FindAll(ctx, tx, kodeOpd, jenisDataId)
	helper.PanicIfError(err)

	if len(results) == 0 {
		return []web.DataKinerjaOpdResponse{}
	}

	responses := helper.ToDataKinerjaOpdResponses(results)

	for _, resp := range responses {
		log.Printf("Data Kinerja OPD ID: %d, OPD: %s - %s, Nama: %s", resp.Id, resp.KodeOpd, resp.NamaOpd, resp.NamaData)
		log.Printf("Jumlah Target: %d", len(resp.Target))
	}

	return responses
}

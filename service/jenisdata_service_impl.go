package service

import (
	"alurkerjaService/helper"
	"alurkerjaService/model/domain"
	"alurkerjaService/model/web"
	"alurkerjaService/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/go-playground/validator/v10"
)

type JenisDataServiceImpl struct {
	JenisDataRepository repository.JenisDataRepository
	DB                  *sql.DB
	Validator           *validator.Validate
}

func NewJenisDataServiceImpl(jenisDataRepository repository.JenisDataRepository, db *sql.DB, validator *validator.Validate) *JenisDataServiceImpl {
	return &JenisDataServiceImpl{
		JenisDataRepository: jenisDataRepository,
		DB:                  db,
		Validator:           validator,
	}
}

func (service *JenisDataServiceImpl) Create(ctx context.Context, jenisData web.JenisDataCreateRequest) (web.JenisDataResponse, error) {
	err := service.Validator.Struct(jenisData)
	if err != nil {
		return web.JenisDataResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JenisDataResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataDomain := domain.JenisData{
		JenisData: jenisData.JenisData,
	}

	jenisDataDomain, err = service.JenisDataRepository.Create(ctx, tx, jenisDataDomain)
	if err != nil {
		return web.JenisDataResponse{}, err
	}

	return web.JenisDataResponse{
		JenisData: jenisDataDomain.JenisData,
	}, nil
}

func (service *JenisDataServiceImpl) Update(ctx context.Context, jenisData web.JenisDataUpdateRequest) (web.JenisDataResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JenisDataResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataDomain := domain.JenisData{
		Id:        jenisData.Id,
		JenisData: jenisData.JenisData,
	}

	jenisDataDomain, err = service.JenisDataRepository.Update(ctx, tx, jenisDataDomain)
	if err != nil {
		return web.JenisDataResponse{}, err
	}

	return web.JenisDataResponse{
		Id:        jenisDataDomain.Id,
		JenisData: jenisDataDomain.JenisData,
	}, nil
}

func (service *JenisDataServiceImpl) Delete(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.JenisDataRepository.Delete(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *JenisDataServiceImpl) FindById(ctx context.Context, id int) (web.JenisDataResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JenisDataResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataDomain, err := service.JenisDataRepository.FindById(ctx, tx, id)
	if err != nil {
		return web.JenisDataResponse{}, errors.New("id tidak ditemukan")
	}

	return web.JenisDataResponse{
		Id:        jenisDataDomain.Id,
		JenisData: jenisDataDomain.JenisData,
	}, nil
}

func (service *JenisDataServiceImpl) FindAll(ctx context.Context) ([]web.JenisDataResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.JenisDataResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataDomains, err := service.JenisDataRepository.FindAll(ctx, tx)
	if err != nil {
		return []web.JenisDataResponse{}, err
	}

	return helper.ToJenisDataResponses(jenisDataDomains), nil
}

// jenis data opd
func (service *JenisDataServiceImpl) CreateOpd(ctx context.Context, jenisDataOpd web.JenisDataOpdCreateRequest) (web.JenisDataOpdResponse, error) {
	err := service.Validator.Struct(jenisDataOpd)
	if err != nil {
		return web.JenisDataOpdResponse{}, err
	}

	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JenisDataOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataOpdDomain := domain.JenisDataOpd{
		KodeOpd:   jenisDataOpd.KodeOpd,
		NamaOpd:   jenisDataOpd.NamaOpd,
		JenisData: jenisDataOpd.JenisData,
	}

	jenisDataOpdDomain, err = service.JenisDataRepository.CreateOpd(ctx, tx, jenisDataOpdDomain)
	if err != nil {
		return web.JenisDataOpdResponse{}, err
	}

	return web.JenisDataOpdResponse{
		KodeOpd:   jenisDataOpdDomain.KodeOpd,
		NamaOpd:   jenisDataOpdDomain.NamaOpd,
		JenisData: jenisDataOpdDomain.JenisData,
	}, nil
}

func (service *JenisDataServiceImpl) UpdateOpd(ctx context.Context, jenisDataOpd web.JenisDataOpdUpdateRequest) (web.JenisDataOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JenisDataOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataOpdDomain := domain.JenisDataOpd{
		Id:        jenisDataOpd.Id,
		KodeOpd:   jenisDataOpd.KodeOpd,
		NamaOpd:   jenisDataOpd.NamaOpd,
		JenisData: jenisDataOpd.JenisData,
	}

	jenisDataOpdDomain, err = service.JenisDataRepository.UpdateOpd(ctx, tx, jenisDataOpdDomain)
	if err != nil {
		return web.JenisDataOpdResponse{}, err
	}

	return web.JenisDataOpdResponse{
		Id:        jenisDataOpdDomain.Id,
		KodeOpd:   jenisDataOpdDomain.KodeOpd,
		NamaOpd:   jenisDataOpdDomain.NamaOpd,
		JenisData: jenisDataOpdDomain.JenisData,
	}, nil
}

func (service *JenisDataServiceImpl) DeleteOpd(ctx context.Context, id int) error {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer helper.CommitOrRollback(tx)

	err = service.JenisDataRepository.DeleteOpd(ctx, tx, id)
	if err != nil {
		return err
	}

	return nil
}

func (service *JenisDataServiceImpl) FindByIdOpd(ctx context.Context, id int) (web.JenisDataOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return web.JenisDataOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataOpdDomain, err := service.JenisDataRepository.FindByIdOpd(ctx, tx, id)
	if err != nil {
		return web.JenisDataOpdResponse{}, errors.New("id tidak ditemukan")
	}

	return web.JenisDataOpdResponse{
		Id:        jenisDataOpdDomain.Id,
		KodeOpd:   jenisDataOpdDomain.KodeOpd,
		NamaOpd:   jenisDataOpdDomain.NamaOpd,
		JenisData: jenisDataOpdDomain.JenisData,
	}, nil
}

func (service *JenisDataServiceImpl) FindAllOpd(ctx context.Context, kodeOpd string) ([]web.JenisDataOpdResponse, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return []web.JenisDataOpdResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	jenisDataOpdDomains, err := service.JenisDataRepository.FindAllOpd(ctx, tx, kodeOpd)
	if err != nil {
		return []web.JenisDataOpdResponse{}, err
	}

	var responses []web.JenisDataOpdResponse
	for _, jenisDataOpdDomain := range jenisDataOpdDomains {
		responses = append(responses, web.JenisDataOpdResponse{
			Id:        jenisDataOpdDomain.Id,
			KodeOpd:   jenisDataOpdDomain.KodeOpd,
			NamaOpd:   jenisDataOpdDomain.NamaOpd,
			JenisData: jenisDataOpdDomain.JenisData,
		})
	}
	return responses, nil
}

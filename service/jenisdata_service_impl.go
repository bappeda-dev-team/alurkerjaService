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

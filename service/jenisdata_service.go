package service

import (
	"alurkerjaService/model/web"
	"context"
)

type JenisDataService interface {
	Create(ctx context.Context, jenisData web.JenisDataCreateRequest) (web.JenisDataResponse, error)
	Update(ctx context.Context, jenisData web.JenisDataUpdateRequest) (web.JenisDataResponse, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (web.JenisDataResponse, error)
	FindAll(ctx context.Context) ([]web.JenisDataResponse, error)

	//jenis data opd
	CreateOpd(ctx context.Context, jenisDataOpd web.JenisDataOpdCreateRequest) (web.JenisDataOpdResponse, error)
	UpdateOpd(ctx context.Context, jenisDataOpd web.JenisDataOpdUpdateRequest) (web.JenisDataOpdResponse, error)
	DeleteOpd(ctx context.Context, id int) error
	FindByIdOpd(ctx context.Context, id int) (web.JenisDataOpdResponse, error)
	FindAllOpd(ctx context.Context, kodeOpd string) ([]web.JenisDataOpdResponse, error)
}

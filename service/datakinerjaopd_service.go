package service

import (
	"alurkerjaService/model/web"
	"context"
)

type DataKinerjaOpdService interface {
	Create(ctx context.Context, request web.DataKinerjaOpdCreateRequest) (web.DataKinerjaOpdResponse, error)
	Update(ctx context.Context, request web.DataKinerjaOpdUpdateRequest) (web.DataKinerjaOpdResponse, error)
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) (web.DataKinerjaOpdResponse, error)
	FindAll(ctx context.Context, kodeOpd string, jenisDataId int) []web.DataKinerjaOpdResponse
}

package service

import (
	"alurkerjaService/model/web"
	"context"
)

type DataKinerjaPemdaService interface {
	Create(ctx context.Context, request web.DataKinerjaPemdaCreateRequest) web.DataKinerjaPemdaResponse
	Update(ctx context.Context, request web.DataKinerjaPemdaUpdateRequest) web.DataKinerjaPemdaResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.DataKinerjaPemdaResponse
	FindAll(ctx context.Context, jenisDataId int) []web.JenisDataPemdaResponse
}

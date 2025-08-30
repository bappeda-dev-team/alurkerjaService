package repository

import (
	"alurkerjaService/model/domain"
	"context"
	"database/sql"
)

type DataKinerjaPemdaRepository interface {
	Create(ctx context.Context, tx *sql.Tx, dataKinerjaPemda domain.DataKinerjaPemda) (domain.DataKinerjaPemda, error)
	Update(ctx context.Context, tx *sql.Tx, dataKinerjaPemda domain.DataKinerjaPemda) (domain.DataKinerjaPemda, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.DataKinerjaPemda, error)
	FindAll(ctx context.Context, tx *sql.Tx, jenisDataId int) ([]domain.DataKinerjaPemda, error)
}

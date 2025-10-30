package repository

import (
	"alurkerjaService/model/domain"
	"context"
	"database/sql"
)

type DataKinerjaOpdRepository interface {
	Create(ctx context.Context, tx *sql.Tx, dataKinerjaOpd domain.DataKinerjaOpd) (domain.DataKinerjaOpd, error)
	Update(ctx context.Context, tx *sql.Tx, dataKinerjaOpd domain.DataKinerjaOpd) (domain.DataKinerjaOpd, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.DataKinerjaOpd, error)
	FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, jenisDataId int) ([]domain.DataKinerjaOpd, error)
}

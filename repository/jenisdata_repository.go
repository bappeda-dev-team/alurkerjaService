package repository

import (
	"alurkerjaService/model/domain"
	"context"
	"database/sql"
)

type JenisDataRepository interface {
	Create(ctx context.Context, tx *sql.Tx, jenisData domain.JenisData) (domain.JenisData, error)
	Update(ctx context.Context, tx *sql.Tx, jenisData domain.JenisData) (domain.JenisData, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.JenisData, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.JenisData, error)
}

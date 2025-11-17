package repository

import (
	"alurkerjaService/model/domain"
	"context"
	"database/sql"
	"errors"
)

type JenisDataRepositoryImpl struct {
}

func NewJenisDataRepositoryImpl() *JenisDataRepositoryImpl {
	return &JenisDataRepositoryImpl{}
}

func (repository *JenisDataRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, jenisData domain.JenisData) (domain.JenisData, error) {
	query := "INSERT INTO tb_jenis_data (jenis_data) VALUES (?)"
	_, err := tx.ExecContext(ctx, query, jenisData.JenisData)
	if err != nil {
		return domain.JenisData{}, err
	}

	return jenisData, nil
}

func (repository *JenisDataRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, jenisData domain.JenisData) (domain.JenisData, error) {
	query := "UPDATE tb_jenis_data SET jenis_data = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, jenisData.JenisData, jenisData.Id)
	if err != nil {
		return domain.JenisData{}, err
	}

	return jenisData, nil
}

func (repository *JenisDataRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM tb_jenis_data WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *JenisDataRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.JenisData, error) {
	query := "SELECT id, jenis_data FROM tb_jenis_data WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.JenisData{}, errors.New("id tidak ditemukan")
	}
	defer rows.Close()

	if rows.Next() {
		var jenisData domain.JenisData
		err := rows.Scan(&jenisData.Id, &jenisData.JenisData)
		if err != nil {
			return domain.JenisData{}, err
		}
		return jenisData, nil
	}

	return domain.JenisData{}, nil
}

func (repository *JenisDataRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.JenisData, error) {
	query := "SELECT id, jenis_data FROM tb_jenis_data ORDER BY id ASC"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return []domain.JenisData{}, err
	}
	defer rows.Close()

	var jenisDataList []domain.JenisData
	for rows.Next() {
		var jenisData domain.JenisData
		err := rows.Scan(&jenisData.Id, &jenisData.JenisData)
		if err != nil {
			return []domain.JenisData{}, err
		}

		jenisDataList = append(jenisDataList, jenisData)
	}

	return jenisDataList, nil
}

// jenis data opd
func (repository *JenisDataRepositoryImpl) CreateOpd(ctx context.Context, tx *sql.Tx, jenisDataOpd domain.JenisDataOpd) (domain.JenisDataOpd, error) {
	query := "INSERT INTO tb_jenis_data_opd (kode_opd, nama_opd, jenis_data) VALUES (?, ?, ?)"
	_, err := tx.ExecContext(ctx, query, jenisDataOpd.KodeOpd, jenisDataOpd.NamaOpd, jenisDataOpd.JenisData)
	if err != nil {
		return domain.JenisDataOpd{}, err
	}
	return jenisDataOpd, nil
}

func (repository *JenisDataRepositoryImpl) UpdateOpd(ctx context.Context, tx *sql.Tx, jenisDataOpd domain.JenisDataOpd) (domain.JenisDataOpd, error) {
	query := "UPDATE tb_jenis_data_opd SET kode_opd = ?, nama_opd = ?, jenis_data = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, jenisDataOpd.KodeOpd, jenisDataOpd.NamaOpd, jenisDataOpd.JenisData, jenisDataOpd.Id)
	if err != nil {
		return domain.JenisDataOpd{}, err
	}
	return jenisDataOpd, nil
}

func (repository *JenisDataRepositoryImpl) DeleteOpd(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM tb_jenis_data_opd WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (repository *JenisDataRepositoryImpl) FindByIdOpd(ctx context.Context, tx *sql.Tx, id int) (domain.JenisDataOpd, error) {
	query := "SELECT id, kode_opd, nama_opd, jenis_data FROM tb_jenis_data_opd WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.JenisDataOpd{}, errors.New("id tidak ditemukan")
	}
	defer rows.Close()

	if rows.Next() {
		var jenisDataOpd domain.JenisDataOpd
		err := rows.Scan(&jenisDataOpd.Id, &jenisDataOpd.KodeOpd, &jenisDataOpd.NamaOpd, &jenisDataOpd.JenisData)
		if err != nil {
			return domain.JenisDataOpd{}, err
		}
		return jenisDataOpd, nil
	}

	return domain.JenisDataOpd{}, nil
}

func (repository *JenisDataRepositoryImpl) FindAllOpd(ctx context.Context, tx *sql.Tx, kodeOpd string) ([]domain.JenisDataOpd, error) {
	var (
		query = "SELECT id, kode_opd, nama_opd, jenis_data FROM tb_jenis_data_opd WHERE 1=1"
		args  []interface{}
	)

	if kodeOpd != "" {
		query += " AND kode_opd = ?"
		args = append(args, kodeOpd)
	}

	query += " ORDER BY id ASC"

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return []domain.JenisDataOpd{}, err
	}
	defer rows.Close()

	var jenisDataOpdList []domain.JenisDataOpd
	for rows.Next() {
		var jenisDataOpd domain.JenisDataOpd
		err := rows.Scan(&jenisDataOpd.Id, &jenisDataOpd.KodeOpd, &jenisDataOpd.NamaOpd, &jenisDataOpd.JenisData)
		if err != nil {
			return []domain.JenisDataOpd{}, err
		}
		jenisDataOpdList = append(jenisDataOpdList, jenisDataOpd)
	}
	return jenisDataOpdList, nil
}

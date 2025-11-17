package repository

import (
	"alurkerjaService/helper"
	"alurkerjaService/model/domain"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sort"
)

type DataKinerjaOpdRepositoryImpl struct {
}

func NewDataKinerjaOpdRepositoryImpl() *DataKinerjaOpdRepositoryImpl {
	return &DataKinerjaOpdRepositoryImpl{}
}

func (repository *DataKinerjaOpdRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, dataKinerjaOpd domain.DataKinerjaOpd) (domain.DataKinerjaOpd, error) {
	query := `
		INSERT INTO tb_data_kinerja_opd
			(jenis_data_id, kode_opd, nama_opd, nama_data, rumus_perhitungan, sumber_data, instansi_produsen_data, keterangan)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(
		ctx, query,
		dataKinerjaOpd.JenisDataId,
		dataKinerjaOpd.KodeOpd,
		dataKinerjaOpd.NamaOpd,
		dataKinerjaOpd.NamaData,
		dataKinerjaOpd.RumusPerhitungan,
		dataKinerjaOpd.SumberData,
		dataKinerjaOpd.InstansiProdusenData,
		dataKinerjaOpd.Keterangan,
	)
	if err != nil {
		return domain.DataKinerjaOpd{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.DataKinerjaOpd{}, err
	}
	dataKinerjaOpd.Id = int(id)

	for _, target := range dataKinerjaOpd.Target {
		queryTarget := `
			INSERT INTO tb_target_opd (data_kinerja_opd_id, target, satuan, tahun)
			VALUES (?, ?, ?, ?)`
		_, err := tx.ExecContext(ctx, queryTarget, id, target.Target, target.Satuan, target.Tahun)
		if err != nil {
			return domain.DataKinerjaOpd{}, err
		}
	}

	return dataKinerjaOpd, nil
}

func (repository *DataKinerjaOpdRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, dataKinerjaOpd domain.DataKinerjaOpd) (domain.DataKinerjaOpd, error) {
	existing, err := repository.FindById(ctx, tx, dataKinerjaOpd.Id)
	if err != nil {
		return domain.DataKinerjaOpd{}, err
	}

	log.Printf("Updating Data Kinerja OPD ID: %d", dataKinerjaOpd.Id)
	log.Printf("Existing Data: %+v", existing)
	log.Printf("New Data: %+v", dataKinerjaOpd)

	query := `
		UPDATE tb_data_kinerja_opd
		SET
			jenis_data_id = ?,
			kode_opd = ?,
			nama_opd = ?,
			nama_data = ?,
			rumus_perhitungan = ?,
			sumber_data = ?,
			instansi_produsen_data = ?,
			keterangan = ?
		WHERE id = ?`
	_, err = tx.ExecContext(ctx, query,
		dataKinerjaOpd.JenisDataId,
		dataKinerjaOpd.KodeOpd,
		dataKinerjaOpd.NamaOpd,
		dataKinerjaOpd.NamaData,
		dataKinerjaOpd.RumusPerhitungan,
		dataKinerjaOpd.SumberData,
		dataKinerjaOpd.InstansiProdusenData,
		dataKinerjaOpd.Keterangan,
		dataKinerjaOpd.Id,
	)
	if err != nil {
		return domain.DataKinerjaOpd{}, fmt.Errorf("error updating data kinerja opd: %v", err)
	}

	existingTargets := make(map[int]domain.TargetOpd)
	for _, t := range existing.Target {
		existingTargets[t.Id] = t
	}

	targetToUpdate := make(map[int]bool)
	var newTargets []domain.TargetOpd

	for _, t := range dataKinerjaOpd.Target {
		if t.Id > 0 {
			targetToUpdate[t.Id] = true
			queryUpdateTarget := `
				UPDATE tb_target_opd
				SET target = ?, satuan = ?, tahun = ?
				WHERE id = ? AND data_kinerja_opd_id = ?`
			_, err := tx.ExecContext(ctx, queryUpdateTarget,
				t.Target, t.Satuan, t.Tahun, t.Id, dataKinerjaOpd.Id,
			)
			if err != nil {
				return domain.DataKinerjaOpd{}, fmt.Errorf("error updating target opd: %v", err)
			}
		} else {
			newTargets = append(newTargets, t)
		}
	}

	for _, t := range newTargets {
		queryInsertTarget := `
			INSERT INTO tb_target_opd (data_kinerja_opd_id, target, satuan, tahun)
			VALUES (?, ?, ?, ?)`
		_, err := tx.ExecContext(ctx, queryInsertTarget,
			dataKinerjaOpd.Id, t.Target, t.Satuan, t.Tahun,
		)
		if err != nil {
			return domain.DataKinerjaOpd{}, fmt.Errorf("error inserting target opd: %v", err)
		}
	}

	for targetId := range existingTargets {
		if !targetToUpdate[targetId] {
			_, err = tx.ExecContext(ctx, "DELETE FROM tb_target_opd WHERE id = ? AND data_kinerja_opd_id = ?",
				targetId, dataKinerjaOpd.Id,
			)
			if err != nil {
				return domain.DataKinerjaOpd{}, fmt.Errorf("error deleting target opd: %v", err)
			}
		}
	}

	return repository.FindById(ctx, tx, dataKinerjaOpd.Id)
}

func (repository *DataKinerjaOpdRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.DataKinerjaOpd, error) {
	queryCheck := "SELECT COUNT(*) FROM tb_data_kinerja_opd WHERE id = ?"
	var count int
	if err := tx.QueryRowContext(ctx, queryCheck, id).Scan(&count); err != nil {
		return domain.DataKinerjaOpd{}, err
	}
	if count == 0 {
		return domain.DataKinerjaOpd{}, errors.New("data tidak ditemukan")
	}

	query := `
		SELECT
			dk.id,
			dk.jenis_data_id,
			jd.jenis_data AS nama_jenis_data,
			dk.kode_opd,
			dk.nama_opd,
			dk.nama_data,
			dk.rumus_perhitungan,
			dk.sumber_data,
			dk.instansi_produsen_data,
			dk.keterangan,
			t.id AS target_id,
			t.target,
			t.satuan,
			t.tahun AS target_tahun
		FROM tb_data_kinerja_opd dk
		LEFT JOIN tb_target_opd t ON dk.id = t.data_kinerja_opd_id
		JOIN tb_jenis_data_opd jd ON dk.jenis_data_id = jd.id AND dk.kode_opd = jd.kode_opd
		WHERE dk.id = ?
		ORDER BY t.tahun DESC`
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.DataKinerjaOpd{}, err
	}
	defer rows.Close()

	var data *domain.DataKinerjaOpd
	var targets []domain.TargetOpd

	for rows.Next() {
		var (
			dk          domain.DataKinerjaOpd
			targetId    sql.NullInt64
			targetVal   sql.NullString
			satuan      sql.NullString
			targetTahun sql.NullString
		)
		if err := rows.Scan(
			&dk.Id,
			&dk.JenisDataId,
			&dk.JenisData,
			&dk.KodeOpd,
			&dk.NamaOpd,
			&dk.NamaData,
			&dk.RumusPerhitungan,
			&dk.SumberData,
			&dk.InstansiProdusenData,
			&dk.Keterangan,
			&targetId,
			&targetVal,
			&satuan,
			&targetTahun,
		); err != nil {
			return domain.DataKinerjaOpd{}, err
		}

		if data == nil {
			data = &dk
			data.Target = make([]domain.TargetOpd, 0)
		}

		if targetId.Valid && targetVal.Valid && satuan.Valid && targetTahun.Valid {
			targets = append(targets, domain.TargetOpd{
				Id:            int(targetId.Int64),
				DataKinerjaId: dk.Id,
				Target:        targetVal.String,
				Satuan:        satuan.String,
				Tahun:         targetTahun.String,
			})
		}
	}

	if data == nil {
		return domain.DataKinerjaOpd{}, errors.New("data tidak ditemukan")
	}

	sort.Slice(targets, func(i, j int) bool { return targets[i].Tahun > targets[j].Tahun })
	data.Target = targets

	log.Printf("Found Data Kinerja OPD ID: %d", data.Id)
	log.Printf("- Nama Data: %s", data.NamaData)
	log.Printf("- Jumlah Target: %d", len(data.Target))

	return *data, nil
}

func (repository *DataKinerjaOpdRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, kodeOpd string, jenisDataId int) ([]domain.DataKinerjaOpd, error) {
	var (
		query = `
			SELECT
				dk.id,
				dk.jenis_data_id,
				jd.jenis_data AS nama_jenis_data,
				dk.kode_opd,
				dk.nama_opd,
				dk.nama_data,
				dk.rumus_perhitungan,
				dk.sumber_data,
				dk.instansi_produsen_data,
				dk.keterangan,
				t.id AS target_id,
				t.target,
				t.satuan,
				t.tahun AS target_tahun
			FROM tb_data_kinerja_opd dk
			LEFT JOIN tb_target_opd t ON dk.id = t.data_kinerja_opd_id
			JOIN tb_jenis_data_opd jd ON dk.jenis_data_id = jd.id AND dk.kode_opd = jd.kode_opd
			WHERE 1=1`
		args []interface{}
	)

	if kodeOpd != "" {
		query += " AND dk.kode_opd = ?"
		args = append(args, kodeOpd)
	}
	if jenisDataId > 0 {
		query += " AND dk.jenis_data_id = ?"
		args = append(args, jenisDataId)
	}
	query += " ORDER BY dk.id ASC, t.tahun DESC"

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataMap := make(map[int]*domain.DataKinerjaOpd)
	var orderIds []int

	for rows.Next() {
		var (
			item        domain.DataKinerjaOpd
			targetId    sql.NullInt64
			targetVal   sql.NullString
			satuan      sql.NullString
			targetTahun sql.NullString
		)
		if err := rows.Scan(
			&item.Id,
			&item.JenisDataId,
			&item.JenisData,
			&item.KodeOpd,
			&item.NamaOpd,
			&item.NamaData,
			&item.RumusPerhitungan,
			&item.SumberData,
			&item.InstansiProdusenData,
			&item.Keterangan,
			&targetId,
			&targetVal,
			&satuan,
			&targetTahun,
		); err != nil {
			return nil, err
		}

		existing, ok := dataMap[item.Id]
		if !ok {
			item.Target = make([]domain.TargetOpd, 0)
			dataMap[item.Id] = &item
			orderIds = append(orderIds, item.Id)
			existing = &item
		}

		if targetId.Valid && targetVal.Valid && satuan.Valid && targetTahun.Valid {
			existing.Target = append(existing.Target, domain.TargetOpd{
				Id:            int(targetId.Int64),
				DataKinerjaId: item.Id,
				Target:        targetVal.String,
				Satuan:        satuan.String,
				Tahun:         targetTahun.String,
			})
		}
	}

	var result []domain.DataKinerjaOpd
	for _, id := range orderIds {
		if data, ok := dataMap[id]; ok {
			sort.Slice(data.Target, func(i, j int) bool { return data.Target[i].Tahun > data.Target[j].Tahun })
			result = append(result, *data)
		}
	}

	if len(result) == 0 {
		return []domain.DataKinerjaOpd{}, nil
	}
	return result, nil
}

func (repository *DataKinerjaOpdRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	_, err := tx.ExecContext(ctx, "DELETE FROM tb_data_kinerja_opd WHERE id = ?", id)
	helper.PanicIfError(err)
	return nil
}

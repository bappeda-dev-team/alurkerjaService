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

type DataKinerjaPemdaRepositoryImpl struct {
}

func NewDataKinerjaPemdaRepositoryImpl() *DataKinerjaPemdaRepositoryImpl {
	return &DataKinerjaPemdaRepositoryImpl{}
}

func (repository *DataKinerjaPemdaRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, dataKinerjaPemda domain.DataKinerjaPemda) (domain.DataKinerjaPemda, error) {
	// Insert data kinerja pemda
	query := "INSERT INTO tb_data_kinerja (jenis_data_id, nama_data, rumus_perhitungan, sumber_data, instansi_produsen_data, keterangan) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, dataKinerjaPemda.JenisDataId, dataKinerjaPemda.NamaData, dataKinerjaPemda.RumusPerhitungan, dataKinerjaPemda.SumberData, dataKinerjaPemda.InstansiProdusenData, dataKinerjaPemda.Keterangan)
	if err != nil {
		return domain.DataKinerjaPemda{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.DataKinerjaPemda{}, err
	}
	dataKinerjaPemda.Id = int(id)

	// Insert target data
	for _, target := range dataKinerjaPemda.Target {
		queryTarget := "INSERT INTO tb_target (data_kinerja_id, target, satuan) VALUES (?, ?, ?)"
		_, err := tx.ExecContext(ctx, queryTarget, id, target.Target, target.Satuan)
		if err != nil {
			return domain.DataKinerjaPemda{}, err
		}
	}

	return dataKinerjaPemda, nil
}

func (repository *DataKinerjaPemdaRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, dataKinerjaPemda domain.DataKinerjaPemda) (domain.DataKinerjaPemda, error) {
	// Cek dulu apakah data ada
	existing, err := repository.FindById(ctx, tx, dataKinerjaPemda.Id)
	if err != nil {
		return domain.DataKinerjaPemda{}, err
	}

	// Log data yang akan diupdate
	log.Printf("Updating Data Kinerja ID: %d", dataKinerjaPemda.Id)
	log.Printf("Existing Data: %+v", existing)
	log.Printf("New Data: %+v", dataKinerjaPemda)

	// Update data kinerja pemda
	query := `UPDATE tb_data_kinerja 
        SET nama_data = ?, 
            rumus_perhitungan = ?, 
            sumber_data = ?, 
            instansi_produsen_data = ?, 
            keterangan = ?
        WHERE id = ?`

	_, err = tx.ExecContext(ctx, query,
		dataKinerjaPemda.NamaData,
		dataKinerjaPemda.RumusPerhitungan,
		dataKinerjaPemda.SumberData,
		dataKinerjaPemda.InstansiProdusenData,
		dataKinerjaPemda.Keterangan,
		dataKinerjaPemda.Id)
	if err != nil {
		return domain.DataKinerjaPemda{}, fmt.Errorf("error updating data kinerja: %v", err)
	}

	// Buat map untuk existing target
	existingTargets := make(map[int]domain.Target)
	for _, target := range existing.Target {
		existingTargets[target.Id] = target
	}

	// Buat map untuk tracking target yang akan diupdate
	targetToUpdate := make(map[int]bool)
	var newTargets []domain.Target

	// Pisahkan target yang akan diupdate dan yang baru
	for _, target := range dataKinerjaPemda.Target {
		if target.Id > 0 {
			// Target dengan ID akan diupdate
			targetToUpdate[target.Id] = true

			// Update existing target
			queryUpdateTarget := `UPDATE tb_target 
                SET target = ?,
                    satuan = ?,
                    tahun = ?
                WHERE id = ? AND data_kinerja_id = ?`

			_, err := tx.ExecContext(ctx, queryUpdateTarget,
				target.Target,
				target.Satuan,
				target.Tahun,
				target.Id,
				dataKinerjaPemda.Id)
			if err != nil {
				return domain.DataKinerjaPemda{}, fmt.Errorf("error updating target: %v", err)
			}
		} else {
			// Target tanpa ID akan diinsert sebagai baru
			newTargets = append(newTargets, target)
		}
	}

	// Insert target baru
	for _, target := range newTargets {
		queryInsertTarget := `INSERT INTO tb_target 
            (data_kinerja_id, target, satuan, tahun) 
            VALUES (?, ?, ?, ?)`

		_, err := tx.ExecContext(ctx, queryInsertTarget,
			dataKinerjaPemda.Id,
			target.Target,
			target.Satuan,
			target.Tahun)
		if err != nil {
			return domain.DataKinerjaPemda{}, fmt.Errorf("error inserting target: %v", err)
		}

	}

	// Delete targets yang tidak ada di request update
	for targetId := range existingTargets {
		if !targetToUpdate[targetId] {
			// Target ini tidak ada di request update, jadi dihapus
			_, err = tx.ExecContext(ctx, "DELETE FROM tb_target WHERE id = ? AND data_kinerja_id = ?",
				targetId, dataKinerjaPemda.Id)
			if err != nil {
				return domain.DataKinerjaPemda{}, fmt.Errorf("error deleting target: %v", err)
			}
		}
	}

	// Get updated data
	return repository.FindById(ctx, tx, dataKinerjaPemda.Id)
}

func (repository *DataKinerjaPemdaRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.DataKinerjaPemda, error) {
	// Cek dulu apakah data kinerja ada
	queryCheck := "SELECT COUNT(*) FROM tb_data_kinerja WHERE id = ?"
	var count int
	err := tx.QueryRowContext(ctx, queryCheck, id).Scan(&count)
	if err != nil {
		return domain.DataKinerjaPemda{}, err
	}
	if count == 0 {
		return domain.DataKinerjaPemda{}, errors.New("data tidak ditemukan")
	}

	// Base query dengan LEFT JOIN untuk mendapatkan target
	query := `
        SELECT 
            dk.id,
            dk.jenis_data_id,
            dk.nama_data,
            dk.rumus_perhitungan,
            dk.sumber_data,
            dk.instansi_produsen_data,
            dk.keterangan,
            t.id as target_id,
            t.target,
            t.satuan,
            t.tahun as target_tahun
        FROM tb_data_kinerja dk
        LEFT JOIN tb_target t ON dk.id = t.data_kinerja_id
        WHERE dk.id = ?
        ORDER BY t.tahun ASC`

	// Execute query
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return domain.DataKinerjaPemda{}, err
	}
	defer rows.Close()

	var dataKinerjaPemda *domain.DataKinerjaPemda
	var targets []domain.Target

	for rows.Next() {
		var (
			dk          domain.DataKinerjaPemda
			targetId    sql.NullInt64
			targetVal   sql.NullString
			satuan      sql.NullString
			targetTahun sql.NullString
		)

		err := rows.Scan(
			&dk.Id,
			&dk.JenisDataId,
			&dk.NamaData,
			&dk.RumusPerhitungan,
			&dk.SumberData,
			&dk.InstansiProdusenData,
			&dk.Keterangan,

			&targetId,
			&targetVal,
			&satuan,
			&targetTahun,
		)
		if err != nil {
			return domain.DataKinerjaPemda{}, err
		}

		// Inisialisasi dataKinerjaPemda jika belum
		if dataKinerjaPemda == nil {
			dataKinerjaPemda = &dk
			dataKinerjaPemda.Target = make([]domain.Target, 0)
		}

		// Tambahkan target jika valid
		if targetId.Valid && targetVal.Valid && satuan.Valid && targetTahun.Valid {
			target := domain.Target{
				Id:            int(targetId.Int64),
				DataKinerjaId: dk.Id,
				Target:        targetVal.String,
				Satuan:        satuan.String,
				Tahun:         targetTahun.String,
			}
			targets = append(targets, target)
		}
	}

	// Tidak perlu cek dataKinerjaPemda == nil karena sudah dicek di awal

	// Sort target berdasarkan tahun
	sort.Slice(targets, func(i, j int) bool {
		return targets[i].Tahun < targets[j].Tahun
	})

	dataKinerjaPemda.Target = targets

	// Debug log
	log.Printf("Found Data Kinerja ID: %d", dataKinerjaPemda.Id)
	log.Printf("- Nama Data: %s", dataKinerjaPemda.NamaData)
	log.Printf("- Jumlah Target: %d", len(dataKinerjaPemda.Target))
	for _, t := range dataKinerjaPemda.Target {
		log.Printf("  * Target: %s, Tahun: %s, Satuan: %s", t.Target, t.Tahun, t.Satuan)
	}

	return *dataKinerjaPemda, nil
}

func (repository *DataKinerjaPemdaRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, jenisDataId int) ([]domain.DataKinerjaPemda, error) {
	var query string
	var args []interface{}

	// Base query dengan LEFT JOIN untuk mendapatkan target
	query = `
        SELECT 
            dk.id,
            dk.jenis_data_id,
            dk.nama_data,
            dk.rumus_perhitungan,
            dk.sumber_data,
            dk.instansi_produsen_data,
            dk.keterangan
            t.id as target_id,
            t.target,
            t.satuan,
            t.tahun as target_tahun
        FROM tb_data_kinerja dk
        LEFT JOIN tb_target t ON dk.id = t.data_kinerja_id
        WHERE 1=1`

	// Add filters
	if jenisDataId > 0 {
		query += " AND dk.jenis_data_id = ?"
		args = append(args, jenisDataId)
	}

	// Add ordering untuk data kinerja dan target
	query += " ORDER BY dk.id ASC, t.tahun ASC"

	// Execute query with filters
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map untuk menyimpan data kinerja yang sudah diproses
	dataKinerjaMap := make(map[int]*domain.DataKinerjaPemda)
	var resultIds []int // Untuk menjaga urutan data kinerja

	for rows.Next() {
		var (
			dataKinerja domain.DataKinerjaPemda
			targetId    sql.NullInt64
			targetVal   sql.NullString
			satuan      sql.NullString
			targetTahun sql.NullString
		)

		err := rows.Scan(
			&dataKinerja.Id,
			&dataKinerja.JenisDataId,
			&dataKinerja.NamaData,
			&dataKinerja.RumusPerhitungan,
			&dataKinerja.SumberData,
			&dataKinerja.InstansiProdusenData,
			&dataKinerja.Keterangan,
			&targetId,
			&targetVal,
			&satuan,
			&targetTahun,
		)
		if err != nil {
			return nil, err
		}

		// Cek apakah data kinerja sudah ada di map
		existingData, exists := dataKinerjaMap[dataKinerja.Id]
		if !exists {
			// Buat data kinerja baru
			dataKinerja.Target = make([]domain.Target, 0)
			dataKinerjaMap[dataKinerja.Id] = &dataKinerja
			resultIds = append(resultIds, dataKinerja.Id)
			existingData = &dataKinerja
		}

		// Tambahkan target jika valid
		if targetId.Valid && targetVal.Valid && satuan.Valid && targetTahun.Valid {
			target := domain.Target{
				Id:            int(targetId.Int64),
				DataKinerjaId: dataKinerja.Id,
				Target:        targetVal.String,
				Satuan:        satuan.String,
				Tahun:         targetTahun.String,
			}
			existingData.Target = append(existingData.Target, target)
		}
	}

	// Konversi map ke slice dengan urutan yang benar
	var result []domain.DataKinerjaPemda
	for _, id := range resultIds {
		if data, ok := dataKinerjaMap[id]; ok {
			// Sort target berdasarkan tahun
			sort.Slice(data.Target, func(i, j int) bool {
				return data.Target[i].Tahun < data.Target[j].Tahun
			})
			result = append(result, *data)
		}
	}

	if len(result) == 0 {
		return []domain.DataKinerjaPemda{}, nil
	}

	return result, nil
}

func (repository *DataKinerjaPemdaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	// Delete targets first due to foreign key constraint
	_, err := tx.ExecContext(ctx, "DELETE FROM tb_target WHERE data_kinerja_id = ?", id)
	helper.PanicIfError(err)

	return nil
}

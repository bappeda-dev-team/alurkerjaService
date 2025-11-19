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
		queryTarget := "INSERT INTO tb_target (data_kinerja_id, target, satuan, tahun) VALUES (?, ?, ?, ?)"
		_, err := tx.ExecContext(ctx, queryTarget, id, target.Target, target.Satuan, target.Tahun)
		if err != nil {
			log.Printf("Error inserting target: %v", err)
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

	// UBAH: Mulai dari tb_jenis_data, bukan dari tb_data_kinerja
	// Jadi jenis data tetap muncul meskipun tidak ada data kinerja
	query = `
	SELECT 
		jd.id as jenis_data_id,
		jd.jenis_data as nama_jenis_data,
		dk.id,
		dk.nama_data,
		dk.rumus_perhitungan,
		dk.sumber_data,
		dk.instansi_produsen_data,
		dk.keterangan,
		t.id as target_id,
		t.target,
		t.satuan,
		t.tahun as target_tahun
	FROM tb_jenis_data jd
	LEFT JOIN tb_data_kinerja dk ON jd.id = dk.jenis_data_id
	LEFT JOIN tb_target t ON dk.id = t.data_kinerja_id
	WHERE 1=1`

	// Add filters
	if jenisDataId > 0 {
		query += " AND jd.id = ?"
		args = append(args, jenisDataId)
	}

	// Add ordering
	query += " ORDER BY jd.id ASC, dk.id ASC, t.tahun DESC"

	// Execute query with filters
	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Map untuk menyimpan jenis data
	jenisDataMap := make(map[int]*domain.DataKinerjaPemda)
	dataKinerjaMap := make(map[int]*domain.DataKinerjaPemda)
	var jenisDataIds []int

	for rows.Next() {
		var (
			jenisDataId          int
			jenisDataNama        string
			dataKinerjaId        sql.NullInt64
			namaData             sql.NullString
			rumusPerhitungan     sql.NullString
			sumberData           sql.NullString
			instansiProdusenData sql.NullString
			keterangan           sql.NullString
			targetId             sql.NullInt64
			targetVal            sql.NullString
			satuan               sql.NullString
			targetTahun          sql.NullString
		)

		err := rows.Scan(
			&jenisDataId,
			&jenisDataNama,
			&dataKinerjaId,
			&namaData,
			&rumusPerhitungan,
			&sumberData,
			&instansiProdusenData,
			&keterangan,
			&targetId,
			&targetVal,
			&satuan,
			&targetTahun,
		)
		if err != nil {
			return nil, err
		}

		// Pastikan jenis data ada di map (untuk jenis data tanpa data kinerja)
		if _, exists := jenisDataMap[jenisDataId]; !exists {
			jenisDataMap[jenisDataId] = &domain.DataKinerjaPemda{
				JenisDataId: jenisDataId,
				JenisData:   jenisDataNama,
				Target:      make([]domain.Target, 0),
			}
			jenisDataIds = append(jenisDataIds, jenisDataId)
		}

		// Jika ada data kinerja
		if dataKinerjaId.Valid {
			dataKinerjaIdInt := int(dataKinerjaId.Int64)

			// Cek apakah data kinerja sudah ada di map
			existingData, exists := dataKinerjaMap[dataKinerjaIdInt]
			if !exists {
				// Buat data kinerja baru
				dataKinerja := &domain.DataKinerjaPemda{
					Id:                   dataKinerjaIdInt,
					JenisDataId:          jenisDataId,
					JenisData:            jenisDataNama,
					NamaData:             namaData.String,
					RumusPerhitungan:     rumusPerhitungan.String,
					SumberData:           sumberData.String,
					InstansiProdusenData: instansiProdusenData.String,
					Keterangan:           keterangan.String,
					Target:               make([]domain.Target, 0),
				}
				dataKinerjaMap[dataKinerjaIdInt] = dataKinerja
				existingData = dataKinerja
			}

			// Tambahkan target jika valid
			if targetId.Valid && targetVal.Valid && satuan.Valid && targetTahun.Valid {
				target := domain.Target{
					Id:            int(targetId.Int64),
					DataKinerjaId: dataKinerjaIdInt,
					Target:        targetVal.String,
					Satuan:        satuan.String,
					Tahun:         targetTahun.String,
				}
				existingData.Target = append(existingData.Target, target)
			}
		}
	}

	// Konversi map ke slice dengan urutan yang benar
	var result []domain.DataKinerjaPemda

	// Tambahkan semua jenis data (termasuk yang tidak punya data kinerja)
	for _, jenisDataId := range jenisDataIds {
		jenisDataInfo := jenisDataMap[jenisDataId]

		// Cari semua data kinerja untuk jenis data ini
		hasDataKinerja := false
		for _, dataKinerja := range dataKinerjaMap {
			if dataKinerja.JenisDataId == jenisDataId {
				hasDataKinerja = true
				// Sort target berdasarkan tahun DESC
				sort.Slice(dataKinerja.Target, func(i, j int) bool {
					return dataKinerja.Target[i].Tahun > dataKinerja.Target[j].Tahun
				})
				result = append(result, *dataKinerja)
			}
		}

		// Jika tidak ada data kinerja, tambahkan jenis data dengan data kinerja kosong
		if !hasDataKinerja {
			result = append(result, *jenisDataInfo)
		}
	}

	return result, nil
}

func (repository *DataKinerjaPemdaRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	// Delete targets first due to foreign key constraint
	_, err := tx.ExecContext(ctx, "DELETE FROM tb_data_kinerja WHERE id = ?", id)
	helper.PanicIfError(err)

	return nil
}

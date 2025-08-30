CREATE TABLE tb_data_kinerja(
    id INT PRIMARY KEY AUTO_INCREMENT,
    jenis_data_id INT NOT NULL,
    nama_data TEXT,
    rumus_perhitungan TEXT,
    sumber_data TEXT,
    instansi_produsen_data VARCHAR(255) NOT NULL,
    keterangan TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_tb_data_kinerja_jenis_data FOREIGN KEY (jenis_data_id)
    REFERENCES tb_jenis_data(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
)ENGINE=InnoDB;
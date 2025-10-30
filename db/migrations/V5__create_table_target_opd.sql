CREATE TABLE tb_target_opd(
    id INT PRIMARY KEY AUTO_INCREMENT,
    target VARCHAR(255) NOT NULL,
    satuan VARCHAR(255) NOT NULL,
    tahun VARCHAR(255) NOT NULL,
    data_kinerja_opd_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_tb_target_data_kinerja_opd FOREIGN KEY (data_kinerja_opd_id)
    REFERENCES tb_data_kinerja_opd(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
)ENGINE=InnoDB;
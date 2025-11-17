-- DROP foreign key lama
ALTER TABLE tb_data_kinerja_opd
DROP FOREIGN KEY fk_tb_data_kinerja_opd_jenis_data;

-- ADD foreign key baru
ALTER TABLE tb_data_kinerja_opd
ADD CONSTRAINT fk_tb_data_kinerja_opd_jenis_data_opd FOREIGN KEY (jenis_data_id)
REFERENCES tb_jenis_data_opd(id)
ON DELETE CASCADE
ON UPDATE CASCADE;
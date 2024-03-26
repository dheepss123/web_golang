package models

import (
	"database/sql"
	"fmt"
	"web_golang/config"
	"web_golang/entities"
)

type BiodataModel struct {
	conn *sql.DB
}

func NewBiodataModel() *BiodataModel {
	conn, err := config.DBConection()
	if err != nil {
		panic(err)
	}

	return &BiodataModel{
		conn: conn,
	}
}

// Menampilkan isi database ke index
func (p *BiodataModel) FindAll() ([]entities.Biodata, error) {

	rows, err := p.conn.Query("SELECT * FROM biodata")
	if err != nil {
		return []entities.Biodata{}, err
	}
	defer rows.Close()

	var dataBiodata []entities.Biodata
	for rows.Next() {
		var biodata entities.Biodata
		rows.Scan(
			&biodata.Id,
			&biodata.Nama,
			&biodata.Nim,
			&biodata.Prodi,
			&biodata.Divisi,
		)
		dataBiodata = append(dataBiodata, biodata)
	}
	return dataBiodata, nil

}

// Simpan data ke database
func (p *BiodataModel) Create(biodata entities.Biodata) bool {
	result, err := p.conn.Exec("insert into biodata(nama, nim, prodi, divisi) values(?,?,?,?)", biodata.Nama, biodata.Nim, biodata.Prodi, biodata.Divisi)

	if err != nil {
		fmt.Print(err)
		return false
	}
	LastInsertId, _ := result.LastInsertId()
	return LastInsertId > 0
}

func (p *BiodataModel) Find(id int64, biodata *entities.Biodata) error {
	return p.conn.QueryRow("select * from biodata where id = ?", id).Scan(
		&biodata.Id,
		&biodata.Nama,
		&biodata.Nim,
		&biodata.Prodi,
		&biodata.Divisi,
	)
}
func (p *BiodataModel) Update(biodata entities.Biodata) error {

	_, err := p.conn.Exec(
		"update pasien set nama_lengkap = ?, nik = ?, jenis_kelamin = ?, tempat_lahir = ?, tanggal_lahir = ?, alamat = ?, no_hp = ? where id = ?",
		biodata.Nama, biodata.Nim, biodata.Prodi, biodata.Divisi)

	if err != nil {
		return err
	}

	return nil
}

func (p *BiodataModel) Delete(id int64) {
	p.conn.Exec("delete from biodata where id = ?", id)
}
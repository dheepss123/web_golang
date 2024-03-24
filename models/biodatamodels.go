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

package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb() {
	db, err := sql.Open("mysql", "root:root@/go_product")
	if err != nil {

		panic(err)
	}
	log.Println("Database conected")
	DB = db
}

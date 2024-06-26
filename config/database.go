package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DBConection() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "infinity"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return db, err
}

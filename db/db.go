package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"meli/config"
	)

func DbConn() (db *sql.DB) {
	db, err := sql.Open(config.DB_DRIVER, config.DB_STRING_CONNECTION)
	if err != nil {
		panic(err.Error())
	}
	return db
}
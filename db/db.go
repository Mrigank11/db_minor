package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *sql.DB
	err error
)

func Query(q string, args ...interface{}) (rows *sql.Rows) {
	rows, err := db.Query(q, args...)
	if err != nil {
		log.Error("Failed to execute query", err)
		return nil
	}
	return rows
}

func init() {
	db, err = sql.Open("mysql", "root:root@/db_minor")
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("init DB.go")
}

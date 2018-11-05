package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	DB  *sql.DB
	err error
)

func Query(q string, args ...interface{}) (rows *sql.Rows) {
	rows, err := DB.Query(q, args...)
	if err != nil {
		log.Error("Failed to execute query: ", q, args, err)
		return nil
	}
	return rows
}

func init() {
	DB, err = sql.Open("mysql", "root:root@/db_minor")
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("init DB.go")
}

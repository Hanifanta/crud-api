package app

import (
	"crud-api/helper"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func NewDB(dbDriver string, dbSource string) *sql.DB {
	db, err := sql.Open(dbDriver, dbSource)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

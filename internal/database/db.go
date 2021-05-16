package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// var (
// 	db *sql.DB
// )

func New(dbDriver, dbSource string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	if db, err = sql.Open(dbDriver, dbSource); err != nil {
		return nil, err
	}
	return db, nil
}

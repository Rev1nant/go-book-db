package datebase

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

func NewDB(dsn string) *DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return &DB{DB: db}
}

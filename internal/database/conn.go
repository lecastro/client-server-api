package database

import (
	"database/sql"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func Conn() (*sql.DB, error) {
	path := "./internal/database/dollarPrice.db"

	db, err := sql.Open("sqlite3", path)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db, nil
}

func Close(db *sql.DB) {
	db.Close()
}

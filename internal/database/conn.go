package database

import (
	"database/sql"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func Conn() error {
	path := "./internal/database/dollarPrice.db"

	db, err := sql.Open("sqlite3", path)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return nil
}

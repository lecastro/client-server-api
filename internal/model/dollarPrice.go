package model

import (
	"context"
	"database/sql"
	"time"

	"github.com/lecastro/client-server-api/internal/database"
)

type DollarPrice struct {
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}

func Create(dp *DollarPrice) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)

	defer cancel()
	db, err := database.Conn()

	defer database.Close(db)

	if err != nil {
		panic(err)
	}

	makeTable(db)

	stmt, err := db.PrepareContext(ctx, "INSERT INTO dollar_price (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, dp.Code, dp.Codein, dp.Name, dp.High, dp.Low, dp.VarBid, dp.PctChange, dp.Bid, dp.Ask, dp.Timestamp, dp.CreateDate)

	if err != nil {
		return err
	}

	return nil
}

func makeTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS dollar_price (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code VARCHAR(30),
			codein VARCHAR(30),
			name VARCHAR(30),
			high VARCHAR(30),
			low VARCHAR(30),
			varBid VARCHAR(30),
			pctChange VARCHAR(30),
			bid VARCHAR(30),
			ask VARCHAR(30),
			timestamp VARCHAR(30),
			create_date VARCHAR(30)
		)
	`)

	if err != nil {
		panic(err)
	}
}

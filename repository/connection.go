package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"github.com/paul-ilves/wanaku-api-go/config"
)

var (
	client *sqlx.DB
)

func OpenDBConnection() {
	db, err := PostgreSQLConnection()
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}

	client = db
}

func CloseDBConnection() {
	err := client.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		panic(err)
	}
	client = nil
}

func PostgreSQLConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", config.C.DBConnectionURL)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(10)       // the default is 0 (unlimited)
	db.SetMaxIdleConns(2)        // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(10000) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

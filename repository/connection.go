package repository

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"
	"os"
)

var (
	client *sqlx.DB

	//maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	//maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	//maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	//host := os.Getenv("DB_HOST")
)

func OpenDBConnection() {
	// Define a new PostgreSQL connection.
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
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	url := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, name, user, password)
	// Define database connection for PostgreSQL.
	//url := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, name, user, password)
	db, err := sqlx.Connect("pgx", url)
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

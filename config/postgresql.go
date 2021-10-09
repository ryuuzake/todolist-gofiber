package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

func PostgreSQLConnection() (*sqlx.DB, error) {
	// Define database connection settings.
	//maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	//maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	//maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	maxConn := Config.DBMaxConnections
	maxIdleConn := Config.DBMaxIdleConnections
	maxLifetimeConn := Config.DBMaxLifetimeConnections

	// Define database connection for PostgreSQL.
	//dbServerUrl := os.Getenv("DB_SERVER_URL")
	dbServerUrl := Config.DBServerUrl
	db, err := sqlx.Connect("pgx", dbServerUrl)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Set database connection settings.
	db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

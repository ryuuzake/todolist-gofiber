package config

import "github.com/jmoiron/sqlx"

func OpenDBConnection() (*sqlx.DB, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return db, nil
}

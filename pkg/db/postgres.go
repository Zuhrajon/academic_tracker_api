package db

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const defaultPostgresDSN = "postgresql://academic-tracker-user:academic-tracker-password@localhost:5432/academic-tracker-db?sslmode=disable"

func NewPostgresConnection() (*sql.DB, error) {
	postgresDSN := os.Getenv("POSTGRES_DSN")
	if postgresDSN == "" {
		postgresDSN = defaultPostgresDSN
	}

	database, err := sql.Open("pgx", postgresDSN)
	if err != nil {
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	return database, nil
}

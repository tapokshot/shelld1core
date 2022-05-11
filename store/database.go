package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store struct {
	DB *sql.DB
}

func NewPostgresDB(conf *PostgresDBConfig) (*Store, error) {
	dbUrl := createPostgresDBUrl(conf)
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		DB: db,
	}, nil
}

func createPostgresDBUrl(conf *PostgresDBConfig) string {
	return fmt.Sprintf("host=%s dbname=%s sslmode=disable user=%s password=%s",
		conf.Host,
		conf.Dbname,
		conf.Username,
		conf.Password)
}

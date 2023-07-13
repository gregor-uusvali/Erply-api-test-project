package dbrepo

import (
	"Erply-api-test-project/repository"
	"database/sql"
)

type sqliteDBRepo struct {
	DB *sql.DB
}

func NewSQLiteRepo(conn *sql.DB) repository.DatabaseRepo {
	return &sqliteDBRepo{
		DB: conn,
	}
}

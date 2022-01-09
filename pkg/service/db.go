package service

import "database/sql"

type DatabaseService struct {
	Db *sql.DB
}

func NewDatabaseService(db *sql.DB) *DatabaseService {
	return &DatabaseService{Db: db}
}

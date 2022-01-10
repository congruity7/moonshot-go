package service

import (
	"gorm.io/gorm"
)

type DatabaseService struct {
	Db *gorm.DB
}

func NewDatabaseService(db *gorm.DB) *DatabaseService {
	return &DatabaseService{Db: db}
}

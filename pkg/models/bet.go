package models

import (
	"time"

	"gorm.io/gorm"
)

type Bet struct {
	gorm.Model
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Trash      bool
	Multiplier float32
}

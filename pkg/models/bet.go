package models

import (
	"gorm.io/gorm"
)

type Bet struct {
	gorm.Model
	Trash      bool
	Multiplier float32
}

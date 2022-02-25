package models

import (
	"gorm.io/gorm"
)

type Round struct {
	gorm.Model
	Trash      bool
	Multiplier float32
}

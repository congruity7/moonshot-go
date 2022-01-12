package models

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Trash         bool
	MaxMultiplier float32 `json:"max_multiplier"`
	MinMultiplier float32 `json:"min_multiplier"`
}

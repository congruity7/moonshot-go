package models

import (
	"time"
)

type Bet struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Trash      bool
	Multiplier float32
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type PlacedBet struct {
	gorm.Model
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Trash            bool
	User             User
	Wallet           Wallet
	Bet              Bet
	Status           uint8
	PayoutMultiplier float32
	Amount           float32
}

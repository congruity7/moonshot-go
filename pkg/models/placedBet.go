package models

import (
	"time"
)

type PlacedBet struct {
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

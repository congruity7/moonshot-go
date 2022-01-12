package models

import (
	"gorm.io/gorm"
)

type PlacedBet struct {
	gorm.Model
	Trash            bool
	UserID           uint
	WalletID         uint
	BetID            uint
	Status           uint8
	PayoutMultiplier float32
	Amount           float32
	// User             User   `gorm:"foreignKey:UserID"`
	// Wallet           Wallet `gorm:"foreignKey:WalletID"`
	// Bet              Bet    `gorm:"foreignKey:BetID"`
}

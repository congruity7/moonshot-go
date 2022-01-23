package models

import (
	"gorm.io/gorm"
)

type PlacedBet struct {
	gorm.Model
	Trash            bool
	UserID           uint
	WalletID         uint
	RoundID          uint
	Status           uint8
	PayoutMultiplier float32
	Amount           float32
	// User             User   `gorm:"foreignKey:UserID"`
	// Wallet           Wallet `gorm:"foreignKey:WalletID"`
	// Round            Round  `gorm:"foreignKey:BetID"`
}

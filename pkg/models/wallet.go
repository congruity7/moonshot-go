package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	Trash          bool
	WalletAddress  string
	WalletStatus   uint8
	TradingBalance float32
	FundingBalance float32
	User           User `gorm:"foreignKey:ID"`
}

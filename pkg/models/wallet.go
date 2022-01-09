package models

import (
	"time"
)

type Wallet struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Trash          bool
	WalletAddress  string
	WalletStatus   uint8
	TradingBalance float32
	FundingBalance float32
	UserID         uint
}

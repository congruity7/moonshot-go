package models

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Trash          bool
	UserID         uint    `json:"user_id"`
	WalletAddress  string  `json:"wallet_address"`
	WalletStatus   uint8   `json:"wallet_status"`
	TradingBalance float32 `json:"trading_balance"`
	FundingBalance float32 `json:"funding_balance"`
	WagerAmount    float32 `json:"wager_amount"`
	//User           User    `gorm:"foreignKey:UserID"`
}

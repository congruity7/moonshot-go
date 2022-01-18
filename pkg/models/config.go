package models

import (
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model
	Trash           bool
	MaxMultiplier   float32 `json:"max_multiplier"`
	MinMultiplier   float32 `json:"min_multiplier"`
	Max             float32 `json:"max"`
	Min             float32 `json:"min"`
	SpeedSetting    float32 `json:"speed_setting"`
	CooldownSetting float32 `json:"cooldown_setting"`
	HouseEdge       float32 `json:"house_edge"`
	Round           float32 `json:"round"`
	MinTotalWager   float32 `json:"min_total_wager"`
}

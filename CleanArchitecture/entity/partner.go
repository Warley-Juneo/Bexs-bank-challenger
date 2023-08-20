package entity

import (
	"time"
)

type Partner struct {
	ID           int32  `json:"id"`
	Trading_name string `json:"trading_name"`
	Document     string `json:"document"`
	Currency     string `json:"currency"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

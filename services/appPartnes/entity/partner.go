package entity

import (
	"time"
)

type Partner struct {
	ID           int32  `json:"id"`
	Trading_name string `json:"trading_name"`
	Document     string `json:"document"`
	Currency     string `json:"currency"`
	Created_at   time.Time
	Updated_at   time.Time
}

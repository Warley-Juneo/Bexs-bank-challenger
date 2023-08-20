package entity

import (
	"time"
)

type Consumer struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	National_id string `json:"national_id"`

	Created_at time.Time
	Updated_at time.Time
}

type Payment struct {
	ID         int32    `json:"id"`
	Partner_id int32    `json:"partner_id"`
	Amount     float64  `json:"amount"`
	Consumer   Consumer `json:"consumer"`

	Created_at time.Time
	Updated_at time.Time
}

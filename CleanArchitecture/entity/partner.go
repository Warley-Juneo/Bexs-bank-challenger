package entity

import (
	"errors"
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

func (partner *Partner) Validate() error {
	// função para consultar o documento no banco de dados e verificar se já existe
	list_cucurrency_valid := []string{"USD", "EUR", "GBP"}

	for _, currency := range list_cucurrency_valid {
		if partner.Currency == currency {
			return nil
		}
	}
	return errors.New("currency invalid")
}

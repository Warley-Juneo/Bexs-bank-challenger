package utils

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/payment/dto/paymentdto"
	"github.com/payment/consts"
)

func CalculateForeingAmount(partner_id int32, amount float64) float64 {
	url := fmt.Sprintf("http://localhost:3001/api/v1/partners/%d", partner_id)

	r, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer r.Body.Close()

	partner := dto.PartnerData{}
	err = json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		return 0
	}

	if partner.Currency == "USD" {
		return amount * consts.USD
	} else if partner.Currency == "EUR" {
		return amount * consts.EUR
	} else if partner.Currency == "GBP" {
		return amount * consts.GBP
	} else {
		return 0
	}
}
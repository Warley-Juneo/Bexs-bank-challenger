package dto

type ConsumerData struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	National_id string `json:"national_id"`
}

type PaymentData struct {
	Partner_id int32   `json:"partner_id"`
	Amount     float64 `json:"amount"`
	Consumer   ConsumerData	`json:"consumer"`
}
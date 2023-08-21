package dto

type ConsumerData struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	National_id string `json:"national_id"`
}

type PaymentData struct {
	Partner_id int32        `json:"partner_id"`
	Amount     float64      `json:"amount"`
	Consumer   ConsumerData `json:"consumer"`
}

type PaymentResponse struct {
	ID             int32        `json:"id"`
	Partner_id     int32        `json:"partner_id"`
	Amount         float64      `json:"amount"`
	Foreing_amount float64      `json:"foreing_amount"`
	Consumer       ConsumerData `json:"consumer_id"`

	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type PartnerData struct {
	ID           int32  `json:"id"`
	Trading_name string `json:"trading_name"`
	Document     string `json:"document"`
	Currency     string `json:"currency"`
}

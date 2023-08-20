package dto_test

import (
	"encoding/json"
	"testing"

	"github.com/payment/dto"
	"github.com/stretchr/testify/require"
)

func TestPaymentDataJSONEncoding(t *testing.T) {
	consumerData := dto.ConsumerData {
		ID:          1,
		Name:        "Test Name",
		National_id: "123456789",
	}
	
	paymentData := dto.PaymentData{
		Partner_id: 1,
		Amount:     100.00,
		Consumer:   consumerData,
	}

	jsonData, err := json.Marshal(paymentData)
	require.NoError(t, err)

	var decodedPaymentData dto.PaymentData
	err = json.Unmarshal(jsonData, &decodedPaymentData)
	require.NoError(t, err)

	require.Equal(t, paymentData, decodedPaymentData)
}


package entity_test

import (
	"time"
	"testing"

	"github.com/payment/entity"

	"github.com/stretchr/testify/require"
)

func TestPaymentsFields(t *testing.T) {
	consumerData := entity.Consumer {
		ID: 123,
		Name: "Test Name",
		National_id: "123456789",

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}


	paymentData := entity.Payment {
		ID: 123,
		Partner_id: 123,
		Amount: 123.45,
		Consumer: consumerData,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	require.Equal(t, int32(123), paymentData.ID)
	require.Equal(t, int32(123), paymentData.Partner_id)
	require.Equal(t, 123.45, paymentData.Amount)
	require.Equal(t, consumerData, paymentData.Consumer)

}
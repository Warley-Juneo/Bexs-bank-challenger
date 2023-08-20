package entity_test

import (
	"time"
	"testing"

	"github.com/wjuneo/bexs/entity"

	"github.com/stretchr/testify/require"
)

func TestPartnerFields(t *testing.T) {
	partnerData := entity.Partner {
		ID: 			123,
		Trading_name: 	"Test Trading Name",
		Document: 		"123456789",
		Currency: 		"USD",
		CreatedAt: 		time.Now(),
		UpdatedAt: 		time.Now(),
	}

	require.Equal(t, int32(123), partnerData.ID)
	require.Equal(t, "Test Trading Name", partnerData.Trading_name)
	require.Equal(t, "123456789", partnerData.Document)
	require.Equal(t, "USD", partnerData.Currency)
}
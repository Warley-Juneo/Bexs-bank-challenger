package partnerdto_test

import (
	"encoding/json"
	"testing"

	"github.com/wjuneo/bexs/dto/partnerdto"
	"github.com/stretchr/testify/require"
)


func TestPartnerDataJSONEncoding(t *testing.T) {
	partnerData := partnerdto.PartnerData{
		TradingName: "Test Trading Name",
		Document:    "123456789",
		Currency:    "USD",
	}

	jsonData, err := json.Marshal(partnerData)
	require.NoError(t, err)

	var decodedPartnerData partnerdto.PartnerData
	err = json.Unmarshal(jsonData, &decodedPartnerData)
	require.NoError(t, err)

	require.Equal(t, partnerData, decodedPartnerData)
}

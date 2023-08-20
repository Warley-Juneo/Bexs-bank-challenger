package partnerrepository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/wjuneo/bexs/entity"
	"github.com/wjuneo/bexs/repository/partnerrepository/mocks"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreatePartner(t *testing.T) {
	fakeDBPartner := entity.Partner{}
	faker.FakeData(&fakeDBPartner)

	mockCrlr := gomock.NewController(t)
	defer mockCrlr.Finish()
	mockPartnerRepo := mocks.NewMockPartnerRepository(mockCrlr)
	mockPartnerRepo.EXPECT().SavePartners(context.Background(), fakeDBPartner).Return(&fakeDBPartner, nil)

	savedPartner, err := mockPartnerRepo.SavePartners(context.Background(), fakeDBPartner)
	require.NoError(t, err)
	require.Equal(t, &fakeDBPartner, savedPartner)
}

func TestCreatePartnerError(t *testing.T) {
	fakeDBPartner := entity.Partner{
		ID:           123,
		Trading_name: "Test Trading Name",
		Document:     "123456789",
	}

	mockCrlr := gomock.NewController(t)
	defer mockCrlr.Finish()
	mockPartnerRepo := mocks.NewMockPartnerRepository(mockCrlr)
	mockPartnerRepo.EXPECT().SavePartners(context.Background(), fakeDBPartner).Return(nil, fmt.Errorf("error"))

	savedPartner, err := mockPartnerRepo.SavePartners(context.Background(), fakeDBPartner)
	require.Error(t, err)
	require.Nil(t, savedPartner)
}

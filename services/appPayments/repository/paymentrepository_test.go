package repository_test

import (
	"context"
	"testing"
	"fmt"

	"github.com/payment/entity"
	"github.com/payment/repository/mocks"

	"github.com/go-faker/faker/v4"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreatePayment(t *testing.T) {
	fakeDBPayment := entity.Payment{}
	faker.FakeData(&fakeDBPayment)

	mockCrlr := gomock.NewController(t)
	defer mockCrlr.Finish()
	mockPartnerRepo := mocks.NewMockPaymentRepository(mockCrlr)
	mockPartnerRepo.EXPECT().SavePayment(context.Background(), fakeDBPayment).Return(&fakeDBPayment, nil)

	savedPayment, err := mockPartnerRepo.SavePayment(context.Background(), fakeDBPayment)
	require.NoError(t, err)
	require.Equal(t, &fakeDBPayment, savedPayment)
}

func TestCreatePaymentError(t *testing.T) {
	fakeDBPayment := entity.Payment{
		ID: 123,
		Partner_id: 123,
		Amount: 123.45,
		Consumer_id: 123,
	}

	mockCrlr := gomock.NewController(t)
	defer mockCrlr.Finish()
	mockPartnerRepo := mocks.NewMockPaymentRepository(mockCrlr)
	mockPartnerRepo.EXPECT().SavePayment(context.Background(), fakeDBPayment).Return(nil, fmt.Errorf("error"))

	savedPayment, err := mockPartnerRepo.SavePayment(context.Background(), fakeDBPayment)
	require.Error(t, err)
	require.Nil(t, savedPayment)
}
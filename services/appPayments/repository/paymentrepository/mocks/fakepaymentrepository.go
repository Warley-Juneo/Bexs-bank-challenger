// Code generated by MockGen. DO NOT EDIT.
// Source: repository/paymentrepository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/payment/entity"
)

// MockPaymentRepository is a mock of PaymentRepository interface.
type MockPaymentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentRepositoryMockRecorder
}

// MockPaymentRepositoryMockRecorder is the mock recorder for MockPaymentRepository.
type MockPaymentRepositoryMockRecorder struct {
	mock *MockPaymentRepository
}

// NewMockPaymentRepository creates a new mock instance.
func NewMockPaymentRepository(ctrl *gomock.Controller) *MockPaymentRepository {
	mock := &MockPaymentRepository{ctrl: ctrl}
	mock.recorder = &MockPaymentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentRepository) EXPECT() *MockPaymentRepositoryMockRecorder {
	return m.recorder
}

// FindConsumer mocks base method.
func (m *MockPaymentRepository) FindConsumer(ctx context.Context, national_id string) (*entity.Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindConsumer", ctx, national_id)
	ret0, _ := ret[0].(*entity.Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindConsumer indicates an expected call of FindConsumer.
func (mr *MockPaymentRepositoryMockRecorder) FindConsumer(ctx, national_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindConsumer", reflect.TypeOf((*MockPaymentRepository)(nil).FindConsumer), ctx, national_id)
}

// FindPayment mocks base method.
func (m *MockPaymentRepository) FindPayment(ctx context.Context, payment_id int32) (*entity.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPayment", ctx, payment_id)
	ret0, _ := ret[0].(*entity.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPayment indicates an expected call of FindPayment.
func (mr *MockPaymentRepositoryMockRecorder) FindPayment(ctx, payment_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPayment", reflect.TypeOf((*MockPaymentRepository)(nil).FindPayment), ctx, payment_id)
}

// SavePayment mocks base method.
func (m *MockPaymentRepository) SavePayment(ctx context.Context, payment entity.Payment) (*entity.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePayment", ctx, payment)
	ret0, _ := ret[0].(*entity.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SavePayment indicates an expected call of SavePayment.
func (mr *MockPaymentRepositoryMockRecorder) SavePayment(ctx, payment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePayment", reflect.TypeOf((*MockPaymentRepository)(nil).SavePayment), ctx, payment)
}

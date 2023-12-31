// Code generated by MockGen. DO NOT EDIT.
// Source: repository/partnerrepository/partnerrepository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/wjuneo/bexs/entity"
)

// MockPartnerRepository is a mock of PartnerRepository interface.
type MockPartnerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPartnerRepositoryMockRecorder
}

// MockPartnerRepositoryMockRecorder is the mock recorder for MockPartnerRepository.
type MockPartnerRepositoryMockRecorder struct {
	mock *MockPartnerRepository
}

// NewMockPartnerRepository creates a new mock instance.
func NewMockPartnerRepository(ctrl *gomock.Controller) *MockPartnerRepository {
	mock := &MockPartnerRepository{ctrl: ctrl}
	mock.recorder = &MockPartnerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartnerRepository) EXPECT() *MockPartnerRepositoryMockRecorder {
	return m.recorder
}

// FindPartnerByDocument mocks base method.
func (m *MockPartnerRepository) FindPartnerByDocument(ctx context.Context, document string) (*entity.Partner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPartnerByDocument", ctx, document)
	ret0, _ := ret[0].(*entity.Partner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPartnerByDocument indicates an expected call of FindPartnerByDocument.
func (mr *MockPartnerRepositoryMockRecorder) FindPartnerByDocument(ctx, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPartnerByDocument", reflect.TypeOf((*MockPartnerRepository)(nil).FindPartnerByDocument), ctx, document)
}

// SavePartners mocks base method.
func (m *MockPartnerRepository) SavePartners(ctx context.Context, partner entity.Partner) (*entity.Partner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePartners", ctx, partner)
	ret0, _ := ret[0].(*entity.Partner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SavePartners indicates an expected call of SavePartners.
func (mr *MockPartnerRepositoryMockRecorder) SavePartners(ctx, partner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePartners", reflect.TypeOf((*MockPartnerRepository)(nil).SavePartners), ctx, partner)
}

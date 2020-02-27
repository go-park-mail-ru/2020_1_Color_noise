// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/pin/usecase/interface_usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "pinterest/pkg/models"
	reflect "reflect"
)

// MockIPinUsecase is a mock of IPinUsecase interface
type MockIPinUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIPinUsecaseMockRecorder
}

// MockIPinUsecaseMockRecorder is the mock recorder for MockIPinUsecase
type MockIPinUsecaseMockRecorder struct {
	mock *MockIPinUsecase
}

// NewMockIPinUsecase creates a new mock instance
func NewMockIPinUsecase(ctrl *gomock.Controller) *MockIPinUsecase {
	mock := &MockIPinUsecase{ctrl: ctrl}
	mock.recorder = &MockIPinUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPinUsecase) EXPECT() *MockIPinUsecaseMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockIPinUsecase) Add(pin *models.Pin) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", pin)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockIPinUsecaseMockRecorder) Add(pin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockIPinUsecase)(nil).Add), pin)
}

// Get mocks base method
func (m *MockIPinUsecase) Get(id uint) (*models.Pin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*models.Pin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIPinUsecaseMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIPinUsecase)(nil).Get), id)
}

// SaveImage mocks base method
func (m *MockIPinUsecase) SaveImage(pin *models.Pin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveImage", pin)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveImage indicates an expected call of SaveImage
func (mr *MockIPinUsecaseMockRecorder) SaveImage(pin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveImage", reflect.TypeOf((*MockIPinUsecase)(nil).SaveImage), pin)
}

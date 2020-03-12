// Code generated by MockGen. DO NOT EDIT.
// Source: ./user/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	models "2020_1_Color_noise/internal/models"
	bytes "bytes"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIUsecase is a mock of IUsecase interface
type MockIUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIUsecaseMockRecorder
}

// MockIUsecaseMockRecorder is the mock recorder for MockIUsecase
type MockIUsecaseMockRecorder struct {
	mock *MockIUsecase
}

// NewMockIUsecase creates a new mock instance
func NewMockIUsecase(ctrl *gomock.Controller) *MockIUsecase {
	mock := &MockIUsecase{ctrl: ctrl}
	mock.recorder = &MockIUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUsecase) EXPECT() *MockIUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIUsecase) Create(input *models.SignUpInput) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockIUsecaseMockRecorder) Create(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUsecase)(nil).Create), input)
}

// GetById mocks base method
func (m *MockIUsecase) GetById(id uint) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockIUsecaseMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIUsecase)(nil).GetById), id)
}

// GetByLogin mocks base method
func (m *MockIUsecase) GetByLogin(login string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByLogin", login)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByLogin indicates an expected call of GetByLogin
func (mr *MockIUsecaseMockRecorder) GetByLogin(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLogin", reflect.TypeOf((*MockIUsecase)(nil).GetByLogin), login)
}

// Update mocks base method
func (m *MockIUsecase) Update(id uint, input *models.UpdateInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockIUsecaseMockRecorder) Update(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIUsecase)(nil).Update), id, input)
}

// UpdatePassword mocks base method
func (m *MockIUsecase) UpdatePassword(id uint, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", id, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockIUsecaseMockRecorder) UpdatePassword(id, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockIUsecase)(nil).UpdatePassword), id, password)
}

// Delete mocks base method
func (m *MockIUsecase) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIUsecaseMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIUsecase)(nil).Delete), id)
}

// ComparePassword mocks base method
func (m *MockIUsecase) ComparePassword(user *models.User, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComparePassword", user, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// ComparePassword indicates an expected call of ComparePassword
func (mr *MockIUsecaseMockRecorder) ComparePassword(user, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComparePassword", reflect.TypeOf((*MockIUsecase)(nil).ComparePassword), user, password)
}

// UpdateAvatar mocks base method
func (m *MockIUsecase) UpdateAvatar(id uint, buffer *bytes.Buffer) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", id, buffer)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAvatar indicates an expected call of UpdateAvatar
func (mr *MockIUsecaseMockRecorder) UpdateAvatar(id, buffer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockIUsecase)(nil).UpdateAvatar), id, buffer)
}

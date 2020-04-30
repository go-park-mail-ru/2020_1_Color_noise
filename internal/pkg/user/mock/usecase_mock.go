// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/user/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	models "2020_1_Color_noise/internal/models"
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
func (m *MockIUsecase) Create(input *models.SignUpInput) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", input)
	ret0, _ := ret[0].(*models.User)
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

// UpdateProfile mocks base method
func (m *MockIUsecase) UpdateProfile(id uint, input *models.UpdateProfileInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockIUsecaseMockRecorder) UpdateProfile(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockIUsecase)(nil).UpdateProfile), id, input)
}

// UpdateDescription mocks base method
func (m *MockIUsecase) UpdateDescription(id uint, input *models.UpdateDescriptionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDescription", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDescription indicates an expected call of UpdateDescription
func (mr *MockIUsecaseMockRecorder) UpdateDescription(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescription", reflect.TypeOf((*MockIUsecase)(nil).UpdateDescription), id, input)
}

// UpdatePassword mocks base method
func (m *MockIUsecase) UpdatePassword(id uint, input *models.UpdatePasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", id, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockIUsecaseMockRecorder) UpdatePassword(id, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockIUsecase)(nil).UpdatePassword), id, input)
}

// UpdateAvatar mocks base method
func (m *MockIUsecase) UpdateAvatar(id uint, buffer []byte) (string, error) {
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

// Follow mocks base method
func (m *MockIUsecase) Follow(id, subId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", id, subId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Follow indicates an expected call of Follow
func (mr *MockIUsecaseMockRecorder) Follow(id, subId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockIUsecase)(nil).Follow), id, subId)
}

// Unfollow mocks base method
func (m *MockIUsecase) Unfollow(id, subId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", id, subId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unfollow indicates an expected call of Unfollow
func (mr *MockIUsecaseMockRecorder) Unfollow(id, subId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockIUsecase)(nil).Unfollow), id, subId)
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

// Search mocks base method
func (m *MockIUsecase) Search(login string, start, limit int) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", login, start, limit)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockIUsecaseMockRecorder) Search(login, start, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIUsecase)(nil).Search), login, start, limit)
}

// GetSubscribers mocks base method
func (m *MockIUsecase) GetSubscribers(id uint, start, limit int) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", id, start, limit)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers
func (mr *MockIUsecaseMockRecorder) GetSubscribers(id, start, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockIUsecase)(nil).GetSubscribers), id, start, limit)
}

// GetSubscriptions mocks base method
func (m *MockIUsecase) GetSubscriptions(id uint, start, limit int) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptions", id, start, limit)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions
func (mr *MockIUsecaseMockRecorder) GetSubscriptions(id, start, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptions", reflect.TypeOf((*MockIUsecase)(nil).GetSubscriptions), id, start, limit)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ./user.pb.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	user "2020_1_Color_noise/internal/pkg/proto/user"
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockUserServiceClient is a mock of UserServiceClient interface
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserServiceClient) Create(ctx context.Context, in *user.SignUp, opts ...grpc.CallOption) (*user.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserServiceClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserServiceClient)(nil).Create), varargs...)
}

// GetById mocks base method
func (m *MockUserServiceClient) GetById(ctx context.Context, in *user.UserID, opts ...grpc.CallOption) (*user.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetById", varargs...)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockUserServiceClientMockRecorder) GetById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserServiceClient)(nil).GetById), varargs...)
}

// GetByLogin mocks base method
func (m *MockUserServiceClient) GetByLogin(ctx context.Context, in *user.Login, opts ...grpc.CallOption) (*user.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByLogin", varargs...)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByLogin indicates an expected call of GetByLogin
func (mr *MockUserServiceClientMockRecorder) GetByLogin(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLogin", reflect.TypeOf((*MockUserServiceClient)(nil).GetByLogin), varargs...)
}

// UpdateProfile mocks base method
func (m *MockUserServiceClient) UpdateProfile(ctx context.Context, in *user.Profile, opts ...grpc.CallOption) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProfile", varargs...)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockUserServiceClientMockRecorder) UpdateProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserServiceClient)(nil).UpdateProfile), varargs...)
}

// UpdateDescription mocks base method
func (m *MockUserServiceClient) UpdateDescription(ctx context.Context, in *user.Description, opts ...grpc.CallOption) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDescription", varargs...)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDescription indicates an expected call of UpdateDescription
func (mr *MockUserServiceClientMockRecorder) UpdateDescription(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescription", reflect.TypeOf((*MockUserServiceClient)(nil).UpdateDescription), varargs...)
}

// UpdatePassword mocks base method
func (m *MockUserServiceClient) UpdatePassword(ctx context.Context, in *user.Password, opts ...grpc.CallOption) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePassword", varargs...)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockUserServiceClientMockRecorder) UpdatePassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserServiceClient)(nil).UpdatePassword), varargs...)
}

// UpdateAvatar mocks base method
func (m *MockUserServiceClient) UpdateAvatar(ctx context.Context, in *user.Avatar, opts ...grpc.CallOption) (*user.Address, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateAvatar", varargs...)
	ret0, _ := ret[0].(*user.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAvatar indicates an expected call of UpdateAvatar
func (mr *MockUserServiceClientMockRecorder) UpdateAvatar(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockUserServiceClient)(nil).UpdateAvatar), varargs...)
}

// Follow mocks base method
func (m *MockUserServiceClient) Follow(ctx context.Context, in *user.Following, opts ...grpc.CallOption) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Follow", varargs...)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Follow indicates an expected call of Follow
func (mr *MockUserServiceClientMockRecorder) Follow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockUserServiceClient)(nil).Follow), varargs...)
}

// Unfollow mocks base method
func (m *MockUserServiceClient) Unfollow(ctx context.Context, in *user.Following, opts ...grpc.CallOption) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Unfollow", varargs...)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow
func (mr *MockUserServiceClientMockRecorder) Unfollow(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockUserServiceClient)(nil).Unfollow), varargs...)
}

// Search mocks base method
func (m *MockUserServiceClient) Search(ctx context.Context, in *user.Searching, opts ...grpc.CallOption) (*user.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockUserServiceClientMockRecorder) Search(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUserServiceClient)(nil).Search), varargs...)
}

// GetSubscribers mocks base method
func (m *MockUserServiceClient) GetSubscribers(ctx context.Context, in *user.Sub, opts ...grpc.CallOption) (*user.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubscribers", varargs...)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers
func (mr *MockUserServiceClientMockRecorder) GetSubscribers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockUserServiceClient)(nil).GetSubscribers), varargs...)
}

// GetSubscriptions mocks base method
func (m *MockUserServiceClient) GetSubscriptions(ctx context.Context, in *user.Sub, opts ...grpc.CallOption) (*user.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubscriptions", varargs...)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions
func (mr *MockUserServiceClientMockRecorder) GetSubscriptions(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptions", reflect.TypeOf((*MockUserServiceClient)(nil).GetSubscriptions), varargs...)
}

// MockUserServiceServer is a mock of UserServiceServer interface
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserServiceServer) Create(arg0 context.Context, arg1 *user.SignUp) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserServiceServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserServiceServer)(nil).Create), arg0, arg1)
}

// GetById mocks base method
func (m *MockUserServiceServer) GetById(arg0 context.Context, arg1 *user.UserID) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockUserServiceServerMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUserServiceServer)(nil).GetById), arg0, arg1)
}

// GetByLogin mocks base method
func (m *MockUserServiceServer) GetByLogin(arg0 context.Context, arg1 *user.Login) (*user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByLogin", arg0, arg1)
	ret0, _ := ret[0].(*user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByLogin indicates an expected call of GetByLogin
func (mr *MockUserServiceServerMockRecorder) GetByLogin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByLogin", reflect.TypeOf((*MockUserServiceServer)(nil).GetByLogin), arg0, arg1)
}

// UpdateProfile mocks base method
func (m *MockUserServiceServer) UpdateProfile(arg0 context.Context, arg1 *user.Profile) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile
func (mr *MockUserServiceServerMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserServiceServer)(nil).UpdateProfile), arg0, arg1)
}

// UpdateDescription mocks base method
func (m *MockUserServiceServer) UpdateDescription(arg0 context.Context, arg1 *user.Description) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDescription", arg0, arg1)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDescription indicates an expected call of UpdateDescription
func (mr *MockUserServiceServerMockRecorder) UpdateDescription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDescription", reflect.TypeOf((*MockUserServiceServer)(nil).UpdateDescription), arg0, arg1)
}

// UpdatePassword mocks base method
func (m *MockUserServiceServer) UpdatePassword(arg0 context.Context, arg1 *user.Password) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePassword indicates an expected call of UpdatePassword
func (mr *MockUserServiceServerMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserServiceServer)(nil).UpdatePassword), arg0, arg1)
}

// UpdateAvatar mocks base method
func (m *MockUserServiceServer) UpdateAvatar(arg0 context.Context, arg1 *user.Avatar) (*user.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1)
	ret0, _ := ret[0].(*user.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAvatar indicates an expected call of UpdateAvatar
func (mr *MockUserServiceServerMockRecorder) UpdateAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*MockUserServiceServer)(nil).UpdateAvatar), arg0, arg1)
}

// Follow mocks base method
func (m *MockUserServiceServer) Follow(arg0 context.Context, arg1 *user.Following) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Follow", arg0, arg1)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Follow indicates an expected call of Follow
func (mr *MockUserServiceServerMockRecorder) Follow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Follow", reflect.TypeOf((*MockUserServiceServer)(nil).Follow), arg0, arg1)
}

// Unfollow mocks base method
func (m *MockUserServiceServer) Unfollow(arg0 context.Context, arg1 *user.Following) (*user.Nothing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unfollow", arg0, arg1)
	ret0, _ := ret[0].(*user.Nothing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unfollow indicates an expected call of Unfollow
func (mr *MockUserServiceServerMockRecorder) Unfollow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unfollow", reflect.TypeOf((*MockUserServiceServer)(nil).Unfollow), arg0, arg1)
}

// Search mocks base method
func (m *MockUserServiceServer) Search(arg0 context.Context, arg1 *user.Searching) (*user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0, arg1)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search
func (mr *MockUserServiceServerMockRecorder) Search(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUserServiceServer)(nil).Search), arg0, arg1)
}

// GetSubscribers mocks base method
func (m *MockUserServiceServer) GetSubscribers(arg0 context.Context, arg1 *user.Sub) (*user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscribers", arg0, arg1)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscribers indicates an expected call of GetSubscribers
func (mr *MockUserServiceServerMockRecorder) GetSubscribers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscribers", reflect.TypeOf((*MockUserServiceServer)(nil).GetSubscribers), arg0, arg1)
}

// GetSubscriptions mocks base method
func (m *MockUserServiceServer) GetSubscriptions(arg0 context.Context, arg1 *user.Sub) (*user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptions", arg0, arg1)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions
func (mr *MockUserServiceServerMockRecorder) GetSubscriptions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptions", reflect.TypeOf((*MockUserServiceServer)(nil).GetSubscriptions), arg0, arg1)
}

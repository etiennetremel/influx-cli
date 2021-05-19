// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influx-cli/v2/internal/api (interfaces: UsersApi)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	api "github.com/influxdata/influx-cli/v2/internal/api"
)

// MockUsersApi is a mock of UsersApi interface.
type MockUsersApi struct {
	ctrl     *gomock.Controller
	recorder *MockUsersApiMockRecorder
}

// MockUsersApiMockRecorder is the mock recorder for MockUsersApi.
type MockUsersApiMockRecorder struct {
	mock *MockUsersApi
}

// NewMockUsersApi creates a new mock instance.
func NewMockUsersApi(ctrl *gomock.Controller) *MockUsersApi {
	mock := &MockUsersApi{ctrl: ctrl}
	mock.recorder = &MockUsersApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersApi) EXPECT() *MockUsersApiMockRecorder {
	return m.recorder
}

// DeleteUsersID mocks base method.
func (m *MockUsersApi) DeleteUsersID(arg0 context.Context, arg1 string) api.ApiDeleteUsersIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsersID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiDeleteUsersIDRequest)
	return ret0
}

// DeleteUsersID indicates an expected call of DeleteUsersID.
func (mr *MockUsersApiMockRecorder) DeleteUsersID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsersID", reflect.TypeOf((*MockUsersApi)(nil).DeleteUsersID), arg0, arg1)
}

// DeleteUsersIDExecute mocks base method.
func (m *MockUsersApi) DeleteUsersIDExecute(arg0 api.ApiDeleteUsersIDRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsersIDExecute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsersIDExecute indicates an expected call of DeleteUsersIDExecute.
func (mr *MockUsersApiMockRecorder) DeleteUsersIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsersIDExecute", reflect.TypeOf((*MockUsersApi)(nil).DeleteUsersIDExecute), arg0)
}

// GetUsersID mocks base method.
func (m *MockUsersApi) GetUsersID(arg0 context.Context, arg1 string) api.ApiGetUsersIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiGetUsersIDRequest)
	return ret0
}

// GetUsersID indicates an expected call of GetUsersID.
func (mr *MockUsersApiMockRecorder) GetUsersID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersID", reflect.TypeOf((*MockUsersApi)(nil).GetUsersID), arg0, arg1)
}

// GetUsersIDExecute mocks base method.
func (m *MockUsersApi) GetUsersIDExecute(arg0 api.ApiGetUsersIDRequest) (api.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersIDExecute", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersIDExecute indicates an expected call of GetUsersIDExecute.
func (mr *MockUsersApiMockRecorder) GetUsersIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersIDExecute", reflect.TypeOf((*MockUsersApi)(nil).GetUsersIDExecute), arg0)
}

// PatchUsersID mocks base method.
func (m *MockUsersApi) PatchUsersID(arg0 context.Context, arg1 string) api.ApiPatchUsersIDRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchUsersID", arg0, arg1)
	ret0, _ := ret[0].(api.ApiPatchUsersIDRequest)
	return ret0
}

// PatchUsersID indicates an expected call of PatchUsersID.
func (mr *MockUsersApiMockRecorder) PatchUsersID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchUsersID", reflect.TypeOf((*MockUsersApi)(nil).PatchUsersID), arg0, arg1)
}

// PatchUsersIDExecute mocks base method.
func (m *MockUsersApi) PatchUsersIDExecute(arg0 api.ApiPatchUsersIDRequest) (api.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchUsersIDExecute", arg0)
	ret0, _ := ret[0].(api.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PatchUsersIDExecute indicates an expected call of PatchUsersIDExecute.
func (mr *MockUsersApiMockRecorder) PatchUsersIDExecute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchUsersIDExecute", reflect.TypeOf((*MockUsersApi)(nil).PatchUsersIDExecute), arg0)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: stub/github.com/nileshdt/pb/logs/logs.trpc.go

// Package logs is a generated GoMock package.
package logs

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "trpc.group/trpc-go/trpc-go/client"
)

// MockLogServiceService is a mock of LogServiceService interface.
type MockLogServiceService struct {
	ctrl     *gomock.Controller
	recorder *MockLogServiceServiceMockRecorder
}

// MockLogServiceServiceMockRecorder is the mock recorder for MockLogServiceService.
type MockLogServiceServiceMockRecorder struct {
	mock *MockLogServiceService
}

// NewMockLogServiceService creates a new mock instance.
func NewMockLogServiceService(ctrl *gomock.Controller) *MockLogServiceService {
	mock := &MockLogServiceService{ctrl: ctrl}
	mock.recorder = &MockLogServiceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogServiceService) EXPECT() *MockLogServiceServiceMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockLogServiceService) Log(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log", ctx, req)
	ret0, _ := ret[0].(*LogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Log indicates an expected call of Log.
func (mr *MockLogServiceServiceMockRecorder) Log(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogServiceService)(nil).Log), ctx, req)
}

// MockLogServiceClientProxy is a mock of LogServiceClientProxy interface.
type MockLogServiceClientProxy struct {
	ctrl     *gomock.Controller
	recorder *MockLogServiceClientProxyMockRecorder
}

// MockLogServiceClientProxyMockRecorder is the mock recorder for MockLogServiceClientProxy.
type MockLogServiceClientProxyMockRecorder struct {
	mock *MockLogServiceClientProxy
}

// NewMockLogServiceClientProxy creates a new mock instance.
func NewMockLogServiceClientProxy(ctrl *gomock.Controller) *MockLogServiceClientProxy {
	mock := &MockLogServiceClientProxy{ctrl: ctrl}
	mock.recorder = &MockLogServiceClientProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogServiceClientProxy) EXPECT() *MockLogServiceClientProxyMockRecorder {
	return m.recorder
}

// Log mocks base method.
func (m *MockLogServiceClientProxy) Log(ctx context.Context, req *LogRequest, opts ...client.Option) (*LogResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Log", varargs...)
	ret0, _ := ret[0].(*LogResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Log indicates an expected call of Log.
func (mr *MockLogServiceClientProxyMockRecorder) Log(ctx, req interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogServiceClientProxy)(nil).Log), varargs...)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go

// Package debugger_test is a generated GoMock package.
package debugger_test

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServer is a mock of Server interface
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// Serve mocks base method
func (m *MockServer) Serve(arg0 net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve
func (mr *MockServerMockRecorder) Serve(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockServer)(nil).Serve), arg0)
}

// RegisterOnShutdown mocks base method
func (m *MockServer) RegisterOnShutdown(arg0 func()) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterOnShutdown", arg0)
}

// RegisterOnShutdown indicates an expected call of RegisterOnShutdown
func (mr *MockServerMockRecorder) RegisterOnShutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOnShutdown", reflect.TypeOf((*MockServer)(nil).RegisterOnShutdown), arg0)
}

// Shutdown mocks base method
func (m *MockServer) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown
func (mr *MockServerMockRecorder) Shutdown(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockServer)(nil).Shutdown), arg0)
}

// MockListener is a mock of Listener interface
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
}

// MockListenerMockRecorder is the mock recorder for MockListener
type MockListenerMockRecorder struct {
	mock *MockListener
}

// NewMockListener creates a new mock instance
func NewMockListener(ctrl *gomock.Controller) *MockListener {
	mock := &MockListener{ctrl: ctrl}
	mock.recorder = &MockListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockListener) EXPECT() *MockListenerMockRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockListener) Accept() (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accept indicates an expected call of Accept
func (mr *MockListenerMockRecorder) Accept() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockListener)(nil).Accept))
}

// Close mocks base method
func (m *MockListener) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockListenerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockListener)(nil).Close))
}

// Addr mocks base method
func (m *MockListener) Addr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// Addr indicates an expected call of Addr
func (mr *MockListenerMockRecorder) Addr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockListener)(nil).Addr))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ../internal/transaction/interface.go
//
// Generated by this command:
//
//	mockgen.exe -destination=../internal/transaction/mocks.go -package=transaction -source=../internal/transaction/interface.go
//
// Package transaction is a generated GoMock package.
package transaction

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSession is a mock of Session interface.
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession.
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance.
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockSession) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockSessionMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockSession)(nil).Commit))
}

// CreateNewSession mocks base method.
func (m *MockSession) CreateNewSession() Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewSession")
	ret0, _ := ret[0].(Session)
	return ret0
}

// CreateNewSession indicates an expected call of CreateNewSession.
func (mr *MockSessionMockRecorder) CreateNewSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewSession", reflect.TypeOf((*MockSession)(nil).CreateNewSession))
}

// Rollback mocks base method.
func (m *MockSession) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockSessionMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockSession)(nil).Rollback))
}

// Start mocks base method.
func (m *MockSession) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockSessionMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSession)(nil).Start))
}

// Tx mocks base method.
func (m *MockSession) Tx() any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx")
	ret0, _ := ret[0].(any)
	return ret0
}

// Tx indicates an expected call of Tx.
func (mr *MockSessionMockRecorder) Tx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockSession)(nil).Tx))
}

// TxIsActive mocks base method.
func (m *MockSession) TxIsActive() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxIsActive")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TxIsActive indicates an expected call of TxIsActive.
func (mr *MockSessionMockRecorder) TxIsActive() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxIsActive", reflect.TypeOf((*MockSession)(nil).TxIsActive))
}

// MockSessionManager is a mock of SessionManager interface.
type MockSessionManager struct {
	ctrl     *gomock.Controller
	recorder *MockSessionManagerMockRecorder
}

// MockSessionManagerMockRecorder is the mock recorder for MockSessionManager.
type MockSessionManagerMockRecorder struct {
	mock *MockSessionManager
}

// NewMockSessionManager creates a new mock instance.
func NewMockSessionManager(ctrl *gomock.Controller) *MockSessionManager {
	mock := &MockSessionManager{ctrl: ctrl}
	mock.recorder = &MockSessionManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionManager) EXPECT() *MockSessionManagerMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockSessionManager) CreateSession() Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession")
	ret0, _ := ret[0].(Session)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionManagerMockRecorder) CreateSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessionManager)(nil).CreateSession))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: pow.go

// Package pow is a generated GoMock package.
package pow

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVerifier is a mock of Verifier interface.
type MockVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockVerifierMockRecorder
}

// MockVerifierMockRecorder is the mock recorder for MockVerifier.
type MockVerifierMockRecorder struct {
	mock *MockVerifier
}

// NewMockVerifier creates a new mock instance.
func NewMockVerifier(ctrl *gomock.Controller) *MockVerifier {
	mock := &MockVerifier{ctrl: ctrl}
	mock.recorder = &MockVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVerifier) EXPECT() *MockVerifierMockRecorder {
	return m.recorder
}

// Challenge mocks base method.
func (m *MockVerifier) Challenge() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Challenge")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Challenge indicates an expected call of Challenge.
func (mr *MockVerifierMockRecorder) Challenge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Challenge", reflect.TypeOf((*MockVerifier)(nil).Challenge))
}

// Verify mocks base method.
func (m *MockVerifier) Verify(challenge, solution []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", challenge, solution)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockVerifierMockRecorder) Verify(challenge, solution interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockVerifier)(nil).Verify), challenge, solution)
}

// MockSolver is a mock of Solver interface.
type MockSolver struct {
	ctrl     *gomock.Controller
	recorder *MockSolverMockRecorder
}

// MockSolverMockRecorder is the mock recorder for MockSolver.
type MockSolverMockRecorder struct {
	mock *MockSolver
}

// NewMockSolver creates a new mock instance.
func NewMockSolver(ctrl *gomock.Controller) *MockSolver {
	mock := &MockSolver{ctrl: ctrl}
	mock.recorder = &MockSolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSolver) EXPECT() *MockSolverMockRecorder {
	return m.recorder
}

// Solve mocks base method.
func (m *MockSolver) Solve(challenge []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Solve", challenge)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Solve indicates an expected call of Solve.
func (mr *MockSolverMockRecorder) Solve(challenge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Solve", reflect.TypeOf((*MockSolver)(nil).Solve), challenge)
}

// MockPOW is a mock of POW interface.
type MockPOW struct {
	ctrl     *gomock.Controller
	recorder *MockPOWMockRecorder
}

// MockPOWMockRecorder is the mock recorder for MockPOW.
type MockPOWMockRecorder struct {
	mock *MockPOW
}

// NewMockPOW creates a new mock instance.
func NewMockPOW(ctrl *gomock.Controller) *MockPOW {
	mock := &MockPOW{ctrl: ctrl}
	mock.recorder = &MockPOWMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPOW) EXPECT() *MockPOWMockRecorder {
	return m.recorder
}

// Challenge mocks base method.
func (m *MockPOW) Challenge() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Challenge")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Challenge indicates an expected call of Challenge.
func (mr *MockPOWMockRecorder) Challenge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Challenge", reflect.TypeOf((*MockPOW)(nil).Challenge))
}

// Solve mocks base method.
func (m *MockPOW) Solve(challenge []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Solve", challenge)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Solve indicates an expected call of Solve.
func (mr *MockPOWMockRecorder) Solve(challenge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Solve", reflect.TypeOf((*MockPOW)(nil).Solve), challenge)
}

// Verify mocks base method.
func (m *MockPOW) Verify(challenge, solution []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", challenge, solution)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockPOWMockRecorder) Verify(challenge, solution interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockPOW)(nil).Verify), challenge, solution)
}

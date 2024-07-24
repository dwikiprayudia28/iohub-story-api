// Code generated by MockGen. DO NOT EDIT.
// Source: uuid.go

// Package mockutils is a generated GoMock package.
package mockutils

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUUID is a mock of UUID interface.
type MockUUID struct {
	ctrl     *gomock.Controller
	recorder *MockUUIDMockRecorder
}

// MockUUIDMockRecorder is the mock recorder for MockUUID.
type MockUUIDMockRecorder struct {
	mock *MockUUID
}

// NewMockUUID creates a new mock instance.
func NewMockUUID(ctrl *gomock.Controller) *MockUUID {
	mock := &MockUUID{ctrl: ctrl}
	mock.recorder = &MockUUIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUUID) EXPECT() *MockUUIDMockRecorder {
	return m.recorder
}

// NewUUID mocks base method.
func (m *MockUUID) NewUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// NewUUID indicates an expected call of NewUUID.
func (mr *MockUUIDMockRecorder) NewUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUUID", reflect.TypeOf((*MockUUID)(nil).NewUUID))
}

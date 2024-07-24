// Code generated by MockGen. DO NOT EDIT.
// Source: message_repository.go

// Package mockchatrepository is a generated GoMock package.
package mockchatrepository

import (
	reflect "reflect"

	dto "github.com/iammuho/natternet/internal/chat/application/dto"
	values "github.com/iammuho/natternet/internal/chat/domain/values"
	errorhandler "github.com/iammuho/natternet/pkg/errorhandler"
	gomock "go.uber.org/mock/gomock"
)

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageRepository) Create(message *values.MessageDBValue) *errorhandler.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", message)
	ret0, _ := ret[0].(*errorhandler.Response)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMessageRepositoryMockRecorder) Create(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageRepository)(nil).Create), message)
}

// Query mocks base method.
func (m *MockMessageRepository) Query(query *dto.QueryMessagesReqDTO) ([]*values.MessageValue, *errorhandler.Response) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", query)
	ret0, _ := ret[0].([]*values.MessageValue)
	ret1, _ := ret[1].(*errorhandler.Response)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockMessageRepositoryMockRecorder) Query(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockMessageRepository)(nil).Query), query)
}

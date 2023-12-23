// Code generated by MockGen. DO NOT EDIT.
// Source: chatsettings.go
//
// Generated by this command:
//
//	mockgen -source=chatsettings.go -destination=mocks/mock.go
//

// Package mock_chatsettings is a generated GoMock package.
package mock_chatsettings

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	domain "github.com/satont/twitch-notifier/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepository) Create(ctx context.Context, chatSettings domain.ChatSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, chatSettings)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(ctx, chatSettings any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), ctx, chatSettings)
}

// Delete mocks base method.
func (m *MockRepository) Delete(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRepository)(nil).Delete), ctx, id)
}

// GetByChatID mocks base method.
func (m *MockRepository) GetByChatID(ctx context.Context, chatID uuid.UUID) (*domain.ChatSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByChatID", ctx, chatID)
	ret0, _ := ret[0].(*domain.ChatSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByChatID indicates an expected call of GetByChatID.
func (mr *MockRepositoryMockRecorder) GetByChatID(ctx, chatID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByChatID", reflect.TypeOf((*MockRepository)(nil).GetByChatID), ctx, chatID)
}

// GetByID mocks base method.
func (m *MockRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ChatSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*domain.ChatSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRepositoryMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRepository)(nil).GetByID), ctx, id)
}

// Update mocks base method.
func (m *MockRepository) Update(ctx context.Context, chatSettings domain.ChatSettings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, chatSettings)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRepositoryMockRecorder) Update(ctx, chatSettings any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), ctx, chatSettings)
}

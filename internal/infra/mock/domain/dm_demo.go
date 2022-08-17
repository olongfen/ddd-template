// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/./domain/dm_demo.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "ddd-template/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIDemoRepo is a mock of IDemoRepo interface.
type MockIDemoRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIDemoRepoMockRecorder
}

// MockIDemoRepoMockRecorder is the mock recorder for MockIDemoRepo.
type MockIDemoRepoMockRecorder struct {
	mock *MockIDemoRepo
}

// NewMockIDemoRepo creates a new mock instance.
func NewMockIDemoRepo(ctrl *gomock.Controller) *MockIDemoRepo {
	mock := &MockIDemoRepo{ctrl: ctrl}
	mock.recorder = &MockIDemoRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDemoRepo) EXPECT() *MockIDemoRepoMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIDemoRepo) Get(ctx context.Context, demo *domain.Demo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, demo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockIDemoRepoMockRecorder) Get(ctx, demo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIDemoRepo)(nil).Get), ctx, demo)
}

// MockIDemoUseCase is a mock of IDemoUseCase interface.
type MockIDemoUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIDemoUseCaseMockRecorder
}

// MockIDemoUseCaseMockRecorder is the mock recorder for MockIDemoUseCase.
type MockIDemoUseCaseMockRecorder struct {
	mock *MockIDemoUseCase
}

// NewMockIDemoUseCase creates a new mock instance.
func NewMockIDemoUseCase(ctrl *gomock.Controller) *MockIDemoUseCase {
	mock := &MockIDemoUseCase{ctrl: ctrl}
	mock.recorder = &MockIDemoUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDemoUseCase) EXPECT() *MockIDemoUseCaseMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIDemoUseCase) Get(ctx context.Context, id uint) (*domain.Demo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*domain.Demo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIDemoUseCaseMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIDemoUseCase)(nil).Get), ctx, id)
}
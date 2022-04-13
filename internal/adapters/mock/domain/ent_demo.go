// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/./domain/ent_demo.go

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

// SayHello mocks base method.
func (m *MockIDemoRepo) SayHello(ctx context.Context, msg string) *domain.Demo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SayHello", ctx, msg)
	ret0, _ := ret[0].(*domain.Demo)
	return ret0
}

// SayHello indicates an expected call of SayHello.
func (mr *MockIDemoRepoMockRecorder) SayHello(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockIDemoRepo)(nil).SayHello), ctx, msg)
}

// MockIDemoUsecase is a mock of IDemoUsecase interface.
type MockIDemoUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIDemoUsecaseMockRecorder
}

// MockIDemoUsecaseMockRecorder is the mock recorder for MockIDemoUsecase.
type MockIDemoUsecaseMockRecorder struct {
	mock *MockIDemoUsecase
}

// NewMockIDemoUsecase creates a new mock instance.
func NewMockIDemoUsecase(ctrl *gomock.Controller) *MockIDemoUsecase {
	mock := &MockIDemoUsecase{ctrl: ctrl}
	mock.recorder = &MockIDemoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDemoUsecase) EXPECT() *MockIDemoUsecaseMockRecorder {
	return m.recorder
}

// SayHello mocks base method.
func (m *MockIDemoUsecase) SayHello(ctx context.Context, msg string) (*domain.Demo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SayHello", ctx, msg)
	ret0, _ := ret[0].(*domain.Demo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello.
func (mr *MockIDemoUsecaseMockRecorder) SayHello(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockIDemoUsecase)(nil).SayHello), ctx, msg)
}

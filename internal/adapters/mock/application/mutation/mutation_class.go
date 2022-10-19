// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/./application/mutation/mutation_class.go

// Package mock_mutation is a generated GoMock package.
package mock_mutation

import (
	context "context"
	schema "ddd-template/internal/schema"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIClassMutationService is a mock of IClassMutationService interface.
type MockIClassMutationService struct {
	ctrl     *gomock.Controller
	recorder *MockIClassMutationServiceMockRecorder
}

// MockIClassMutationServiceMockRecorder is the mock recorder for MockIClassMutationService.
type MockIClassMutationServiceMockRecorder struct {
	mock *MockIClassMutationService
}

// NewMockIClassMutationService creates a new mock instance.
func NewMockIClassMutationService(ctrl *gomock.Controller) *MockIClassMutationService {
	mock := &MockIClassMutationService{ctrl: ctrl}
	mock.recorder = &MockIClassMutationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIClassMutationService) EXPECT() *MockIClassMutationServiceMockRecorder {
	return m.recorder
}

// AddClass mocks base method.
func (m *MockIClassMutationService) AddClass(ctx context.Context, form *schema.ClassAddForm) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, form)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddClass indicates an expected call of AddClass.
func (mr *MockIClassMutationServiceMockRecorder) AddClass(ctx, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIClassMutationService)(nil).AddClass), ctx, form)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: repository/tree/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=repository/tree/interfaces.go -destination=repository/tree/interfaces.mock.gen.go -package=tree
//

// Package tree is a generated GoMock package.
package tree

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockTreeRepositoryInterface is a mock of TreeRepositoryInterface interface.
type MockTreeRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTreeRepositoryInterfaceMockRecorder
}

// MockTreeRepositoryInterfaceMockRecorder is the mock recorder for MockTreeRepositoryInterface.
type MockTreeRepositoryInterfaceMockRecorder struct {
	mock *MockTreeRepositoryInterface
}

// NewMockTreeRepositoryInterface creates a new mock instance.
func NewMockTreeRepositoryInterface(ctrl *gomock.Controller) *MockTreeRepositoryInterface {
	mock := &MockTreeRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockTreeRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTreeRepositoryInterface) EXPECT() *MockTreeRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateTree mocks base method.
func (m *MockTreeRepositoryInterface) CreateTree(ctx context.Context, tree *Tree) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", ctx, tree)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockTreeRepositoryInterfaceMockRecorder) CreateTree(ctx, tree any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockTreeRepositoryInterface)(nil).CreateTree), ctx, tree)
}

// GetTree mocks base method.
func (m *MockTreeRepositoryInterface) GetTree(ctx context.Context, id uuid.UUID) (*Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", ctx, id)
	ret0, _ := ret[0].(*Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockTreeRepositoryInterfaceMockRecorder) GetTree(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockTreeRepositoryInterface)(nil).GetTree), ctx, id)
}

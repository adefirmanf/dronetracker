// Code generated by MockGen. DO NOT EDIT.
// Source: service/tree/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=service/tree/interfaces.go -destination=service/tree/interfaces.mock.gen.go -package=tree
//

// Package tree is a generated GoMock package.
package tree

import (
	context "context"
	reflect "reflect"

	tree "github.com/SawitProRecruitment/UserService/repository/tree"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateNewTree mocks base method.
func (m *MockService) CreateNewTree(ctx context.Context, estateId string, xCoordinate, yCoordinate, height int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewTree", ctx, estateId, xCoordinate, yCoordinate, height)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewTree indicates an expected call of CreateNewTree.
func (mr *MockServiceMockRecorder) CreateNewTree(ctx, estateId, xCoordinate, yCoordinate, height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewTree", reflect.TypeOf((*MockService)(nil).CreateNewTree), ctx, estateId, xCoordinate, yCoordinate, height)
}

// GetStats mocks base method.
func (m *MockService) GetStats(ctx context.Context, id string) (*tree.TreeStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStats", ctx, id)
	ret0, _ := ret[0].(*tree.TreeStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStats indicates an expected call of GetStats.
func (mr *MockServiceMockRecorder) GetStats(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStats", reflect.TypeOf((*MockService)(nil).GetStats), ctx, id)
}

// RetrievesByEstateID mocks base method.
func (m *MockService) RetrievesByEstateID(ctx context.Context, id string) ([]*tree.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RetrievesByEstateID", ctx, id)
	ret0, _ := ret[0].([]*tree.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RetrievesByEstateID indicates an expected call of RetrievesByEstateID.
func (mr *MockServiceMockRecorder) RetrievesByEstateID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrievesByEstateID", reflect.TypeOf((*MockService)(nil).RetrievesByEstateID), ctx, id)
}
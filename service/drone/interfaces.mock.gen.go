// Code generated by MockGen. DO NOT EDIT.
// Source: service/drone/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=service/drone/interfaces.go -destination=service/drone/interfaces.mock.gen.go -package=drone
//

// Package drone is a generated GoMock package.
package drone

import (
	reflect "reflect"

	estate "github.com/SawitProRecruitment/UserService/repository/estate"
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

// GetDronePlane mocks base method.
func (m *MockService) GetDronePlane(estate *estate.Estate, tree []*tree.Tree, maxDistance int) (int, Coordinate) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDronePlane", estate, tree, maxDistance)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(Coordinate)
	return ret0, ret1
}

// GetDronePlane indicates an expected call of GetDronePlane.
func (mr *MockServiceMockRecorder) GetDronePlane(estate, tree, maxDistance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDronePlane", reflect.TypeOf((*MockService)(nil).GetDronePlane), estate, tree, maxDistance)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: parser.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	calculation "github.com/nvasilev98/calculator/pkg/calculation"
	reflect "reflect"
)

// MockOperationFactorier is a mock of OperationFactorier interface
type MockOperationFactorier struct {
	ctrl     *gomock.Controller
	recorder *MockOperationFactorierMockRecorder
}

// MockOperationFactorierMockRecorder is the mock recorder for MockOperationFactorier
type MockOperationFactorierMockRecorder struct {
	mock *MockOperationFactorier
}

// NewMockOperationFactorier creates a new mock instance
func NewMockOperationFactorier(ctrl *gomock.Controller) *MockOperationFactorier {
	mock := &MockOperationFactorier{ctrl: ctrl}
	mock.recorder = &MockOperationFactorierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOperationFactorier) EXPECT() *MockOperationFactorierMockRecorder {
	return m.recorder
}

// GetOperation mocks base method
func (m *MockOperationFactorier) GetOperation(str string) (calculation.Evaluator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperation", str)
	ret0, _ := ret[0].(calculation.Evaluator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation
func (mr *MockOperationFactorierMockRecorder) GetOperation(str interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockOperationFactorier)(nil).GetOperation), str)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	postgres "github.com/nvasilev98/calculator/pkg/database/postgres"
	reflect "reflect"
)

// MockCalculator is a mock of Calculator interface
type MockCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorMockRecorder
}

// MockCalculatorMockRecorder is the mock recorder for MockCalculator
type MockCalculatorMockRecorder struct {
	mock *MockCalculator
}

// NewMockCalculator creates a new mock instance
func NewMockCalculator(ctrl *gomock.Controller) *MockCalculator {
	mock := &MockCalculator{ctrl: ctrl}
	mock.recorder = &MockCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCalculator) EXPECT() *MockCalculatorMockRecorder {
	return m.recorder
}

// Calculate mocks base method
func (m *MockCalculator) Calculate(expr string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Calculate", expr)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Calculate indicates an expected call of Calculate
func (mr *MockCalculatorMockRecorder) Calculate(expr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Calculate", reflect.TypeOf((*MockCalculator)(nil).Calculate), expr)
}

// MockDataAccess is a mock of DataAccess interface
type MockDataAccess struct {
	ctrl     *gomock.Controller
	recorder *MockDataAccessMockRecorder
}

// MockDataAccessMockRecorder is the mock recorder for MockDataAccess
type MockDataAccessMockRecorder struct {
	mock *MockDataAccess
}

// NewMockDataAccess creates a new mock instance
func NewMockDataAccess(ctrl *gomock.Controller) *MockDataAccess {
	mock := &MockDataAccess{ctrl: ctrl}
	mock.recorder = &MockDataAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataAccess) EXPECT() *MockDataAccessMockRecorder {
	return m.recorder
}

// InsertCalculation mocks base method
func (m *MockDataAccess) InsertCalculation(model postgres.Model) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCalculation", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertCalculation indicates an expected call of InsertCalculation
func (mr *MockDataAccessMockRecorder) InsertCalculation(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCalculation", reflect.TypeOf((*MockDataAccess)(nil).InsertCalculation), model)
}

// SelectResultByID mocks base method
func (m *MockDataAccess) SelectResultByID(id int64) (postgres.Response, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectResultByID", id)
	ret0, _ := ret[0].(postgres.Response)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SelectResultByID indicates an expected call of SelectResultByID
func (mr *MockDataAccessMockRecorder) SelectResultByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectResultByID", reflect.TypeOf((*MockDataAccess)(nil).SelectResultByID), id)
}

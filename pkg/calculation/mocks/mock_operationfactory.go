// Code generated by MockGen. DO NOT EDIT.
// Source: operationfactory.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEvaluator is a mock of Evaluator interface
type MockEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockEvaluatorMockRecorder
}

// MockEvaluatorMockRecorder is the mock recorder for MockEvaluator
type MockEvaluatorMockRecorder struct {
	mock *MockEvaluator
}

// NewMockEvaluator creates a new mock instance
func NewMockEvaluator(ctrl *gomock.Controller) *MockEvaluator {
	mock := &MockEvaluator{ctrl: ctrl}
	mock.recorder = &MockEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEvaluator) EXPECT() *MockEvaluatorMockRecorder {
	return m.recorder
}

// Evaluate mocks base method
func (m *MockEvaluator) Evaluate(num1, num2 float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Evaluate", num1, num2)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Evaluate indicates an expected call of Evaluate
func (mr *MockEvaluatorMockRecorder) Evaluate(num1, num2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Evaluate", reflect.TypeOf((*MockEvaluator)(nil).Evaluate), num1, num2)
}

// Weight mocks base method
func (m *MockEvaluator) Weight() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Weight")
	ret0, _ := ret[0].(int)
	return ret0
}

// Weight indicates an expected call of Weight
func (mr *MockEvaluatorMockRecorder) Weight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Weight", reflect.TypeOf((*MockEvaluator)(nil).Weight))
}

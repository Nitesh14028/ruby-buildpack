// Code generated by MockGen. DO NOT EDIT.
// Source: finalize.go

package finalize_test

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockStager) EXPECT() *MockStagerMockRecorder {
	return _m.recorder
}

// BuildDir mocks base method
func (_m *MockStager) BuildDir() string {
	ret := _m.ctrl.Call(_m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDir indicates an expected call of BuildDir
func (_mr *MockStagerMockRecorder) BuildDir() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "BuildDir", reflect.TypeOf((*MockStager)(nil).BuildDir))
}

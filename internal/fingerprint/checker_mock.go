// Code generated by MockGen. DO NOT EDIT.
// Source: checker.go

// Package fingerprint is a generated GoMock package.
package fingerprint

import (
	context "context"
	reflect "reflect"

	taskfile "github.com/go-task/task/v3/taskfile"
	gomock "github.com/golang/mock/gomock"
)

// MockStatusCheckable is a mock of StatusCheckable interface.
type MockStatusCheckable struct {
	ctrl     *gomock.Controller
	recorder *MockStatusCheckableMockRecorder
}

// MockStatusCheckableMockRecorder is the mock recorder for MockStatusCheckable.
type MockStatusCheckableMockRecorder struct {
	mock *MockStatusCheckable
}

// NewMockStatusCheckable creates a new mock instance.
func NewMockStatusCheckable(ctrl *gomock.Controller) *MockStatusCheckable {
	mock := &MockStatusCheckable{ctrl: ctrl}
	mock.recorder = &MockStatusCheckableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusCheckable) EXPECT() *MockStatusCheckableMockRecorder {
	return m.recorder
}

// IsUpToDate mocks base method.
func (m *MockStatusCheckable) IsUpToDate(ctx context.Context, t *taskfile.Task) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpToDate", ctx, t)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUpToDate indicates an expected call of IsUpToDate.
func (mr *MockStatusCheckableMockRecorder) IsUpToDate(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpToDate", reflect.TypeOf((*MockStatusCheckable)(nil).IsUpToDate), ctx, t)
}

// MockSourcesCheckable is a mock of SourcesCheckable interface.
type MockSourcesCheckable struct {
	ctrl     *gomock.Controller
	recorder *MockSourcesCheckableMockRecorder
}

// MockSourcesCheckableMockRecorder is the mock recorder for MockSourcesCheckable.
type MockSourcesCheckableMockRecorder struct {
	mock *MockSourcesCheckable
}

// NewMockSourcesCheckable creates a new mock instance.
func NewMockSourcesCheckable(ctrl *gomock.Controller) *MockSourcesCheckable {
	mock := &MockSourcesCheckable{ctrl: ctrl}
	mock.recorder = &MockSourcesCheckableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourcesCheckable) EXPECT() *MockSourcesCheckableMockRecorder {
	return m.recorder
}

// IsUpToDate mocks base method.
func (m *MockSourcesCheckable) IsUpToDate(t *taskfile.Task) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUpToDate", t)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUpToDate indicates an expected call of IsUpToDate.
func (mr *MockSourcesCheckableMockRecorder) IsUpToDate(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUpToDate", reflect.TypeOf((*MockSourcesCheckable)(nil).IsUpToDate), t)
}

// Kind mocks base method.
func (m *MockSourcesCheckable) Kind() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(string)
	return ret0
}

// Kind indicates an expected call of Kind.
func (mr *MockSourcesCheckableMockRecorder) Kind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockSourcesCheckable)(nil).Kind))
}

// OnError mocks base method.
func (m *MockSourcesCheckable) OnError(t *taskfile.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnError", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// OnError indicates an expected call of OnError.
func (mr *MockSourcesCheckableMockRecorder) OnError(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockSourcesCheckable)(nil).OnError), t)
}

// Value mocks base method.
func (m *MockSourcesCheckable) Value(t *taskfile.Task) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value", t)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Value indicates an expected call of Value.
func (mr *MockSourcesCheckableMockRecorder) Value(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockSourcesCheckable)(nil).Value), t)
}
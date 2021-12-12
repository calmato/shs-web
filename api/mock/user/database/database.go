// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	database "github.com/calmato/shs-web/api/internal/user/database"
	entity "github.com/calmato/shs-web/api/internal/user/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTeacher is a mock of Teacher interface.
type MockTeacher struct {
	ctrl     *gomock.Controller
	recorder *MockTeacherMockRecorder
}

// MockTeacherMockRecorder is the mock recorder for MockTeacher.
type MockTeacherMockRecorder struct {
	mock *MockTeacher
}

// NewMockTeacher creates a new mock instance.
func NewMockTeacher(ctrl *gomock.Controller) *MockTeacher {
	mock := &MockTeacher{ctrl: ctrl}
	mock.recorder = &MockTeacherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeacher) EXPECT() *MockTeacherMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTeacher) Create(ctx context.Context, t *entity.Teacher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTeacherMockRecorder) Create(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTeacher)(nil).Create), ctx, t)
}

// Get mocks base method.
func (m *MockTeacher) Get(ctx context.Context, id string, fields ...string) (*entity.Teacher, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Teacher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTeacherMockRecorder) Get(ctx, id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTeacher)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockTeacher) List(ctx context.Context, p *database.ListTeachersParams, fields ...string) (entity.Teachers, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, p}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Teachers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockTeacherMockRecorder) List(ctx, p interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, p}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockTeacher)(nil).List), varargs...)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"
	time "time"

	database "github.com/calmato/shs-web/api/internal/classroom/database"
	entity "github.com/calmato/shs-web/api/internal/classroom/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockSubject is a mock of Subject interface.
type MockSubject struct {
	ctrl     *gomock.Controller
	recorder *MockSubjectMockRecorder
}

// MockSubjectMockRecorder is the mock recorder for MockSubject.
type MockSubjectMockRecorder struct {
	mock *MockSubject
}

// NewMockSubject creates a new mock instance.
func NewMockSubject(ctrl *gomock.Controller) *MockSubject {
	mock := &MockSubject{ctrl: ctrl}
	mock.recorder = &MockSubjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubject) EXPECT() *MockSubjectMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockSubject) Count(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSubjectMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSubject)(nil).Count), ctx)
}

// Get mocks base method.
func (m *MockSubject) Get(ctx context.Context, id int64, fields ...string) (*entity.Subject, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Subject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSubjectMockRecorder) Get(ctx, id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSubject)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockSubject) List(ctx context.Context, p *database.ListSubjectsParams, fields ...string) (entity.Subjects, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, p}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Subjects)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSubjectMockRecorder) List(ctx, p interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, p}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSubject)(nil).List), varargs...)
}

// MultiGet mocks base method.
func (m *MockSubject) MultiGet(ctx context.Context, ids []int64, fields ...string) (entity.Subjects, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.Subjects)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockSubjectMockRecorder) MultiGet(ctx, ids interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockSubject)(nil).MultiGet), varargs...)
}

// MockTeacherSubject is a mock of TeacherSubject interface.
type MockTeacherSubject struct {
	ctrl     *gomock.Controller
	recorder *MockTeacherSubjectMockRecorder
}

// MockTeacherSubjectMockRecorder is the mock recorder for MockTeacherSubject.
type MockTeacherSubjectMockRecorder struct {
	mock *MockTeacherSubject
}

// NewMockTeacherSubject creates a new mock instance.
func NewMockTeacherSubject(ctrl *gomock.Controller) *MockTeacherSubject {
	mock := &MockTeacherSubject{ctrl: ctrl}
	mock.recorder = &MockTeacherSubjectMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeacherSubject) EXPECT() *MockTeacherSubjectMockRecorder {
	return m.recorder
}

// ListByTeacherIDs mocks base method.
func (m *MockTeacherSubject) ListByTeacherIDs(ctx context.Context, teacherIDs []string, fields ...string) (entity.TeacherSubjects, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, teacherIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByTeacherIDs", varargs...)
	ret0, _ := ret[0].(entity.TeacherSubjects)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByTeacherIDs indicates an expected call of ListByTeacherIDs.
func (mr *MockTeacherSubjectMockRecorder) ListByTeacherIDs(ctx, teacherIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, teacherIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByTeacherIDs", reflect.TypeOf((*MockTeacherSubject)(nil).ListByTeacherIDs), varargs...)
}

// Replace mocks base method.
func (m *MockTeacherSubject) Replace(ctx context.Context, schoolType entity.SchoolType, subjects entity.TeacherSubjects) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", ctx, schoolType, subjects)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replace indicates an expected call of Replace.
func (mr *MockTeacherSubjectMockRecorder) Replace(ctx, schoolType, subjects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockTeacherSubject)(nil).Replace), ctx, schoolType, subjects)
}

// MockSchedule is a mock of Schedule interface.
type MockSchedule struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleMockRecorder
}

// MockScheduleMockRecorder is the mock recorder for MockSchedule.
type MockScheduleMockRecorder struct {
	mock *MockSchedule
}

// NewMockSchedule creates a new mock instance.
func NewMockSchedule(ctrl *gomock.Controller) *MockSchedule {
	mock := &MockSchedule{ctrl: ctrl}
	mock.recorder = &MockScheduleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedule) EXPECT() *MockScheduleMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSchedule) Get(ctx context.Context, weekday time.Weekday, fields ...string) (*entity.Schedule, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, weekday}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Schedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockScheduleMockRecorder) Get(ctx, weekday interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, weekday}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSchedule)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockSchedule) List(ctx context.Context, fields ...string) (entity.Schedules, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Schedules)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockScheduleMockRecorder) List(ctx interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSchedule)(nil).List), varargs...)
}

// MultipleUpdate mocks base method.
func (m *MockSchedule) MultipleUpdate(ctx context.Context, schedules entity.Schedules) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipleUpdate", ctx, schedules)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultipleUpdate indicates an expected call of MultipleUpdate.
func (mr *MockScheduleMockRecorder) MultipleUpdate(ctx, schedules interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipleUpdate", reflect.TypeOf((*MockSchedule)(nil).MultipleUpdate), ctx, schedules)
}

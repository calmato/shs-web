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

// MockStudent is a mock of Student interface.
type MockStudent struct {
	ctrl     *gomock.Controller
	recorder *MockStudentMockRecorder
}

// MockStudentMockRecorder is the mock recorder for MockStudent.
type MockStudentMockRecorder struct {
	mock *MockStudent
}

// NewMockStudent creates a new mock instance.
func NewMockStudent(ctrl *gomock.Controller) *MockStudent {
	mock := &MockStudent{ctrl: ctrl}
	mock.recorder = &MockStudentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudent) EXPECT() *MockStudentMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStudent) Create(ctx context.Context, s *entity.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStudentMockRecorder) Create(ctx, s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStudent)(nil).Create), ctx, s)
}

// Get mocks base method.
func (m *MockStudent) Get(ctx context.Context, id string, fields ...string) (*entity.Student, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStudentMockRecorder) Get(ctx, id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStudent)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockStudent) List(ctx context.Context, p *database.ListStudentsParams, fields ...string) (entity.Students, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, p}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Students)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStudentMockRecorder) List(ctx, p interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, p}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStudent)(nil).List), varargs...)
}

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

// Count mocks base method.
func (m *MockTeacher) Count(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockTeacherMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockTeacher)(nil).Count), ctx)
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

// Delete mocks base method.
func (m *MockTeacher) Delete(ctx context.Context, teacherID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, teacherID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTeacherMockRecorder) Delete(ctx, teacherID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTeacher)(nil).Delete), ctx, teacherID)
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

// UpdateMail mocks base method.
func (m *MockTeacher) UpdateMail(ctx context.Context, teacherID, mail string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMail", ctx, teacherID, mail)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMail indicates an expected call of UpdateMail.
func (mr *MockTeacherMockRecorder) UpdateMail(ctx, teacherID, mail interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMail", reflect.TypeOf((*MockTeacher)(nil).UpdateMail), ctx, teacherID, mail)
}

// UpdatePassword mocks base method.
func (m *MockTeacher) UpdatePassword(ctx context.Context, teacherID, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, teacherID, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockTeacherMockRecorder) UpdatePassword(ctx, teacherID, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockTeacher)(nil).UpdatePassword), ctx, teacherID, password)
}

// UpdateRole mocks base method.
func (m *MockTeacher) UpdateRole(ctx context.Context, teacherID string, role entity.Role) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", ctx, teacherID, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockTeacherMockRecorder) UpdateRole(ctx, teacherID, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockTeacher)(nil).UpdateRole), ctx, teacherID, role)
}

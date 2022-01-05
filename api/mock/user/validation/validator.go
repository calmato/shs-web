// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	user "github.com/calmato/shs-web/api/proto/user"
	gomock "github.com/golang/mock/gomock"
)

// MockRequestValidation is a mock of RequestValidation interface.
type MockRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockRequestValidationMockRecorder
}

// MockRequestValidationMockRecorder is the mock recorder for MockRequestValidation.
type MockRequestValidationMockRecorder struct {
	mock *MockRequestValidation
}

// NewMockRequestValidation creates a new mock instance.
func NewMockRequestValidation(ctrl *gomock.Controller) *MockRequestValidation {
	mock := &MockRequestValidation{ctrl: ctrl}
	mock.recorder = &MockRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequestValidation) EXPECT() *MockRequestValidationMockRecorder {
	return m.recorder
}

// CreateTeacher mocks base method.
func (m *MockRequestValidation) CreateTeacher(req *user.CreateTeacherRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeacher", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTeacher indicates an expected call of CreateTeacher.
func (mr *MockRequestValidationMockRecorder) CreateTeacher(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeacher", reflect.TypeOf((*MockRequestValidation)(nil).CreateTeacher), req)
}

// DeleteTeacher mocks base method.
func (m *MockRequestValidation) DeleteTeacher(req *user.DeleteTeacherRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeacher", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTeacher indicates an expected call of DeleteTeacher.
func (mr *MockRequestValidationMockRecorder) DeleteTeacher(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeacher", reflect.TypeOf((*MockRequestValidation)(nil).DeleteTeacher), req)
}

// GetStudent mocks base method.
func (m *MockRequestValidation) GetStudent(req *user.GetStudentRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudent", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetStudent indicates an expected call of GetStudent.
func (mr *MockRequestValidationMockRecorder) GetStudent(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudent", reflect.TypeOf((*MockRequestValidation)(nil).GetStudent), req)
}

// GetTeacher mocks base method.
func (m *MockRequestValidation) GetTeacher(req *user.GetTeacherRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeacher", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTeacher indicates an expected call of GetTeacher.
func (mr *MockRequestValidationMockRecorder) GetTeacher(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacher", reflect.TypeOf((*MockRequestValidation)(nil).GetTeacher), req)
}

// ListTeachers mocks base method.
func (m *MockRequestValidation) ListTeachers(req *user.ListTeachersRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeachers", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTeachers indicates an expected call of ListTeachers.
func (mr *MockRequestValidationMockRecorder) ListTeachers(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeachers", reflect.TypeOf((*MockRequestValidation)(nil).ListTeachers), req)
}

// UpdateTeacherMail mocks base method.
func (m *MockRequestValidation) UpdateTeacherMail(req *user.UpdateTeacherMailRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacherMail", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTeacherMail indicates an expected call of UpdateTeacherMail.
func (mr *MockRequestValidationMockRecorder) UpdateTeacherMail(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherMail", reflect.TypeOf((*MockRequestValidation)(nil).UpdateTeacherMail), req)
}

// UpdateTeacherPassword mocks base method.
func (m *MockRequestValidation) UpdateTeacherPassword(req *user.UpdateTeacherPasswordRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacherPassword", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTeacherPassword indicates an expected call of UpdateTeacherPassword.
func (mr *MockRequestValidationMockRecorder) UpdateTeacherPassword(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherPassword", reflect.TypeOf((*MockRequestValidation)(nil).UpdateTeacherPassword), req)
}

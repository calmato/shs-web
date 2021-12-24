// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	classroom "github.com/calmato/shs-web/api/proto/classroom"
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

// GetSchedule mocks base method.
func (m *MockRequestValidation) GetSchedule(req *classroom.GetScheduleRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetSchedule indicates an expected call of GetSchedule.
func (mr *MockRequestValidationMockRecorder) GetSchedule(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockRequestValidation)(nil).GetSchedule), req)
}

// GetSubject mocks base method.
func (m *MockRequestValidation) GetSubject(req *classroom.GetSubjectRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubject", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetSubject indicates an expected call of GetSubject.
func (mr *MockRequestValidationMockRecorder) GetSubject(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubject", reflect.TypeOf((*MockRequestValidation)(nil).GetSubject), req)
}

// GetTeacherSubject mocks base method.
func (m *MockRequestValidation) GetTeacherSubject(req *classroom.GetTeacherSubjectRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeacherSubject", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTeacherSubject indicates an expected call of GetTeacherSubject.
func (mr *MockRequestValidationMockRecorder) GetTeacherSubject(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherSubject", reflect.TypeOf((*MockRequestValidation)(nil).GetTeacherSubject), req)
}

// ListSchedules mocks base method.
func (m *MockRequestValidation) ListSchedules(req *classroom.ListSchedulesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedules", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListSchedules indicates an expected call of ListSchedules.
func (mr *MockRequestValidationMockRecorder) ListSchedules(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedules", reflect.TypeOf((*MockRequestValidation)(nil).ListSchedules), req)
}

// ListSubjects mocks base method.
func (m *MockRequestValidation) ListSubjects(req *classroom.ListSubjectsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjects", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListSubjects indicates an expected call of ListSubjects.
func (mr *MockRequestValidationMockRecorder) ListSubjects(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjects", reflect.TypeOf((*MockRequestValidation)(nil).ListSubjects), req)
}

// MultiGetSubjects mocks base method.
func (m *MockRequestValidation) MultiGetSubjects(req *classroom.MultiGetSubjectsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetSubjects", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultiGetSubjects indicates an expected call of MultiGetSubjects.
func (mr *MockRequestValidationMockRecorder) MultiGetSubjects(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetSubjects", reflect.TypeOf((*MockRequestValidation)(nil).MultiGetSubjects), req)
}

// MultiGetTeacherSubjects mocks base method.
func (m *MockRequestValidation) MultiGetTeacherSubjects(req *classroom.MultiGetTeacherSubjectsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetTeacherSubjects", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultiGetTeacherSubjects indicates an expected call of MultiGetTeacherSubjects.
func (mr *MockRequestValidationMockRecorder) MultiGetTeacherSubjects(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetTeacherSubjects", reflect.TypeOf((*MockRequestValidation)(nil).MultiGetTeacherSubjects), req)
}

// UpdateTeacherSubject mocks base method.
func (m *MockRequestValidation) UpdateTeacherSubject(req *classroom.UpdateTeacherSubjectRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacherSubject", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTeacherSubject indicates an expected call of UpdateTeacherSubject.
func (mr *MockRequestValidationMockRecorder) UpdateTeacherSubject(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherSubject", reflect.TypeOf((*MockRequestValidation)(nil).UpdateTeacherSubject), req)
}

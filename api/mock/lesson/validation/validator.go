// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	reflect "reflect"

	lesson "github.com/calmato/shs-web/api/proto/lesson"
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

// CreateLesson mocks base method.
func (m *MockRequestValidation) CreateLesson(req *lesson.CreateLessonRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLesson", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateLesson indicates an expected call of CreateLesson.
func (mr *MockRequestValidationMockRecorder) CreateLesson(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLesson", reflect.TypeOf((*MockRequestValidation)(nil).CreateLesson), req)
}

// CreateShifts mocks base method.
func (m *MockRequestValidation) CreateShifts(req *lesson.CreateShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateShifts indicates an expected call of CreateShifts.
func (mr *MockRequestValidationMockRecorder) CreateShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShifts", reflect.TypeOf((*MockRequestValidation)(nil).CreateShifts), req)
}

// DeleteLesson mocks base method.
func (m *MockRequestValidation) DeleteLesson(req *lesson.DeleteLessonRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLesson", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLesson indicates an expected call of DeleteLesson.
func (mr *MockRequestValidationMockRecorder) DeleteLesson(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLesson", reflect.TypeOf((*MockRequestValidation)(nil).DeleteLesson), req)
}

// DeleteShiftSummary mocks base method.
func (m *MockRequestValidation) DeleteShiftSummary(req *lesson.DeleteShiftSummaryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteShiftSummary", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteShiftSummary indicates an expected call of DeleteShiftSummary.
func (mr *MockRequestValidationMockRecorder) DeleteShiftSummary(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShiftSummary", reflect.TypeOf((*MockRequestValidation)(nil).DeleteShiftSummary), req)
}

// GetShiftSummary mocks base method.
func (m *MockRequestValidation) GetShiftSummary(req *lesson.GetShiftSummaryRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShiftSummary", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetShiftSummary indicates an expected call of GetShiftSummary.
func (mr *MockRequestValidationMockRecorder) GetShiftSummary(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShiftSummary", reflect.TypeOf((*MockRequestValidation)(nil).GetShiftSummary), req)
}

// GetStudentShiftTemplate mocks base method.
func (m *MockRequestValidation) GetStudentShiftTemplate(req *lesson.GetStudentShiftTemplateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentShiftTemplate", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetStudentShiftTemplate indicates an expected call of GetStudentShiftTemplate.
func (mr *MockRequestValidationMockRecorder) GetStudentShiftTemplate(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentShiftTemplate", reflect.TypeOf((*MockRequestValidation)(nil).GetStudentShiftTemplate), req)
}

// GetStudentShifts mocks base method.
func (m *MockRequestValidation) GetStudentShifts(req *lesson.GetStudentShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetStudentShifts indicates an expected call of GetStudentShifts.
func (mr *MockRequestValidationMockRecorder) GetStudentShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentShifts", reflect.TypeOf((*MockRequestValidation)(nil).GetStudentShifts), req)
}

// GetTeacherShifts mocks base method.
func (m *MockRequestValidation) GetTeacherShifts(req *lesson.GetTeacherShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeacherShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetTeacherShifts indicates an expected call of GetTeacherShifts.
func (mr *MockRequestValidationMockRecorder) GetTeacherShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherShifts", reflect.TypeOf((*MockRequestValidation)(nil).GetTeacherShifts), req)
}

// ListLessons mocks base method.
func (m *MockRequestValidation) ListLessons(req *lesson.ListLessonsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLessons", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListLessons indicates an expected call of ListLessons.
func (mr *MockRequestValidationMockRecorder) ListLessons(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLessons", reflect.TypeOf((*MockRequestValidation)(nil).ListLessons), req)
}

// ListLessonsByDuration mocks base method.
func (m *MockRequestValidation) ListLessonsByDuration(req *lesson.ListLessonsByDurationRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLessonsByDuration", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListLessonsByDuration indicates an expected call of ListLessonsByDuration.
func (mr *MockRequestValidationMockRecorder) ListLessonsByDuration(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLessonsByDuration", reflect.TypeOf((*MockRequestValidation)(nil).ListLessonsByDuration), req)
}

// ListShiftSummaries mocks base method.
func (m *MockRequestValidation) ListShiftSummaries(req *lesson.ListShiftSummariesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShiftSummaries", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListShiftSummaries indicates an expected call of ListShiftSummaries.
func (mr *MockRequestValidationMockRecorder) ListShiftSummaries(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShiftSummaries", reflect.TypeOf((*MockRequestValidation)(nil).ListShiftSummaries), req)
}

// ListShifts mocks base method.
func (m *MockRequestValidation) ListShifts(req *lesson.ListShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListShifts indicates an expected call of ListShifts.
func (mr *MockRequestValidationMockRecorder) ListShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShifts", reflect.TypeOf((*MockRequestValidation)(nil).ListShifts), req)
}

// ListStudentShifts mocks base method.
func (m *MockRequestValidation) ListStudentShifts(req *lesson.ListStudentShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudentShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListStudentShifts indicates an expected call of ListStudentShifts.
func (mr *MockRequestValidationMockRecorder) ListStudentShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentShifts", reflect.TypeOf((*MockRequestValidation)(nil).ListStudentShifts), req)
}

// ListStudentSubmissionsByShiftSummaryIDs mocks base method.
func (m *MockRequestValidation) ListStudentSubmissionsByShiftSummaryIDs(req *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudentSubmissionsByShiftSummaryIDs", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListStudentSubmissionsByShiftSummaryIDs indicates an expected call of ListStudentSubmissionsByShiftSummaryIDs.
func (mr *MockRequestValidationMockRecorder) ListStudentSubmissionsByShiftSummaryIDs(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentSubmissionsByShiftSummaryIDs", reflect.TypeOf((*MockRequestValidation)(nil).ListStudentSubmissionsByShiftSummaryIDs), req)
}

// ListStudentSubmissionsByStudentIDs mocks base method.
func (m *MockRequestValidation) ListStudentSubmissionsByStudentIDs(req *lesson.ListStudentSubmissionsByStudentIDsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudentSubmissionsByStudentIDs", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListStudentSubmissionsByStudentIDs indicates an expected call of ListStudentSubmissionsByStudentIDs.
func (mr *MockRequestValidationMockRecorder) ListStudentSubmissionsByStudentIDs(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentSubmissionsByStudentIDs", reflect.TypeOf((*MockRequestValidation)(nil).ListStudentSubmissionsByStudentIDs), req)
}

// ListSubmissions mocks base method.
func (m *MockRequestValidation) ListSubmissions(req *lesson.ListSubmissionsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubmissions", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListSubmissions indicates an expected call of ListSubmissions.
func (mr *MockRequestValidationMockRecorder) ListSubmissions(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubmissions", reflect.TypeOf((*MockRequestValidation)(nil).ListSubmissions), req)
}

// ListTeacherShifts mocks base method.
func (m *MockRequestValidation) ListTeacherShifts(req *lesson.ListTeacherShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeacherShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTeacherShifts indicates an expected call of ListTeacherShifts.
func (mr *MockRequestValidationMockRecorder) ListTeacherShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherShifts", reflect.TypeOf((*MockRequestValidation)(nil).ListTeacherShifts), req)
}

// ListTeacherSubmissionsByShiftSummaryIDs mocks base method.
func (m *MockRequestValidation) ListTeacherSubmissionsByShiftSummaryIDs(req *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeacherSubmissionsByShiftSummaryIDs", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTeacherSubmissionsByShiftSummaryIDs indicates an expected call of ListTeacherSubmissionsByShiftSummaryIDs.
func (mr *MockRequestValidationMockRecorder) ListTeacherSubmissionsByShiftSummaryIDs(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherSubmissionsByShiftSummaryIDs", reflect.TypeOf((*MockRequestValidation)(nil).ListTeacherSubmissionsByShiftSummaryIDs), req)
}

// ListTeacherSubmissionsByTeacherIDs mocks base method.
func (m *MockRequestValidation) ListTeacherSubmissionsByTeacherIDs(req *lesson.ListTeacherSubmissionsByTeacherIDsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeacherSubmissionsByTeacherIDs", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListTeacherSubmissionsByTeacherIDs indicates an expected call of ListTeacherSubmissionsByTeacherIDs.
func (mr *MockRequestValidationMockRecorder) ListTeacherSubmissionsByTeacherIDs(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherSubmissionsByTeacherIDs", reflect.TypeOf((*MockRequestValidation)(nil).ListTeacherSubmissionsByTeacherIDs), req)
}

// UpdateLesson mocks base method.
func (m *MockRequestValidation) UpdateLesson(req *lesson.UpdateLessonRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLesson", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLesson indicates an expected call of UpdateLesson.
func (mr *MockRequestValidationMockRecorder) UpdateLesson(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLesson", reflect.TypeOf((*MockRequestValidation)(nil).UpdateLesson), req)
}

// UpdateShiftSummaryDecided mocks base method.
func (m *MockRequestValidation) UpdateShiftSummaryDecided(req *lesson.UpdateShiftSummaryDecidedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShiftSummaryDecided", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateShiftSummaryDecided indicates an expected call of UpdateShiftSummaryDecided.
func (mr *MockRequestValidationMockRecorder) UpdateShiftSummaryDecided(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShiftSummaryDecided", reflect.TypeOf((*MockRequestValidation)(nil).UpdateShiftSummaryDecided), req)
}

// UpdateShiftSummarySchedule mocks base method.
func (m *MockRequestValidation) UpdateShiftSummarySchedule(req *lesson.UpdateShiftSummaryScheduleRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShiftSummarySchedule", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateShiftSummarySchedule indicates an expected call of UpdateShiftSummarySchedule.
func (mr *MockRequestValidationMockRecorder) UpdateShiftSummarySchedule(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShiftSummarySchedule", reflect.TypeOf((*MockRequestValidation)(nil).UpdateShiftSummarySchedule), req)
}

// UpsertStudentShiftTemplate mocks base method.
func (m *MockRequestValidation) UpsertStudentShiftTemplate(req *lesson.UpsertStudentShiftTemplateRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertStudentShiftTemplate", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertStudentShiftTemplate indicates an expected call of UpsertStudentShiftTemplate.
func (mr *MockRequestValidationMockRecorder) UpsertStudentShiftTemplate(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertStudentShiftTemplate", reflect.TypeOf((*MockRequestValidation)(nil).UpsertStudentShiftTemplate), req)
}

// UpsertStudentShifts mocks base method.
func (m *MockRequestValidation) UpsertStudentShifts(req *lesson.UpsertStudentShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertStudentShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertStudentShifts indicates an expected call of UpsertStudentShifts.
func (mr *MockRequestValidationMockRecorder) UpsertStudentShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertStudentShifts", reflect.TypeOf((*MockRequestValidation)(nil).UpsertStudentShifts), req)
}

// UpsertTeacherShifts mocks base method.
func (m *MockRequestValidation) UpsertTeacherShifts(req *lesson.UpsertTeacherShiftsRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertTeacherShifts", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertTeacherShifts indicates an expected call of UpsertTeacherShifts.
func (mr *MockRequestValidationMockRecorder) UpsertTeacherShifts(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTeacherShifts", reflect.TypeOf((*MockRequestValidation)(nil).UpsertTeacherShifts), req)
}

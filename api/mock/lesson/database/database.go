// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"
	time "time"

	database "github.com/calmato/shs-web/api/internal/lesson/database"
	entity "github.com/calmato/shs-web/api/internal/lesson/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockShiftSummary is a mock of ShiftSummary interface.
type MockShiftSummary struct {
	ctrl     *gomock.Controller
	recorder *MockShiftSummaryMockRecorder
}

// MockShiftSummaryMockRecorder is the mock recorder for MockShiftSummary.
type MockShiftSummaryMockRecorder struct {
	mock *MockShiftSummary
}

// NewMockShiftSummary creates a new mock instance.
func NewMockShiftSummary(ctrl *gomock.Controller) *MockShiftSummary {
	mock := &MockShiftSummary{ctrl: ctrl}
	mock.recorder = &MockShiftSummaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShiftSummary) EXPECT() *MockShiftSummaryMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockShiftSummary) Count(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockShiftSummaryMockRecorder) Count(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockShiftSummary)(nil).Count), ctx)
}

// Delete mocks base method.
func (m *MockShiftSummary) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockShiftSummaryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockShiftSummary)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockShiftSummary) Get(ctx context.Context, id int64, fields ...string) (*entity.ShiftSummary, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.ShiftSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShiftSummaryMockRecorder) Get(ctx, id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShiftSummary)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockShiftSummary) List(ctx context.Context, p *database.ListShiftSummariesParams, fields ...string) (entity.ShiftSummaries, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, p}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.ShiftSummaries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockShiftSummaryMockRecorder) List(ctx, p interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, p}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockShiftSummary)(nil).List), varargs...)
}

// UpdateSchedule mocks base method.
func (m *MockShiftSummary) UpdateSchedule(ctx context.Context, id int64, openAt, endAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchedule", ctx, id, openAt, endAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSchedule indicates an expected call of UpdateSchedule.
func (mr *MockShiftSummaryMockRecorder) UpdateSchedule(ctx, id, openAt, endAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedule", reflect.TypeOf((*MockShiftSummary)(nil).UpdateSchedule), ctx, id, openAt, endAt)
}

// MockShift is a mock of Shift interface.
type MockShift struct {
	ctrl     *gomock.Controller
	recorder *MockShiftMockRecorder
}

// MockShiftMockRecorder is the mock recorder for MockShift.
type MockShiftMockRecorder struct {
	mock *MockShift
}

// NewMockShift creates a new mock instance.
func NewMockShift(ctrl *gomock.Controller) *MockShift {
	mock := &MockShift{ctrl: ctrl}
	mock.recorder = &MockShiftMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShift) EXPECT() *MockShiftMockRecorder {
	return m.recorder
}

// ListBySummaryID mocks base method.
func (m *MockShift) ListBySummaryID(ctx context.Context, summaryID int64, fields ...string) (entity.Shifts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, summaryID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBySummaryID", varargs...)
	ret0, _ := ret[0].(entity.Shifts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBySummaryID indicates an expected call of ListBySummaryID.
func (mr *MockShiftMockRecorder) ListBySummaryID(ctx, summaryID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, summaryID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySummaryID", reflect.TypeOf((*MockShift)(nil).ListBySummaryID), varargs...)
}

// MultiGet mocks base method.
func (m *MockShift) MultiGet(ctx context.Context, ids []int64, fields ...string) (entity.Shifts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.Shifts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockShiftMockRecorder) MultiGet(ctx, ids interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockShift)(nil).MultiGet), varargs...)
}

// MultipleCreate mocks base method.
func (m *MockShift) MultipleCreate(ctx context.Context, summary *entity.ShiftSummary, shifts entity.Shifts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultipleCreate", ctx, summary, shifts)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultipleCreate indicates an expected call of MultipleCreate.
func (mr *MockShiftMockRecorder) MultipleCreate(ctx, summary, shifts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultipleCreate", reflect.TypeOf((*MockShift)(nil).MultipleCreate), ctx, summary, shifts)
}

// MockTeacherShift is a mock of TeacherShift interface.
type MockTeacherShift struct {
	ctrl     *gomock.Controller
	recorder *MockTeacherShiftMockRecorder
}

// MockTeacherShiftMockRecorder is the mock recorder for MockTeacherShift.
type MockTeacherShiftMockRecorder struct {
	mock *MockTeacherShift
}

// NewMockTeacherShift creates a new mock instance.
func NewMockTeacherShift(ctrl *gomock.Controller) *MockTeacherShift {
	mock := &MockTeacherShift{ctrl: ctrl}
	mock.recorder = &MockTeacherShiftMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeacherShift) EXPECT() *MockTeacherShiftMockRecorder {
	return m.recorder
}

// Replace mocks base method.
func (m *MockTeacherShift) Replace(ctx context.Context, submission *entity.TeacherSubmission, shifts entity.TeacherShifts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", ctx, submission, shifts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replace indicates an expected call of Replace.
func (mr *MockTeacherShiftMockRecorder) Replace(ctx, submission, shifts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockTeacherShift)(nil).Replace), ctx, submission, shifts)
}

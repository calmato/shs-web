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

// MockLesson is a mock of Lesson interface.
type MockLesson struct {
	ctrl     *gomock.Controller
	recorder *MockLessonMockRecorder
}

// MockLessonMockRecorder is the mock recorder for MockLesson.
type MockLessonMockRecorder struct {
	mock *MockLesson
}

// NewMockLesson creates a new mock instance.
func NewMockLesson(ctrl *gomock.Controller) *MockLesson {
	mock := &MockLesson{ctrl: ctrl}
	mock.recorder = &MockLessonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLesson) EXPECT() *MockLessonMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockLesson) Count(ctx context.Context, p *database.ListLessonsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, p)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockLessonMockRecorder) Count(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockLesson)(nil).Count), ctx, p)
}

// Create mocks base method.
func (m *MockLesson) Create(ctx context.Context, lesson *entity.Lesson) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, lesson)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockLessonMockRecorder) Create(ctx, lesson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockLesson)(nil).Create), ctx, lesson)
}

// Delete mocks base method.
func (m *MockLesson) Delete(ctx context.Context, lessonID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, lessonID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockLessonMockRecorder) Delete(ctx, lessonID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLesson)(nil).Delete), ctx, lessonID)
}

// List mocks base method.
func (m *MockLesson) List(ctx context.Context, p *database.ListLessonsParams, fields ...string) (entity.Lessons, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, p}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Lessons)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockLessonMockRecorder) List(ctx, p interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, p}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLesson)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockLesson) Update(ctx context.Context, lessonID int64, lesson *entity.Lesson) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, lessonID, lesson)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockLessonMockRecorder) Update(ctx, lessonID, lesson interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLesson)(nil).Update), ctx, lessonID, lesson)
}

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

// MultiGet mocks base method.
func (m *MockShiftSummary) MultiGet(ctx context.Context, ids []int64, fields ...string) (entity.ShiftSummaries, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, ids}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.ShiftSummaries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockShiftSummaryMockRecorder) MultiGet(ctx, ids interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, ids}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockShiftSummary)(nil).MultiGet), varargs...)
}

// UpdateDecided mocks base method.
func (m *MockShiftSummary) UpdateDecided(ctx context.Context, id int64, decided bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDecided", ctx, id, decided)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDecided indicates an expected call of UpdateDecided.
func (mr *MockShiftSummaryMockRecorder) UpdateDecided(ctx, id, decided interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDecided", reflect.TypeOf((*MockShiftSummary)(nil).UpdateDecided), ctx, id, decided)
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

// Get mocks base method.
func (m *MockShift) Get(ctx context.Context, id int64, fields ...string) (*entity.Shift, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, id}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Shift)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShiftMockRecorder) Get(ctx, id interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, id}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShift)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockShift) List(ctx context.Context, p *database.ListShiftsParams, fields ...string) (entity.Shifts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, p}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Shifts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockShiftMockRecorder) List(ctx, p interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, p}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockShift)(nil).List), varargs...)
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

// MockTeacherSubmission is a mock of TeacherSubmission interface.
type MockTeacherSubmission struct {
	ctrl     *gomock.Controller
	recorder *MockTeacherSubmissionMockRecorder
}

// MockTeacherSubmissionMockRecorder is the mock recorder for MockTeacherSubmission.
type MockTeacherSubmissionMockRecorder struct {
	mock *MockTeacherSubmission
}

// NewMockTeacherSubmission creates a new mock instance.
func NewMockTeacherSubmission(ctrl *gomock.Controller) *MockTeacherSubmission {
	mock := &MockTeacherSubmission{ctrl: ctrl}
	mock.recorder = &MockTeacherSubmissionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeacherSubmission) EXPECT() *MockTeacherSubmissionMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTeacherSubmission) Get(ctx context.Context, teacherID string, summaryID int64, fields ...string) (*entity.TeacherSubmission, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, teacherID, summaryID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.TeacherSubmission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTeacherSubmissionMockRecorder) Get(ctx, teacherID, summaryID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, teacherID, summaryID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTeacherSubmission)(nil).Get), varargs...)
}

// ListByShiftSummaryIDs mocks base method.
func (m *MockTeacherSubmission) ListByShiftSummaryIDs(ctx context.Context, teacherID string, summaryIDs []int64, fields ...string) (entity.TeacherSubmissions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, teacherID, summaryIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByShiftSummaryIDs", varargs...)
	ret0, _ := ret[0].(entity.TeacherSubmissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShiftSummaryIDs indicates an expected call of ListByShiftSummaryIDs.
func (mr *MockTeacherSubmissionMockRecorder) ListByShiftSummaryIDs(ctx, teacherID, summaryIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, teacherID, summaryIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShiftSummaryIDs", reflect.TypeOf((*MockTeacherSubmission)(nil).ListByShiftSummaryIDs), varargs...)
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

// ListByShiftID mocks base method.
func (m *MockTeacherShift) ListByShiftID(ctx context.Context, shiftID int64, fields ...string) (entity.TeacherShifts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, shiftID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByShiftID", varargs...)
	ret0, _ := ret[0].(entity.TeacherShifts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShiftID indicates an expected call of ListByShiftID.
func (mr *MockTeacherShiftMockRecorder) ListByShiftID(ctx, shiftID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, shiftID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShiftID", reflect.TypeOf((*MockTeacherShift)(nil).ListByShiftID), varargs...)
}

// ListByShiftSummaryID mocks base method.
func (m *MockTeacherShift) ListByShiftSummaryID(ctx context.Context, teacherIDs []string, summaryID int64, fields ...string) (entity.TeacherShifts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, teacherIDs, summaryID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByShiftSummaryID", varargs...)
	ret0, _ := ret[0].(entity.TeacherShifts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShiftSummaryID indicates an expected call of ListByShiftSummaryID.
func (mr *MockTeacherShiftMockRecorder) ListByShiftSummaryID(ctx, teacherIDs, summaryID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, teacherIDs, summaryID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShiftSummaryID", reflect.TypeOf((*MockTeacherShift)(nil).ListByShiftSummaryID), varargs...)
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

// MockStudentSubmission is a mock of StudentSubmission interface.
type MockStudentSubmission struct {
	ctrl     *gomock.Controller
	recorder *MockStudentSubmissionMockRecorder
}

// MockStudentSubmissionMockRecorder is the mock recorder for MockStudentSubmission.
type MockStudentSubmissionMockRecorder struct {
	mock *MockStudentSubmission
}

// NewMockStudentSubmission creates a new mock instance.
func NewMockStudentSubmission(ctrl *gomock.Controller) *MockStudentSubmission {
	mock := &MockStudentSubmission{ctrl: ctrl}
	mock.recorder = &MockStudentSubmissionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentSubmission) EXPECT() *MockStudentSubmissionMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockStudentSubmission) Get(ctx context.Context, studentID string, summaryID int64, fields ...string) (*entity.StudentSubmission, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, studentID, summaryID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.StudentSubmission)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStudentSubmissionMockRecorder) Get(ctx, studentID, summaryID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, studentID, summaryID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStudentSubmission)(nil).Get), varargs...)
}

// ListByShiftSummaryIDs mocks base method.
func (m *MockStudentSubmission) ListByShiftSummaryIDs(ctx context.Context, studentID string, summaryIDs []int64, fields ...string) (entity.StudentSubmissions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, studentID, summaryIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByShiftSummaryIDs", varargs...)
	ret0, _ := ret[0].(entity.StudentSubmissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShiftSummaryIDs indicates an expected call of ListByShiftSummaryIDs.
func (mr *MockStudentSubmissionMockRecorder) ListByShiftSummaryIDs(ctx, studentID, summaryIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, studentID, summaryIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShiftSummaryIDs", reflect.TypeOf((*MockStudentSubmission)(nil).ListByShiftSummaryIDs), varargs...)
}

// ListByStudentIDs mocks base method.
func (m *MockStudentSubmission) ListByStudentIDs(ctx context.Context, studentIDs []string, summaryID int64, fields ...string) (entity.StudentSubmissions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, studentIDs, summaryID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByStudentIDs", varargs...)
	ret0, _ := ret[0].(entity.StudentSubmissions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByStudentIDs indicates an expected call of ListByStudentIDs.
func (mr *MockStudentSubmissionMockRecorder) ListByStudentIDs(ctx, studentIDs, summaryID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, studentIDs, summaryID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByStudentIDs", reflect.TypeOf((*MockStudentSubmission)(nil).ListByStudentIDs), varargs...)
}

// MockStudentShift is a mock of StudentShift interface.
type MockStudentShift struct {
	ctrl     *gomock.Controller
	recorder *MockStudentShiftMockRecorder
}

// MockStudentShiftMockRecorder is the mock recorder for MockStudentShift.
type MockStudentShiftMockRecorder struct {
	mock *MockStudentShift
}

// NewMockStudentShift creates a new mock instance.
func NewMockStudentShift(ctrl *gomock.Controller) *MockStudentShift {
	mock := &MockStudentShift{ctrl: ctrl}
	mock.recorder = &MockStudentShiftMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentShift) EXPECT() *MockStudentShiftMockRecorder {
	return m.recorder
}

// ListByShiftID mocks base method.
func (m *MockStudentShift) ListByShiftID(ctx context.Context, shiftID int64, fields ...string) (entity.StudentShifts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, shiftID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByShiftID", varargs...)
	ret0, _ := ret[0].(entity.StudentShifts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShiftID indicates an expected call of ListByShiftID.
func (mr *MockStudentShiftMockRecorder) ListByShiftID(ctx, shiftID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, shiftID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShiftID", reflect.TypeOf((*MockStudentShift)(nil).ListByShiftID), varargs...)
}

// ListByShiftSummaryID mocks base method.
func (m *MockStudentShift) ListByShiftSummaryID(ctx context.Context, studentIDs []string, summaryID int64, fields ...string) (entity.StudentShifts, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, studentIDs, summaryID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByShiftSummaryID", varargs...)
	ret0, _ := ret[0].(entity.StudentShifts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByShiftSummaryID indicates an expected call of ListByShiftSummaryID.
func (mr *MockStudentShiftMockRecorder) ListByShiftSummaryID(ctx, studentIDs, summaryID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, studentIDs, summaryID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByShiftSummaryID", reflect.TypeOf((*MockStudentShift)(nil).ListByShiftSummaryID), varargs...)
}

// Replace mocks base method.
func (m *MockStudentShift) Replace(ctx context.Context, submission *entity.StudentSubmission, shifts entity.StudentShifts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", ctx, submission, shifts)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replace indicates an expected call of Replace.
func (mr *MockStudentShiftMockRecorder) Replace(ctx, submission, shifts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockStudentShift)(nil).Replace), ctx, submission, shifts)
}

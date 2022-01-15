// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/lesson/service_grpc.pb.go

// Package mock_lesson is a generated GoMock package.
package mock_lesson

import (
	context "context"
	reflect "reflect"

	lesson "github.com/calmato/shs-web/api/proto/lesson"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockLessonServiceClient is a mock of LessonServiceClient interface.
type MockLessonServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockLessonServiceClientMockRecorder
}

// MockLessonServiceClientMockRecorder is the mock recorder for MockLessonServiceClient.
type MockLessonServiceClientMockRecorder struct {
	mock *MockLessonServiceClient
}

// NewMockLessonServiceClient creates a new mock instance.
func NewMockLessonServiceClient(ctrl *gomock.Controller) *MockLessonServiceClient {
	mock := &MockLessonServiceClient{ctrl: ctrl}
	mock.recorder = &MockLessonServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLessonServiceClient) EXPECT() *MockLessonServiceClientMockRecorder {
	return m.recorder
}

// CreateShifts mocks base method.
func (m *MockLessonServiceClient) CreateShifts(ctx context.Context, in *lesson.CreateShiftsRequest, opts ...grpc.CallOption) (*lesson.CreateShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateShifts", varargs...)
	ret0, _ := ret[0].(*lesson.CreateShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShifts indicates an expected call of CreateShifts.
func (mr *MockLessonServiceClientMockRecorder) CreateShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).CreateShifts), varargs...)
}

// DeleteShiftSummary mocks base method.
func (m *MockLessonServiceClient) DeleteShiftSummary(ctx context.Context, in *lesson.DeleteShiftSummaryRequest, opts ...grpc.CallOption) (*lesson.DeleteShiftSummaryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteShiftSummary", varargs...)
	ret0, _ := ret[0].(*lesson.DeleteShiftSummaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteShiftSummary indicates an expected call of DeleteShiftSummary.
func (mr *MockLessonServiceClientMockRecorder) DeleteShiftSummary(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShiftSummary", reflect.TypeOf((*MockLessonServiceClient)(nil).DeleteShiftSummary), varargs...)
}

// GetShiftSummary mocks base method.
func (m *MockLessonServiceClient) GetShiftSummary(ctx context.Context, in *lesson.GetShiftSummaryRequest, opts ...grpc.CallOption) (*lesson.GetShiftSummaryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetShiftSummary", varargs...)
	ret0, _ := ret[0].(*lesson.GetShiftSummaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShiftSummary indicates an expected call of GetShiftSummary.
func (mr *MockLessonServiceClientMockRecorder) GetShiftSummary(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShiftSummary", reflect.TypeOf((*MockLessonServiceClient)(nil).GetShiftSummary), varargs...)
}

// GetStudentShifts mocks base method.
func (m *MockLessonServiceClient) GetStudentShifts(ctx context.Context, in *lesson.GetStudentShiftsRequest, opts ...grpc.CallOption) (*lesson.GetStudentShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudentShifts", varargs...)
	ret0, _ := ret[0].(*lesson.GetStudentShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentShifts indicates an expected call of GetStudentShifts.
func (mr *MockLessonServiceClientMockRecorder) GetStudentShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).GetStudentShifts), varargs...)
}

// GetTeacherShifts mocks base method.
func (m *MockLessonServiceClient) GetTeacherShifts(ctx context.Context, in *lesson.GetTeacherShiftsRequest, opts ...grpc.CallOption) (*lesson.GetTeacherShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTeacherShifts", varargs...)
	ret0, _ := ret[0].(*lesson.GetTeacherShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacherShifts indicates an expected call of GetTeacherShifts.
func (mr *MockLessonServiceClientMockRecorder) GetTeacherShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).GetTeacherShifts), varargs...)
}

// ListShiftSummaries mocks base method.
func (m *MockLessonServiceClient) ListShiftSummaries(ctx context.Context, in *lesson.ListShiftSummariesRequest, opts ...grpc.CallOption) (*lesson.ListShiftSummariesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListShiftSummaries", varargs...)
	ret0, _ := ret[0].(*lesson.ListShiftSummariesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShiftSummaries indicates an expected call of ListShiftSummaries.
func (mr *MockLessonServiceClientMockRecorder) ListShiftSummaries(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShiftSummaries", reflect.TypeOf((*MockLessonServiceClient)(nil).ListShiftSummaries), varargs...)
}

// ListShifts mocks base method.
func (m *MockLessonServiceClient) ListShifts(ctx context.Context, in *lesson.ListShiftsRequest, opts ...grpc.CallOption) (*lesson.ListShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListShifts", varargs...)
	ret0, _ := ret[0].(*lesson.ListShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShifts indicates an expected call of ListShifts.
func (mr *MockLessonServiceClientMockRecorder) ListShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).ListShifts), varargs...)
}

// ListStudentShifts mocks base method.
func (m *MockLessonServiceClient) ListStudentShifts(ctx context.Context, in *lesson.ListStudentShiftsRequest, opts ...grpc.CallOption) (*lesson.ListStudentShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStudentShifts", varargs...)
	ret0, _ := ret[0].(*lesson.ListStudentShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudentShifts indicates an expected call of ListStudentShifts.
func (mr *MockLessonServiceClientMockRecorder) ListStudentShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).ListStudentShifts), varargs...)
}

// ListStudentSubmissionsByShiftSummaryIDs mocks base method.
func (m *MockLessonServiceClient) ListStudentSubmissionsByShiftSummaryIDs(ctx context.Context, in *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest, opts ...grpc.CallOption) (*lesson.ListStudentSubmissionsByShiftSummaryIDsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStudentSubmissionsByShiftSummaryIDs", varargs...)
	ret0, _ := ret[0].(*lesson.ListStudentSubmissionsByShiftSummaryIDsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudentSubmissionsByShiftSummaryIDs indicates an expected call of ListStudentSubmissionsByShiftSummaryIDs.
func (mr *MockLessonServiceClientMockRecorder) ListStudentSubmissionsByShiftSummaryIDs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentSubmissionsByShiftSummaryIDs", reflect.TypeOf((*MockLessonServiceClient)(nil).ListStudentSubmissionsByShiftSummaryIDs), varargs...)
}

// ListStudentSubmissionsByStudentIDs mocks base method.
func (m *MockLessonServiceClient) ListStudentSubmissionsByStudentIDs(ctx context.Context, in *lesson.ListStudentSubmissionsByStudentIDsRequest, opts ...grpc.CallOption) (*lesson.ListStudentSubmissionsByStudentIDsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStudentSubmissionsByStudentIDs", varargs...)
	ret0, _ := ret[0].(*lesson.ListStudentSubmissionsByStudentIDsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudentSubmissionsByStudentIDs indicates an expected call of ListStudentSubmissionsByStudentIDs.
func (mr *MockLessonServiceClientMockRecorder) ListStudentSubmissionsByStudentIDs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentSubmissionsByStudentIDs", reflect.TypeOf((*MockLessonServiceClient)(nil).ListStudentSubmissionsByStudentIDs), varargs...)
}

// ListTeacherShifts mocks base method.
func (m *MockLessonServiceClient) ListTeacherShifts(ctx context.Context, in *lesson.ListTeacherShiftsRequest, opts ...grpc.CallOption) (*lesson.ListTeacherShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTeacherShifts", varargs...)
	ret0, _ := ret[0].(*lesson.ListTeacherShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeacherShifts indicates an expected call of ListTeacherShifts.
func (mr *MockLessonServiceClientMockRecorder) ListTeacherShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).ListTeacherShifts), varargs...)
}

// ListTeacherSubmissionsByShiftSummaryIDs mocks base method.
func (m *MockLessonServiceClient) ListTeacherSubmissionsByShiftSummaryIDs(ctx context.Context, in *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest, opts ...grpc.CallOption) (*lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTeacherSubmissionsByShiftSummaryIDs", varargs...)
	ret0, _ := ret[0].(*lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeacherSubmissionsByShiftSummaryIDs indicates an expected call of ListTeacherSubmissionsByShiftSummaryIDs.
func (mr *MockLessonServiceClientMockRecorder) ListTeacherSubmissionsByShiftSummaryIDs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherSubmissionsByShiftSummaryIDs", reflect.TypeOf((*MockLessonServiceClient)(nil).ListTeacherSubmissionsByShiftSummaryIDs), varargs...)
}

// UpdateShiftSummarySchedule mocks base method.
func (m *MockLessonServiceClient) UpdateShiftSummarySchedule(ctx context.Context, in *lesson.UpdateShiftSummaryScheduleRequest, opts ...grpc.CallOption) (*lesson.UpdateShiftSummaryShceduleResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateShiftSummarySchedule", varargs...)
	ret0, _ := ret[0].(*lesson.UpdateShiftSummaryShceduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateShiftSummarySchedule indicates an expected call of UpdateShiftSummarySchedule.
func (mr *MockLessonServiceClientMockRecorder) UpdateShiftSummarySchedule(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShiftSummarySchedule", reflect.TypeOf((*MockLessonServiceClient)(nil).UpdateShiftSummarySchedule), varargs...)
}

// UpsertStudentShifts mocks base method.
func (m *MockLessonServiceClient) UpsertStudentShifts(ctx context.Context, in *lesson.UpsertStudentShiftsRequest, opts ...grpc.CallOption) (*lesson.UpsertStudentShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertStudentShifts", varargs...)
	ret0, _ := ret[0].(*lesson.UpsertStudentShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertStudentShifts indicates an expected call of UpsertStudentShifts.
func (mr *MockLessonServiceClientMockRecorder) UpsertStudentShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertStudentShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).UpsertStudentShifts), varargs...)
}

// UpsertTeacherShifts mocks base method.
func (m *MockLessonServiceClient) UpsertTeacherShifts(ctx context.Context, in *lesson.UpsertTeacherShiftsRequest, opts ...grpc.CallOption) (*lesson.UpsertTeacherShiftsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertTeacherShifts", varargs...)
	ret0, _ := ret[0].(*lesson.UpsertTeacherShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertTeacherShifts indicates an expected call of UpsertTeacherShifts.
func (mr *MockLessonServiceClientMockRecorder) UpsertTeacherShifts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTeacherShifts", reflect.TypeOf((*MockLessonServiceClient)(nil).UpsertTeacherShifts), varargs...)
}

// MockLessonServiceServer is a mock of LessonServiceServer interface.
type MockLessonServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockLessonServiceServerMockRecorder
}

// MockLessonServiceServerMockRecorder is the mock recorder for MockLessonServiceServer.
type MockLessonServiceServerMockRecorder struct {
	mock *MockLessonServiceServer
}

// NewMockLessonServiceServer creates a new mock instance.
func NewMockLessonServiceServer(ctrl *gomock.Controller) *MockLessonServiceServer {
	mock := &MockLessonServiceServer{ctrl: ctrl}
	mock.recorder = &MockLessonServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLessonServiceServer) EXPECT() *MockLessonServiceServerMockRecorder {
	return m.recorder
}

// CreateShifts mocks base method.
func (m *MockLessonServiceServer) CreateShifts(arg0 context.Context, arg1 *lesson.CreateShiftsRequest) (*lesson.CreateShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.CreateShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShifts indicates an expected call of CreateShifts.
func (mr *MockLessonServiceServerMockRecorder) CreateShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).CreateShifts), arg0, arg1)
}

// DeleteShiftSummary mocks base method.
func (m *MockLessonServiceServer) DeleteShiftSummary(arg0 context.Context, arg1 *lesson.DeleteShiftSummaryRequest) (*lesson.DeleteShiftSummaryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteShiftSummary", arg0, arg1)
	ret0, _ := ret[0].(*lesson.DeleteShiftSummaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteShiftSummary indicates an expected call of DeleteShiftSummary.
func (mr *MockLessonServiceServerMockRecorder) DeleteShiftSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShiftSummary", reflect.TypeOf((*MockLessonServiceServer)(nil).DeleteShiftSummary), arg0, arg1)
}

// GetShiftSummary mocks base method.
func (m *MockLessonServiceServer) GetShiftSummary(arg0 context.Context, arg1 *lesson.GetShiftSummaryRequest) (*lesson.GetShiftSummaryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShiftSummary", arg0, arg1)
	ret0, _ := ret[0].(*lesson.GetShiftSummaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShiftSummary indicates an expected call of GetShiftSummary.
func (mr *MockLessonServiceServerMockRecorder) GetShiftSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShiftSummary", reflect.TypeOf((*MockLessonServiceServer)(nil).GetShiftSummary), arg0, arg1)
}

// GetStudentShifts mocks base method.
func (m *MockLessonServiceServer) GetStudentShifts(arg0 context.Context, arg1 *lesson.GetStudentShiftsRequest) (*lesson.GetStudentShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.GetStudentShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentShifts indicates an expected call of GetStudentShifts.
func (mr *MockLessonServiceServerMockRecorder) GetStudentShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).GetStudentShifts), arg0, arg1)
}

// GetTeacherShifts mocks base method.
func (m *MockLessonServiceServer) GetTeacherShifts(arg0 context.Context, arg1 *lesson.GetTeacherShiftsRequest) (*lesson.GetTeacherShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeacherShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.GetTeacherShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacherShifts indicates an expected call of GetTeacherShifts.
func (mr *MockLessonServiceServerMockRecorder) GetTeacherShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).GetTeacherShifts), arg0, arg1)
}

// ListShiftSummaries mocks base method.
func (m *MockLessonServiceServer) ListShiftSummaries(arg0 context.Context, arg1 *lesson.ListShiftSummariesRequest) (*lesson.ListShiftSummariesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShiftSummaries", arg0, arg1)
	ret0, _ := ret[0].(*lesson.ListShiftSummariesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShiftSummaries indicates an expected call of ListShiftSummaries.
func (mr *MockLessonServiceServerMockRecorder) ListShiftSummaries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShiftSummaries", reflect.TypeOf((*MockLessonServiceServer)(nil).ListShiftSummaries), arg0, arg1)
}

// ListShifts mocks base method.
func (m *MockLessonServiceServer) ListShifts(arg0 context.Context, arg1 *lesson.ListShiftsRequest) (*lesson.ListShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.ListShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShifts indicates an expected call of ListShifts.
func (mr *MockLessonServiceServerMockRecorder) ListShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).ListShifts), arg0, arg1)
}

// ListStudentShifts mocks base method.
func (m *MockLessonServiceServer) ListStudentShifts(arg0 context.Context, arg1 *lesson.ListStudentShiftsRequest) (*lesson.ListStudentShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudentShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.ListStudentShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudentShifts indicates an expected call of ListStudentShifts.
func (mr *MockLessonServiceServerMockRecorder) ListStudentShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).ListStudentShifts), arg0, arg1)
}

// ListStudentSubmissionsByShiftSummaryIDs mocks base method.
func (m *MockLessonServiceServer) ListStudentSubmissionsByShiftSummaryIDs(arg0 context.Context, arg1 *lesson.ListStudentSubmissionsByShiftSummaryIDsRequest) (*lesson.ListStudentSubmissionsByShiftSummaryIDsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudentSubmissionsByShiftSummaryIDs", arg0, arg1)
	ret0, _ := ret[0].(*lesson.ListStudentSubmissionsByShiftSummaryIDsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudentSubmissionsByShiftSummaryIDs indicates an expected call of ListStudentSubmissionsByShiftSummaryIDs.
func (mr *MockLessonServiceServerMockRecorder) ListStudentSubmissionsByShiftSummaryIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentSubmissionsByShiftSummaryIDs", reflect.TypeOf((*MockLessonServiceServer)(nil).ListStudentSubmissionsByShiftSummaryIDs), arg0, arg1)
}

// ListStudentSubmissionsByStudentIDs mocks base method.
func (m *MockLessonServiceServer) ListStudentSubmissionsByStudentIDs(arg0 context.Context, arg1 *lesson.ListStudentSubmissionsByStudentIDsRequest) (*lesson.ListStudentSubmissionsByStudentIDsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudentSubmissionsByStudentIDs", arg0, arg1)
	ret0, _ := ret[0].(*lesson.ListStudentSubmissionsByStudentIDsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudentSubmissionsByStudentIDs indicates an expected call of ListStudentSubmissionsByStudentIDs.
func (mr *MockLessonServiceServerMockRecorder) ListStudentSubmissionsByStudentIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudentSubmissionsByStudentIDs", reflect.TypeOf((*MockLessonServiceServer)(nil).ListStudentSubmissionsByStudentIDs), arg0, arg1)
}

// ListTeacherShifts mocks base method.
func (m *MockLessonServiceServer) ListTeacherShifts(arg0 context.Context, arg1 *lesson.ListTeacherShiftsRequest) (*lesson.ListTeacherShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeacherShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.ListTeacherShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeacherShifts indicates an expected call of ListTeacherShifts.
func (mr *MockLessonServiceServerMockRecorder) ListTeacherShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).ListTeacherShifts), arg0, arg1)
}

// ListTeacherSubmissionsByShiftSummaryIDs mocks base method.
func (m *MockLessonServiceServer) ListTeacherSubmissionsByShiftSummaryIDs(arg0 context.Context, arg1 *lesson.ListTeacherSubmissionsByShiftSummaryIDsRequest) (*lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeacherSubmissionsByShiftSummaryIDs", arg0, arg1)
	ret0, _ := ret[0].(*lesson.ListTeacherSubmissionsByShiftSummaryIDsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeacherSubmissionsByShiftSummaryIDs indicates an expected call of ListTeacherSubmissionsByShiftSummaryIDs.
func (mr *MockLessonServiceServerMockRecorder) ListTeacherSubmissionsByShiftSummaryIDs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeacherSubmissionsByShiftSummaryIDs", reflect.TypeOf((*MockLessonServiceServer)(nil).ListTeacherSubmissionsByShiftSummaryIDs), arg0, arg1)
}

// UpdateShiftSummarySchedule mocks base method.
func (m *MockLessonServiceServer) UpdateShiftSummarySchedule(arg0 context.Context, arg1 *lesson.UpdateShiftSummaryScheduleRequest) (*lesson.UpdateShiftSummaryShceduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShiftSummarySchedule", arg0, arg1)
	ret0, _ := ret[0].(*lesson.UpdateShiftSummaryShceduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateShiftSummarySchedule indicates an expected call of UpdateShiftSummarySchedule.
func (mr *MockLessonServiceServerMockRecorder) UpdateShiftSummarySchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShiftSummarySchedule", reflect.TypeOf((*MockLessonServiceServer)(nil).UpdateShiftSummarySchedule), arg0, arg1)
}

// UpsertStudentShifts mocks base method.
func (m *MockLessonServiceServer) UpsertStudentShifts(arg0 context.Context, arg1 *lesson.UpsertStudentShiftsRequest) (*lesson.UpsertStudentShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertStudentShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.UpsertStudentShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertStudentShifts indicates an expected call of UpsertStudentShifts.
func (mr *MockLessonServiceServerMockRecorder) UpsertStudentShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertStudentShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).UpsertStudentShifts), arg0, arg1)
}

// UpsertTeacherShifts mocks base method.
func (m *MockLessonServiceServer) UpsertTeacherShifts(arg0 context.Context, arg1 *lesson.UpsertTeacherShiftsRequest) (*lesson.UpsertTeacherShiftsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertTeacherShifts", arg0, arg1)
	ret0, _ := ret[0].(*lesson.UpsertTeacherShiftsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertTeacherShifts indicates an expected call of UpsertTeacherShifts.
func (mr *MockLessonServiceServerMockRecorder) UpsertTeacherShifts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTeacherShifts", reflect.TypeOf((*MockLessonServiceServer)(nil).UpsertTeacherShifts), arg0, arg1)
}

// mustEmbedUnimplementedLessonServiceServer mocks base method.
func (m *MockLessonServiceServer) mustEmbedUnimplementedLessonServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedLessonServiceServer")
}

// mustEmbedUnimplementedLessonServiceServer indicates an expected call of mustEmbedUnimplementedLessonServiceServer.
func (mr *MockLessonServiceServerMockRecorder) mustEmbedUnimplementedLessonServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedLessonServiceServer", reflect.TypeOf((*MockLessonServiceServer)(nil).mustEmbedUnimplementedLessonServiceServer))
}

// MockUnsafeLessonServiceServer is a mock of UnsafeLessonServiceServer interface.
type MockUnsafeLessonServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeLessonServiceServerMockRecorder
}

// MockUnsafeLessonServiceServerMockRecorder is the mock recorder for MockUnsafeLessonServiceServer.
type MockUnsafeLessonServiceServerMockRecorder struct {
	mock *MockUnsafeLessonServiceServer
}

// NewMockUnsafeLessonServiceServer creates a new mock instance.
func NewMockUnsafeLessonServiceServer(ctrl *gomock.Controller) *MockUnsafeLessonServiceServer {
	mock := &MockUnsafeLessonServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeLessonServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeLessonServiceServer) EXPECT() *MockUnsafeLessonServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedLessonServiceServer mocks base method.
func (m *MockUnsafeLessonServiceServer) mustEmbedUnimplementedLessonServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedLessonServiceServer")
}

// mustEmbedUnimplementedLessonServiceServer indicates an expected call of mustEmbedUnimplementedLessonServiceServer.
func (mr *MockUnsafeLessonServiceServerMockRecorder) mustEmbedUnimplementedLessonServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedLessonServiceServer", reflect.TypeOf((*MockUnsafeLessonServiceServer)(nil).mustEmbedUnimplementedLessonServiceServer))
}

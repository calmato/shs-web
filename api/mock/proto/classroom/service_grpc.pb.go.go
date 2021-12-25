// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/classroom/service_grpc.pb.go

// Package mock_classroom is a generated GoMock package.
package mock_classroom

import (
	context "context"
	reflect "reflect"

	classroom "github.com/calmato/shs-web/api/proto/classroom"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockClassroomServiceClient is a mock of ClassroomServiceClient interface.
type MockClassroomServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockClassroomServiceClientMockRecorder
}

// MockClassroomServiceClientMockRecorder is the mock recorder for MockClassroomServiceClient.
type MockClassroomServiceClientMockRecorder struct {
	mock *MockClassroomServiceClient
}

// NewMockClassroomServiceClient creates a new mock instance.
func NewMockClassroomServiceClient(ctrl *gomock.Controller) *MockClassroomServiceClient {
	mock := &MockClassroomServiceClient{ctrl: ctrl}
	mock.recorder = &MockClassroomServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClassroomServiceClient) EXPECT() *MockClassroomServiceClientMockRecorder {
	return m.recorder
}

// GetSchedule mocks base method.
func (m *MockClassroomServiceClient) GetSchedule(ctx context.Context, in *classroom.GetScheduleRequest, opts ...grpc.CallOption) (*classroom.GetScheduleResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchedule", varargs...)
	ret0, _ := ret[0].(*classroom.GetScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule.
func (mr *MockClassroomServiceClientMockRecorder) GetSchedule(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockClassroomServiceClient)(nil).GetSchedule), varargs...)
}

// GetSubject mocks base method.
func (m *MockClassroomServiceClient) GetSubject(ctx context.Context, in *classroom.GetSubjectRequest, opts ...grpc.CallOption) (*classroom.GetSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubject", varargs...)
	ret0, _ := ret[0].(*classroom.GetSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubject indicates an expected call of GetSubject.
func (mr *MockClassroomServiceClientMockRecorder) GetSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).GetSubject), varargs...)
}

// GetTeacherSubject mocks base method.
func (m *MockClassroomServiceClient) GetTeacherSubject(ctx context.Context, in *classroom.GetTeacherSubjectRequest, opts ...grpc.CallOption) (*classroom.GetTeacherSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTeacherSubject", varargs...)
	ret0, _ := ret[0].(*classroom.GetTeacherSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacherSubject indicates an expected call of GetTeacherSubject.
func (mr *MockClassroomServiceClientMockRecorder) GetTeacherSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).GetTeacherSubject), varargs...)
}

// ListSchedules mocks base method.
func (m *MockClassroomServiceClient) ListSchedules(ctx context.Context, in *classroom.ListSchedulesRequest, opts ...grpc.CallOption) (*classroom.ListSchedulesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchedules", varargs...)
	ret0, _ := ret[0].(*classroom.ListSchedulesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedules indicates an expected call of ListSchedules.
func (mr *MockClassroomServiceClientMockRecorder) ListSchedules(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedules", reflect.TypeOf((*MockClassroomServiceClient)(nil).ListSchedules), varargs...)
}

// ListSubjects mocks base method.
func (m *MockClassroomServiceClient) ListSubjects(ctx context.Context, in *classroom.ListSubjectsRequest, opts ...grpc.CallOption) (*classroom.ListSubjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSubjects", varargs...)
	ret0, _ := ret[0].(*classroom.ListSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjects indicates an expected call of ListSubjects.
func (mr *MockClassroomServiceClientMockRecorder) ListSubjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjects", reflect.TypeOf((*MockClassroomServiceClient)(nil).ListSubjects), varargs...)
}

// MultiGetSubjects mocks base method.
func (m *MockClassroomServiceClient) MultiGetSubjects(ctx context.Context, in *classroom.MultiGetSubjectsRequest, opts ...grpc.CallOption) (*classroom.MultiGetSubjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGetSubjects", varargs...)
	ret0, _ := ret[0].(*classroom.MultiGetSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetSubjects indicates an expected call of MultiGetSubjects.
func (mr *MockClassroomServiceClientMockRecorder) MultiGetSubjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetSubjects", reflect.TypeOf((*MockClassroomServiceClient)(nil).MultiGetSubjects), varargs...)
}

// MultiGetTeacherSubjects mocks base method.
func (m *MockClassroomServiceClient) MultiGetTeacherSubjects(ctx context.Context, in *classroom.MultiGetTeacherSubjectsRequest, opts ...grpc.CallOption) (*classroom.MultiGetTeacherSubjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGetTeacherSubjects", varargs...)
	ret0, _ := ret[0].(*classroom.MultiGetTeacherSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetTeacherSubjects indicates an expected call of MultiGetTeacherSubjects.
func (mr *MockClassroomServiceClientMockRecorder) MultiGetTeacherSubjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetTeacherSubjects", reflect.TypeOf((*MockClassroomServiceClient)(nil).MultiGetTeacherSubjects), varargs...)
}

// UpdateSchedules mocks base method.
func (m *MockClassroomServiceClient) UpdateSchedules(ctx context.Context, in *classroom.UpdateSchedulesRequest, opts ...grpc.CallOption) (*classroom.UpdateSchedulesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSchedules", varargs...)
	ret0, _ := ret[0].(*classroom.UpdateSchedulesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSchedules indicates an expected call of UpdateSchedules.
func (mr *MockClassroomServiceClientMockRecorder) UpdateSchedules(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedules", reflect.TypeOf((*MockClassroomServiceClient)(nil).UpdateSchedules), varargs...)
}

// UpdateTeacherSubject mocks base method.
func (m *MockClassroomServiceClient) UpdateTeacherSubject(ctx context.Context, in *classroom.UpdateTeacherSubjectRequest, opts ...grpc.CallOption) (*classroom.UpdateTeacherSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTeacherSubject", varargs...)
	ret0, _ := ret[0].(*classroom.UpdateTeacherSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherSubject indicates an expected call of UpdateTeacherSubject.
func (mr *MockClassroomServiceClientMockRecorder) UpdateTeacherSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).UpdateTeacherSubject), varargs...)
}

// MockClassroomServiceServer is a mock of ClassroomServiceServer interface.
type MockClassroomServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockClassroomServiceServerMockRecorder
}

// MockClassroomServiceServerMockRecorder is the mock recorder for MockClassroomServiceServer.
type MockClassroomServiceServerMockRecorder struct {
	mock *MockClassroomServiceServer
}

// NewMockClassroomServiceServer creates a new mock instance.
func NewMockClassroomServiceServer(ctrl *gomock.Controller) *MockClassroomServiceServer {
	mock := &MockClassroomServiceServer{ctrl: ctrl}
	mock.recorder = &MockClassroomServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClassroomServiceServer) EXPECT() *MockClassroomServiceServerMockRecorder {
	return m.recorder
}

// GetSchedule mocks base method.
func (m *MockClassroomServiceServer) GetSchedule(arg0 context.Context, arg1 *classroom.GetScheduleRequest) (*classroom.GetScheduleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchedule", arg0, arg1)
	ret0, _ := ret[0].(*classroom.GetScheduleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchedule indicates an expected call of GetSchedule.
func (mr *MockClassroomServiceServerMockRecorder) GetSchedule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedule", reflect.TypeOf((*MockClassroomServiceServer)(nil).GetSchedule), arg0, arg1)
}

// GetSubject mocks base method.
func (m *MockClassroomServiceServer) GetSubject(arg0 context.Context, arg1 *classroom.GetSubjectRequest) (*classroom.GetSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.GetSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubject indicates an expected call of GetSubject.
func (mr *MockClassroomServiceServerMockRecorder) GetSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).GetSubject), arg0, arg1)
}

// GetTeacherSubject mocks base method.
func (m *MockClassroomServiceServer) GetTeacherSubject(arg0 context.Context, arg1 *classroom.GetTeacherSubjectRequest) (*classroom.GetTeacherSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeacherSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.GetTeacherSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacherSubject indicates an expected call of GetTeacherSubject.
func (mr *MockClassroomServiceServerMockRecorder) GetTeacherSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacherSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).GetTeacherSubject), arg0, arg1)
}

// ListSchedules mocks base method.
func (m *MockClassroomServiceServer) ListSchedules(arg0 context.Context, arg1 *classroom.ListSchedulesRequest) (*classroom.ListSchedulesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSchedules", arg0, arg1)
	ret0, _ := ret[0].(*classroom.ListSchedulesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSchedules indicates an expected call of ListSchedules.
func (mr *MockClassroomServiceServerMockRecorder) ListSchedules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchedules", reflect.TypeOf((*MockClassroomServiceServer)(nil).ListSchedules), arg0, arg1)
}

// ListSubjects mocks base method.
func (m *MockClassroomServiceServer) ListSubjects(arg0 context.Context, arg1 *classroom.ListSubjectsRequest) (*classroom.ListSubjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSubjects", arg0, arg1)
	ret0, _ := ret[0].(*classroom.ListSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSubjects indicates an expected call of ListSubjects.
func (mr *MockClassroomServiceServerMockRecorder) ListSubjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubjects", reflect.TypeOf((*MockClassroomServiceServer)(nil).ListSubjects), arg0, arg1)
}

// MultiGetSubjects mocks base method.
func (m *MockClassroomServiceServer) MultiGetSubjects(arg0 context.Context, arg1 *classroom.MultiGetSubjectsRequest) (*classroom.MultiGetSubjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetSubjects", arg0, arg1)
	ret0, _ := ret[0].(*classroom.MultiGetSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetSubjects indicates an expected call of MultiGetSubjects.
func (mr *MockClassroomServiceServerMockRecorder) MultiGetSubjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetSubjects", reflect.TypeOf((*MockClassroomServiceServer)(nil).MultiGetSubjects), arg0, arg1)
}

// MultiGetTeacherSubjects mocks base method.
func (m *MockClassroomServiceServer) MultiGetTeacherSubjects(arg0 context.Context, arg1 *classroom.MultiGetTeacherSubjectsRequest) (*classroom.MultiGetTeacherSubjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetTeacherSubjects", arg0, arg1)
	ret0, _ := ret[0].(*classroom.MultiGetTeacherSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetTeacherSubjects indicates an expected call of MultiGetTeacherSubjects.
func (mr *MockClassroomServiceServerMockRecorder) MultiGetTeacherSubjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetTeacherSubjects", reflect.TypeOf((*MockClassroomServiceServer)(nil).MultiGetTeacherSubjects), arg0, arg1)
}

// UpdateSchedules mocks base method.
func (m *MockClassroomServiceServer) UpdateSchedules(arg0 context.Context, arg1 *classroom.UpdateSchedulesRequest) (*classroom.UpdateSchedulesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSchedules", arg0, arg1)
	ret0, _ := ret[0].(*classroom.UpdateSchedulesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSchedules indicates an expected call of UpdateSchedules.
func (mr *MockClassroomServiceServerMockRecorder) UpdateSchedules(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSchedules", reflect.TypeOf((*MockClassroomServiceServer)(nil).UpdateSchedules), arg0, arg1)
}

// UpdateTeacherSubject mocks base method.
func (m *MockClassroomServiceServer) UpdateTeacherSubject(arg0 context.Context, arg1 *classroom.UpdateTeacherSubjectRequest) (*classroom.UpdateTeacherSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacherSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.UpdateTeacherSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherSubject indicates an expected call of UpdateTeacherSubject.
func (mr *MockClassroomServiceServerMockRecorder) UpdateTeacherSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).UpdateTeacherSubject), arg0, arg1)
}

// mustEmbedUnimplementedClassroomServiceServer mocks base method.
func (m *MockClassroomServiceServer) mustEmbedUnimplementedClassroomServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedClassroomServiceServer")
}

// mustEmbedUnimplementedClassroomServiceServer indicates an expected call of mustEmbedUnimplementedClassroomServiceServer.
func (mr *MockClassroomServiceServerMockRecorder) mustEmbedUnimplementedClassroomServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedClassroomServiceServer", reflect.TypeOf((*MockClassroomServiceServer)(nil).mustEmbedUnimplementedClassroomServiceServer))
}

// MockUnsafeClassroomServiceServer is a mock of UnsafeClassroomServiceServer interface.
type MockUnsafeClassroomServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeClassroomServiceServerMockRecorder
}

// MockUnsafeClassroomServiceServerMockRecorder is the mock recorder for MockUnsafeClassroomServiceServer.
type MockUnsafeClassroomServiceServerMockRecorder struct {
	mock *MockUnsafeClassroomServiceServer
}

// NewMockUnsafeClassroomServiceServer creates a new mock instance.
func NewMockUnsafeClassroomServiceServer(ctrl *gomock.Controller) *MockUnsafeClassroomServiceServer {
	mock := &MockUnsafeClassroomServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeClassroomServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeClassroomServiceServer) EXPECT() *MockUnsafeClassroomServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedClassroomServiceServer mocks base method.
func (m *MockUnsafeClassroomServiceServer) mustEmbedUnimplementedClassroomServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedClassroomServiceServer")
}

// mustEmbedUnimplementedClassroomServiceServer indicates an expected call of mustEmbedUnimplementedClassroomServiceServer.
func (mr *MockUnsafeClassroomServiceServerMockRecorder) mustEmbedUnimplementedClassroomServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedClassroomServiceServer", reflect.TypeOf((*MockUnsafeClassroomServiceServer)(nil).mustEmbedUnimplementedClassroomServiceServer))
}

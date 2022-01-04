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

// CreateSubject mocks base method.
func (m *MockClassroomServiceClient) CreateSubject(ctx context.Context, in *classroom.CreateSubjectRequest, opts ...grpc.CallOption) (*classroom.CreateSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSubject", varargs...)
	ret0, _ := ret[0].(*classroom.CreateSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubject indicates an expected call of CreateSubject.
func (mr *MockClassroomServiceClientMockRecorder) CreateSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).CreateSubject), varargs...)
}

// DeleteSubject mocks base method.
func (m *MockClassroomServiceClient) DeleteSubject(ctx context.Context, in *classroom.DeleteSubjectRequest, opts ...grpc.CallOption) (*classroom.DeleteSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSubject", varargs...)
	ret0, _ := ret[0].(*classroom.DeleteSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSubject indicates an expected call of DeleteSubject.
func (mr *MockClassroomServiceClientMockRecorder) DeleteSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).DeleteSubject), varargs...)
}

// GetRoomsTotal mocks base method.
func (m *MockClassroomServiceClient) GetRoomsTotal(ctx context.Context, in *classroom.GetRoomsTotalRequest, opts ...grpc.CallOption) (*classroom.GetRoomsTotalResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRoomsTotal", varargs...)
	ret0, _ := ret[0].(*classroom.GetRoomsTotalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomsTotal indicates an expected call of GetRoomsTotal.
func (mr *MockClassroomServiceClientMockRecorder) GetRoomsTotal(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomsTotal", reflect.TypeOf((*MockClassroomServiceClient)(nil).GetRoomsTotal), varargs...)
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

// GetStudentSubject mocks base method.
func (m *MockClassroomServiceClient) GetStudentSubject(ctx context.Context, in *classroom.GetStudentSubjectRequest, opts ...grpc.CallOption) (*classroom.GetStudentSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudentSubject", varargs...)
	ret0, _ := ret[0].(*classroom.GetStudentSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentSubject indicates an expected call of GetStudentSubject.
func (mr *MockClassroomServiceClientMockRecorder) GetStudentSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).GetStudentSubject), varargs...)
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

// MultiGetStudentSubjects mocks base method.
func (m *MockClassroomServiceClient) MultiGetStudentSubjects(ctx context.Context, in *classroom.MultiGetStudentSubjectsRequest, opts ...grpc.CallOption) (*classroom.MultiGetStudentSubjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGetStudentSubjects", varargs...)
	ret0, _ := ret[0].(*classroom.MultiGetStudentSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetStudentSubjects indicates an expected call of MultiGetStudentSubjects.
func (mr *MockClassroomServiceClientMockRecorder) MultiGetStudentSubjects(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetStudentSubjects", reflect.TypeOf((*MockClassroomServiceClient)(nil).MultiGetStudentSubjects), varargs...)
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

// UpdateRoomsTotal mocks base method.
func (m *MockClassroomServiceClient) UpdateRoomsTotal(ctx context.Context, in *classroom.UpdateRoomsTotalRequest, opts ...grpc.CallOption) (*classroom.UpdateRoomsTotalResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateRoomsTotal", varargs...)
	ret0, _ := ret[0].(*classroom.UpdateRoomsTotalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoomsTotal indicates an expected call of UpdateRoomsTotal.
func (mr *MockClassroomServiceClientMockRecorder) UpdateRoomsTotal(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoomsTotal", reflect.TypeOf((*MockClassroomServiceClient)(nil).UpdateRoomsTotal), varargs...)
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

// UpdateSubject mocks base method.
func (m *MockClassroomServiceClient) UpdateSubject(ctx context.Context, in *classroom.UpdateSubjectRequest, opts ...grpc.CallOption) (*classroom.UpdateSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSubject", varargs...)
	ret0, _ := ret[0].(*classroom.UpdateSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubject indicates an expected call of UpdateSubject.
func (mr *MockClassroomServiceClientMockRecorder) UpdateSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).UpdateSubject), varargs...)
}

// UpsertStudentSubject mocks base method.
func (m *MockClassroomServiceClient) UpsertStudentSubject(ctx context.Context, in *classroom.UpsertStudentSubjectRequest, opts ...grpc.CallOption) (*classroom.UpsertStudentSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertStudentSubject", varargs...)
	ret0, _ := ret[0].(*classroom.UpsertStudentSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertStudentSubject indicates an expected call of UpsertStudentSubject.
func (mr *MockClassroomServiceClientMockRecorder) UpsertStudentSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertStudentSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).UpsertStudentSubject), varargs...)
}

// UpsertTeacherSubject mocks base method.
func (m *MockClassroomServiceClient) UpsertTeacherSubject(ctx context.Context, in *classroom.UpsertTeacherSubjectRequest, opts ...grpc.CallOption) (*classroom.UpsertTeacherSubjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertTeacherSubject", varargs...)
	ret0, _ := ret[0].(*classroom.UpsertTeacherSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertTeacherSubject indicates an expected call of UpsertTeacherSubject.
func (mr *MockClassroomServiceClientMockRecorder) UpsertTeacherSubject(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTeacherSubject", reflect.TypeOf((*MockClassroomServiceClient)(nil).UpsertTeacherSubject), varargs...)
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

// CreateSubject mocks base method.
func (m *MockClassroomServiceServer) CreateSubject(arg0 context.Context, arg1 *classroom.CreateSubjectRequest) (*classroom.CreateSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.CreateSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubject indicates an expected call of CreateSubject.
func (mr *MockClassroomServiceServerMockRecorder) CreateSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).CreateSubject), arg0, arg1)
}

// DeleteSubject mocks base method.
func (m *MockClassroomServiceServer) DeleteSubject(arg0 context.Context, arg1 *classroom.DeleteSubjectRequest) (*classroom.DeleteSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.DeleteSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSubject indicates an expected call of DeleteSubject.
func (mr *MockClassroomServiceServerMockRecorder) DeleteSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).DeleteSubject), arg0, arg1)
}

// GetRoomsTotal mocks base method.
func (m *MockClassroomServiceServer) GetRoomsTotal(arg0 context.Context, arg1 *classroom.GetRoomsTotalRequest) (*classroom.GetRoomsTotalResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomsTotal", arg0, arg1)
	ret0, _ := ret[0].(*classroom.GetRoomsTotalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomsTotal indicates an expected call of GetRoomsTotal.
func (mr *MockClassroomServiceServerMockRecorder) GetRoomsTotal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomsTotal", reflect.TypeOf((*MockClassroomServiceServer)(nil).GetRoomsTotal), arg0, arg1)
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

// GetStudentSubject mocks base method.
func (m *MockClassroomServiceServer) GetStudentSubject(arg0 context.Context, arg1 *classroom.GetStudentSubjectRequest) (*classroom.GetStudentSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudentSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.GetStudentSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudentSubject indicates an expected call of GetStudentSubject.
func (mr *MockClassroomServiceServerMockRecorder) GetStudentSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudentSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).GetStudentSubject), arg0, arg1)
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

// MultiGetStudentSubjects mocks base method.
func (m *MockClassroomServiceServer) MultiGetStudentSubjects(arg0 context.Context, arg1 *classroom.MultiGetStudentSubjectsRequest) (*classroom.MultiGetStudentSubjectsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetStudentSubjects", arg0, arg1)
	ret0, _ := ret[0].(*classroom.MultiGetStudentSubjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetStudentSubjects indicates an expected call of MultiGetStudentSubjects.
func (mr *MockClassroomServiceServerMockRecorder) MultiGetStudentSubjects(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetStudentSubjects", reflect.TypeOf((*MockClassroomServiceServer)(nil).MultiGetStudentSubjects), arg0, arg1)
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

// UpdateRoomsTotal mocks base method.
func (m *MockClassroomServiceServer) UpdateRoomsTotal(arg0 context.Context, arg1 *classroom.UpdateRoomsTotalRequest) (*classroom.UpdateRoomsTotalResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRoomsTotal", arg0, arg1)
	ret0, _ := ret[0].(*classroom.UpdateRoomsTotalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRoomsTotal indicates an expected call of UpdateRoomsTotal.
func (mr *MockClassroomServiceServerMockRecorder) UpdateRoomsTotal(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRoomsTotal", reflect.TypeOf((*MockClassroomServiceServer)(nil).UpdateRoomsTotal), arg0, arg1)
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

// UpdateSubject mocks base method.
func (m *MockClassroomServiceServer) UpdateSubject(arg0 context.Context, arg1 *classroom.UpdateSubjectRequest) (*classroom.UpdateSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.UpdateSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubject indicates an expected call of UpdateSubject.
func (mr *MockClassroomServiceServerMockRecorder) UpdateSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).UpdateSubject), arg0, arg1)
}

// UpsertStudentSubject mocks base method.
func (m *MockClassroomServiceServer) UpsertStudentSubject(arg0 context.Context, arg1 *classroom.UpsertStudentSubjectRequest) (*classroom.UpsertStudentSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertStudentSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.UpsertStudentSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertStudentSubject indicates an expected call of UpsertStudentSubject.
func (mr *MockClassroomServiceServerMockRecorder) UpsertStudentSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertStudentSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).UpsertStudentSubject), arg0, arg1)
}

// UpsertTeacherSubject mocks base method.
func (m *MockClassroomServiceServer) UpsertTeacherSubject(arg0 context.Context, arg1 *classroom.UpsertTeacherSubjectRequest) (*classroom.UpsertTeacherSubjectResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertTeacherSubject", arg0, arg1)
	ret0, _ := ret[0].(*classroom.UpsertTeacherSubjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertTeacherSubject indicates an expected call of UpsertTeacherSubject.
func (mr *MockClassroomServiceServerMockRecorder) UpsertTeacherSubject(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTeacherSubject", reflect.TypeOf((*MockClassroomServiceServer)(nil).UpsertTeacherSubject), arg0, arg1)
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

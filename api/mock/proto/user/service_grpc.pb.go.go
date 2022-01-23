// Code generated by MockGen. DO NOT EDIT.
// Source: ./proto/user/service_grpc.pb.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	user "github.com/calmato/shs-web/api/proto/user"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// CreateStudent mocks base method.
func (m *MockUserServiceClient) CreateStudent(ctx context.Context, in *user.CreateStudentRequest, opts ...grpc.CallOption) (*user.CreateStudentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateStudent", varargs...)
	ret0, _ := ret[0].(*user.CreateStudentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockUserServiceClientMockRecorder) CreateStudent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockUserServiceClient)(nil).CreateStudent), varargs...)
}

// CreateTeacher mocks base method.
func (m *MockUserServiceClient) CreateTeacher(ctx context.Context, in *user.CreateTeacherRequest, opts ...grpc.CallOption) (*user.CreateTeacherResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTeacher", varargs...)
	ret0, _ := ret[0].(*user.CreateTeacherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeacher indicates an expected call of CreateTeacher.
func (mr *MockUserServiceClientMockRecorder) CreateTeacher(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeacher", reflect.TypeOf((*MockUserServiceClient)(nil).CreateTeacher), varargs...)
}

// DeleteStudent mocks base method.
func (m *MockUserServiceClient) DeleteStudent(ctx context.Context, in *user.DeleteStudentRequest, opts ...grpc.CallOption) (*user.DeleteStudentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteStudent", varargs...)
	ret0, _ := ret[0].(*user.DeleteStudentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockUserServiceClientMockRecorder) DeleteStudent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockUserServiceClient)(nil).DeleteStudent), varargs...)
}

// DeleteTeacher mocks base method.
func (m *MockUserServiceClient) DeleteTeacher(ctx context.Context, in *user.DeleteTeacherRequest, opts ...grpc.CallOption) (*user.DeleteTeacherResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTeacher", varargs...)
	ret0, _ := ret[0].(*user.DeleteTeacherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTeacher indicates an expected call of DeleteTeacher.
func (mr *MockUserServiceClientMockRecorder) DeleteTeacher(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeacher", reflect.TypeOf((*MockUserServiceClient)(nil).DeleteTeacher), varargs...)
}

// GetStudent mocks base method.
func (m *MockUserServiceClient) GetStudent(ctx context.Context, in *user.GetStudentRequest, opts ...grpc.CallOption) (*user.GetStudentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudent", varargs...)
	ret0, _ := ret[0].(*user.GetStudentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudent indicates an expected call of GetStudent.
func (mr *MockUserServiceClientMockRecorder) GetStudent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudent", reflect.TypeOf((*MockUserServiceClient)(nil).GetStudent), varargs...)
}

// GetTeacher mocks base method.
func (m *MockUserServiceClient) GetTeacher(ctx context.Context, in *user.GetTeacherRequest, opts ...grpc.CallOption) (*user.GetTeacherResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTeacher", varargs...)
	ret0, _ := ret[0].(*user.GetTeacherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacher indicates an expected call of GetTeacher.
func (mr *MockUserServiceClientMockRecorder) GetTeacher(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacher", reflect.TypeOf((*MockUserServiceClient)(nil).GetTeacher), varargs...)
}

// ListStudents mocks base method.
func (m *MockUserServiceClient) ListStudents(ctx context.Context, in *user.ListStudentsRequest, opts ...grpc.CallOption) (*user.ListStudentsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStudents", varargs...)
	ret0, _ := ret[0].(*user.ListStudentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudents indicates an expected call of ListStudents.
func (mr *MockUserServiceClientMockRecorder) ListStudents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudents", reflect.TypeOf((*MockUserServiceClient)(nil).ListStudents), varargs...)
}

// ListTeachers mocks base method.
func (m *MockUserServiceClient) ListTeachers(ctx context.Context, in *user.ListTeachersRequest, opts ...grpc.CallOption) (*user.ListTeachersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTeachers", varargs...)
	ret0, _ := ret[0].(*user.ListTeachersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeachers indicates an expected call of ListTeachers.
func (mr *MockUserServiceClientMockRecorder) ListTeachers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeachers", reflect.TypeOf((*MockUserServiceClient)(nil).ListTeachers), varargs...)
}

// MultiGetTeachers mocks base method.
func (m *MockUserServiceClient) MultiGetTeachers(ctx context.Context, in *user.MultiGetTeachersRequest, opts ...grpc.CallOption) (*user.MultiGetTeachersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGetTeachers", varargs...)
	ret0, _ := ret[0].(*user.MultiGetTeachersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetTeachers indicates an expected call of MultiGetTeachers.
func (mr *MockUserServiceClientMockRecorder) MultiGetTeachers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetTeachers", reflect.TypeOf((*MockUserServiceClient)(nil).MultiGetTeachers), varargs...)
}

// UpdateTeacherMail mocks base method.
func (m *MockUserServiceClient) UpdateTeacherMail(ctx context.Context, in *user.UpdateTeacherMailRequest, opts ...grpc.CallOption) (*user.UpdateTeacherMailResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTeacherMail", varargs...)
	ret0, _ := ret[0].(*user.UpdateTeacherMailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherMail indicates an expected call of UpdateTeacherMail.
func (mr *MockUserServiceClientMockRecorder) UpdateTeacherMail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherMail", reflect.TypeOf((*MockUserServiceClient)(nil).UpdateTeacherMail), varargs...)
}

// UpdateTeacherPassword mocks base method.
func (m *MockUserServiceClient) UpdateTeacherPassword(ctx context.Context, in *user.UpdateTeacherPasswordRequest, opts ...grpc.CallOption) (*user.UpdateTeacherPasswordResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTeacherPassword", varargs...)
	ret0, _ := ret[0].(*user.UpdateTeacherPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherPassword indicates an expected call of UpdateTeacherPassword.
func (mr *MockUserServiceClientMockRecorder) UpdateTeacherPassword(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherPassword", reflect.TypeOf((*MockUserServiceClient)(nil).UpdateTeacherPassword), varargs...)
}

// UpdateTeacherRole mocks base method.
func (m *MockUserServiceClient) UpdateTeacherRole(ctx context.Context, in *user.UpdateTeacherRoleRequest, opts ...grpc.CallOption) (*user.UpdateTeacherRoleResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTeacherRole", varargs...)
	ret0, _ := ret[0].(*user.UpdateTeacherRoleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherRole indicates an expected call of UpdateTeacherRole.
func (mr *MockUserServiceClientMockRecorder) UpdateTeacherRole(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherRole", reflect.TypeOf((*MockUserServiceClient)(nil).UpdateTeacherRole), varargs...)
}

// MockUserServiceServer is a mock of UserServiceServer interface.
type MockUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceServerMockRecorder
}

// MockUserServiceServerMockRecorder is the mock recorder for MockUserServiceServer.
type MockUserServiceServerMockRecorder struct {
	mock *MockUserServiceServer
}

// NewMockUserServiceServer creates a new mock instance.
func NewMockUserServiceServer(ctrl *gomock.Controller) *MockUserServiceServer {
	mock := &MockUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceServer) EXPECT() *MockUserServiceServerMockRecorder {
	return m.recorder
}

// CreateStudent mocks base method.
func (m *MockUserServiceServer) CreateStudent(arg0 context.Context, arg1 *user.CreateStudentRequest) (*user.CreateStudentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", arg0, arg1)
	ret0, _ := ret[0].(*user.CreateStudentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockUserServiceServerMockRecorder) CreateStudent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockUserServiceServer)(nil).CreateStudent), arg0, arg1)
}

// CreateTeacher mocks base method.
func (m *MockUserServiceServer) CreateTeacher(arg0 context.Context, arg1 *user.CreateTeacherRequest) (*user.CreateTeacherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTeacher", arg0, arg1)
	ret0, _ := ret[0].(*user.CreateTeacherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTeacher indicates an expected call of CreateTeacher.
func (mr *MockUserServiceServerMockRecorder) CreateTeacher(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTeacher", reflect.TypeOf((*MockUserServiceServer)(nil).CreateTeacher), arg0, arg1)
}

// DeleteStudent mocks base method.
func (m *MockUserServiceServer) DeleteStudent(arg0 context.Context, arg1 *user.DeleteStudentRequest) (*user.DeleteStudentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", arg0, arg1)
	ret0, _ := ret[0].(*user.DeleteStudentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockUserServiceServerMockRecorder) DeleteStudent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockUserServiceServer)(nil).DeleteStudent), arg0, arg1)
}

// DeleteTeacher mocks base method.
func (m *MockUserServiceServer) DeleteTeacher(arg0 context.Context, arg1 *user.DeleteTeacherRequest) (*user.DeleteTeacherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTeacher", arg0, arg1)
	ret0, _ := ret[0].(*user.DeleteTeacherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTeacher indicates an expected call of DeleteTeacher.
func (mr *MockUserServiceServerMockRecorder) DeleteTeacher(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTeacher", reflect.TypeOf((*MockUserServiceServer)(nil).DeleteTeacher), arg0, arg1)
}

// GetStudent mocks base method.
func (m *MockUserServiceServer) GetStudent(arg0 context.Context, arg1 *user.GetStudentRequest) (*user.GetStudentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudent", arg0, arg1)
	ret0, _ := ret[0].(*user.GetStudentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudent indicates an expected call of GetStudent.
func (mr *MockUserServiceServerMockRecorder) GetStudent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudent", reflect.TypeOf((*MockUserServiceServer)(nil).GetStudent), arg0, arg1)
}

// GetTeacher mocks base method.
func (m *MockUserServiceServer) GetTeacher(arg0 context.Context, arg1 *user.GetTeacherRequest) (*user.GetTeacherResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeacher", arg0, arg1)
	ret0, _ := ret[0].(*user.GetTeacherResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeacher indicates an expected call of GetTeacher.
func (mr *MockUserServiceServerMockRecorder) GetTeacher(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeacher", reflect.TypeOf((*MockUserServiceServer)(nil).GetTeacher), arg0, arg1)
}

// ListStudents mocks base method.
func (m *MockUserServiceServer) ListStudents(arg0 context.Context, arg1 *user.ListStudentsRequest) (*user.ListStudentsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudents", arg0, arg1)
	ret0, _ := ret[0].(*user.ListStudentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudents indicates an expected call of ListStudents.
func (mr *MockUserServiceServerMockRecorder) ListStudents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudents", reflect.TypeOf((*MockUserServiceServer)(nil).ListStudents), arg0, arg1)
}

// ListTeachers mocks base method.
func (m *MockUserServiceServer) ListTeachers(arg0 context.Context, arg1 *user.ListTeachersRequest) (*user.ListTeachersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeachers", arg0, arg1)
	ret0, _ := ret[0].(*user.ListTeachersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTeachers indicates an expected call of ListTeachers.
func (mr *MockUserServiceServerMockRecorder) ListTeachers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeachers", reflect.TypeOf((*MockUserServiceServer)(nil).ListTeachers), arg0, arg1)
}

// MultiGetTeachers mocks base method.
func (m *MockUserServiceServer) MultiGetTeachers(arg0 context.Context, arg1 *user.MultiGetTeachersRequest) (*user.MultiGetTeachersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetTeachers", arg0, arg1)
	ret0, _ := ret[0].(*user.MultiGetTeachersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetTeachers indicates an expected call of MultiGetTeachers.
func (mr *MockUserServiceServerMockRecorder) MultiGetTeachers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetTeachers", reflect.TypeOf((*MockUserServiceServer)(nil).MultiGetTeachers), arg0, arg1)
}

// UpdateTeacherMail mocks base method.
func (m *MockUserServiceServer) UpdateTeacherMail(arg0 context.Context, arg1 *user.UpdateTeacherMailRequest) (*user.UpdateTeacherMailResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacherMail", arg0, arg1)
	ret0, _ := ret[0].(*user.UpdateTeacherMailResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherMail indicates an expected call of UpdateTeacherMail.
func (mr *MockUserServiceServerMockRecorder) UpdateTeacherMail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherMail", reflect.TypeOf((*MockUserServiceServer)(nil).UpdateTeacherMail), arg0, arg1)
}

// UpdateTeacherPassword mocks base method.
func (m *MockUserServiceServer) UpdateTeacherPassword(arg0 context.Context, arg1 *user.UpdateTeacherPasswordRequest) (*user.UpdateTeacherPasswordResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacherPassword", arg0, arg1)
	ret0, _ := ret[0].(*user.UpdateTeacherPasswordResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherPassword indicates an expected call of UpdateTeacherPassword.
func (mr *MockUserServiceServerMockRecorder) UpdateTeacherPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherPassword", reflect.TypeOf((*MockUserServiceServer)(nil).UpdateTeacherPassword), arg0, arg1)
}

// UpdateTeacherRole mocks base method.
func (m *MockUserServiceServer) UpdateTeacherRole(arg0 context.Context, arg1 *user.UpdateTeacherRoleRequest) (*user.UpdateTeacherRoleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTeacherRole", arg0, arg1)
	ret0, _ := ret[0].(*user.UpdateTeacherRoleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTeacherRole indicates an expected call of UpdateTeacherRole.
func (mr *MockUserServiceServerMockRecorder) UpdateTeacherRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTeacherRole", reflect.TypeOf((*MockUserServiceServer)(nil).UpdateTeacherRole), arg0, arg1)
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

// MockUnsafeUserServiceServer is a mock of UnsafeUserServiceServer interface.
type MockUnsafeUserServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeUserServiceServerMockRecorder
}

// MockUnsafeUserServiceServerMockRecorder is the mock recorder for MockUnsafeUserServiceServer.
type MockUnsafeUserServiceServerMockRecorder struct {
	mock *MockUnsafeUserServiceServer
}

// NewMockUnsafeUserServiceServer creates a new mock instance.
func NewMockUnsafeUserServiceServer(ctrl *gomock.Controller) *MockUnsafeUserServiceServer {
	mock := &MockUnsafeUserServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeUserServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeUserServiceServer) EXPECT() *MockUnsafeUserServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedUserServiceServer mocks base method.
func (m *MockUnsafeUserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedUserServiceServer")
}

// mustEmbedUnimplementedUserServiceServer indicates an expected call of mustEmbedUnimplementedUserServiceServer.
func (mr *MockUnsafeUserServiceServerMockRecorder) mustEmbedUnimplementedUserServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedUserServiceServer", reflect.TypeOf((*MockUnsafeUserServiceServer)(nil).mustEmbedUnimplementedUserServiceServer))
}

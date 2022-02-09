// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_pubsub is a generated GoMock package.
package mock_pubsub

import (
	context "context"
	reflect "reflect"

	pubsub "github.com/calmato/shs-web/api/pkg/pubsub"
	gomock "github.com/golang/mock/gomock"
)

// MockPublisher is a mock of Publisher interface.
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher.
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance.
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockPublisher) Publish(ctx context.Context, msg *pubsub.Message) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, msg)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockPublisherMockRecorder) Publish(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockPublisher)(nil).Publish), ctx, msg)
}

// Stop mocks base method.
func (m *MockPublisher) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockPublisherMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockPublisher)(nil).Stop))
}

// MockPuller is a mock of Puller interface.
type MockPuller struct {
	ctrl     *gomock.Controller
	recorder *MockPullerMockRecorder
}

// MockPullerMockRecorder is the mock recorder for MockPuller.
type MockPullerMockRecorder struct {
	mock *MockPuller
}

// NewMockPuller creates a new mock instance.
func NewMockPuller(ctrl *gomock.Controller) *MockPuller {
	mock := &MockPuller{ctrl: ctrl}
	mock.recorder = &MockPullerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPuller) EXPECT() *MockPullerMockRecorder {
	return m.recorder
}

// Pull mocks base method.
func (m *MockPuller) Pull(ctx context.Context, msgCh chan<- *pubsub.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", ctx, msgCh)
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull.
func (mr *MockPullerMockRecorder) Pull(ctx, msgCh interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockPuller)(nil).Pull), ctx, msgCh)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moiot/gravity/pkg/protocol (interfaces: MessageRouter,MessageTracer,MessageWrapper)

// Package mock_protocol is a generated GoMock package.
package mock_protocol

import (
	reflect "reflect"

	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"

	protocol "github.com/moiot/gravity/pkg/protocol"
)

// MockMessageRouter is a mock of MessageRouter interface
type MockMessageRouter struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRouterMockRecorder
}

// MockMessageRouterMockRecorder is the mock recorder for MockMessageRouter
type MockMessageRouterMockRecorder struct {
	mock *MockMessageRouter
}

// NewMockMessageRouter creates a new mock instance
func NewMockMessageRouter(ctrl *gomock.Controller) *MockMessageRouter {
	mock := &MockMessageRouter{ctrl: ctrl}
	mock.recorder = &MockMessageRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageRouter) EXPECT() *MockMessageRouterMockRecorder {
	return m.recorder
}

// GetPartitions mocks base method
func (m *MockMessageRouter) GetPartitions() []int32 {
	ret := m.ctrl.Call(m, "GetPartitions")
	ret0, _ := ret[0].([]int32)
	return ret0
}

// GetPartitions indicates an expected call of GetPartitions
func (mr *MockMessageRouterMockRecorder) GetPartitions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitions", reflect.TypeOf((*MockMessageRouter)(nil).GetPartitions))
}

// GetTopic mocks base method
func (m *MockMessageRouter) GetTopic() string {
	ret := m.ctrl.Call(m, "GetTopic")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTopic indicates an expected call of GetTopic
func (mr *MockMessageRouterMockRecorder) GetTopic() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopic", reflect.TypeOf((*MockMessageRouter)(nil).GetTopic))
}

// MockMessageTracer is a mock of MessageTracer interface
type MockMessageTracer struct {
	ctrl     *gomock.Controller
	recorder *MockMessageTracerMockRecorder
}

// MockMessageTracerMockRecorder is the mock recorder for MockMessageTracer
type MockMessageTracerMockRecorder struct {
	mock *MockMessageTracer
}

// NewMockMessageTracer creates a new mock instance
func NewMockMessageTracer(ctrl *gomock.Controller) *MockMessageTracer {
	mock := &MockMessageTracer{ctrl: ctrl}
	mock.recorder = &MockMessageTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageTracer) EXPECT() *MockMessageTracerMockRecorder {
	return m.recorder
}

// AddMetrics mocks base method
func (m *MockMessageTracer) AddMetrics() {
	m.ctrl.Call(m, "AddMetrics")
}

// AddMetrics indicates an expected call of AddMetrics
func (mr *MockMessageTracerMockRecorder) AddMetrics() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMetrics", reflect.TypeOf((*MockMessageTracer)(nil).AddMetrics))
}

// AddTimestamp mocks base method
func (m *MockMessageTracer) AddTimestamp(arg0 protocol.MsgTimestamp) {
	m.ctrl.Call(m, "AddTimestamp", arg0)
}

// AddTimestamp indicates an expected call of AddTimestamp
func (mr *MockMessageTracerMockRecorder) AddTimestamp(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTimestamp", reflect.TypeOf((*MockMessageTracer)(nil).AddTimestamp), arg0)
}

// MetricsString mocks base method
func (m *MockMessageTracer) MetricsString() string {
	ret := m.ctrl.Call(m, "MetricsString")
	ret0, _ := ret[0].(string)
	return ret0
}

// MetricsString indicates an expected call of MetricsString
func (mr *MockMessageTracerMockRecorder) MetricsString() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetricsString", reflect.TypeOf((*MockMessageTracer)(nil).MetricsString))
}

// MockMessageWrapper is a mock of MessageWrapper interface
type MockMessageWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockMessageWrapperMockRecorder
}

// MockMessageWrapperMockRecorder is the mock recorder for MockMessageWrapper
type MockMessageWrapperMockRecorder struct {
	mock *MockMessageWrapper
}

// NewMockMessageWrapper creates a new mock instance
func NewMockMessageWrapper(ctrl *gomock.Controller) *MockMessageWrapper {
	mock := &MockMessageWrapper{ctrl: ctrl}
	mock.recorder = &MockMessageWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageWrapper) EXPECT() *MockMessageWrapperMockRecorder {
	return m.recorder
}

// GetDatabase mocks base method
func (m *MockMessageWrapper) GetDatabase() string {
	ret := m.ctrl.Call(m, "GetDatabase")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabase indicates an expected call of GetDatabase
func (mr *MockMessageWrapperMockRecorder) GetDatabase() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabase", reflect.TypeOf((*MockMessageWrapper)(nil).GetDatabase))
}

// GetMqMsgType mocks base method
func (m *MockMessageWrapper) GetMqMsgType() protocol.JobMsgType {
	ret := m.ctrl.Call(m, "GetMqMsgType")
	ret0, _ := ret[0].(protocol.JobMsgType)
	return ret0
}

// GetMqMsgType indicates an expected call of GetMqMsgType
func (mr *MockMessageWrapperMockRecorder) GetMqMsgType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMqMsgType", reflect.TypeOf((*MockMessageWrapper)(nil).GetMqMsgType))
}

// GetPartitionKeyFromPayload mocks base method
func (m *MockMessageWrapper) GetPartitionKeyFromPayload(arg0 []string) string {
	ret := m.ctrl.Call(m, "GetPartitionKeyFromPayload", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPartitionKeyFromPayload indicates an expected call of GetPartitionKeyFromPayload
func (mr *MockMessageWrapperMockRecorder) GetPartitionKeyFromPayload(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionKeyFromPayload", reflect.TypeOf((*MockMessageWrapper)(nil).GetPartitionKeyFromPayload), arg0)
}

// GetPayloadKafkaEncoder mocks base method
func (m *MockMessageWrapper) GetPayloadKafkaEncoder() (sarama.Encoder, error) {
	ret := m.ctrl.Call(m, "GetPayloadKafkaEncoder")
	ret0, _ := ret[0].(sarama.Encoder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayloadKafkaEncoder indicates an expected call of GetPayloadKafkaEncoder
func (mr *MockMessageWrapperMockRecorder) GetPayloadKafkaEncoder() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayloadKafkaEncoder", reflect.TypeOf((*MockMessageWrapper)(nil).GetPayloadKafkaEncoder))
}

// GetTable mocks base method
func (m *MockMessageWrapper) GetTable() string {
	ret := m.ctrl.Call(m, "GetTable")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTable indicates an expected call of GetTable
func (mr *MockMessageWrapperMockRecorder) GetTable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTable", reflect.TypeOf((*MockMessageWrapper)(nil).GetTable))
}

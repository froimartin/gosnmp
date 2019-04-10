// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package gosnmp is a generated GoMock package.
package gosnmp

import (
	x "."
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Connect mocks base method
func (m *MockHandler) Connect() error {
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockHandlerMockRecorder) Connect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockHandler)(nil).Connect))
}

// ConnectIPv4 mocks base method
func (m *MockHandler) ConnectIPv4() error {
	ret := m.ctrl.Call(m, "ConnectIPv4")
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectIPv4 indicates an expected call of ConnectIPv4
func (mr *MockHandlerMockRecorder) ConnectIPv4() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectIPv4", reflect.TypeOf((*MockHandler)(nil).ConnectIPv4))
}

// ConnectIPv6 mocks base method
func (m *MockHandler) ConnectIPv6() error {
	ret := m.ctrl.Call(m, "ConnectIPv6")
	ret0, _ := ret[0].(error)
	return ret0
}

// ConnectIPv6 indicates an expected call of ConnectIPv6
func (mr *MockHandlerMockRecorder) ConnectIPv6() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectIPv6", reflect.TypeOf((*MockHandler)(nil).ConnectIPv6))
}

// Get mocks base method
func (m *MockHandler) Get(oids []string) (*x.SnmpPacket, error) {
	ret := m.ctrl.Call(m, "Get", oids)
	ret0, _ := ret[0].(*x.SnmpPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockHandlerMockRecorder) Get(oids interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockHandler)(nil).Get), oids)
}

// GetBulk mocks base method
func (m *MockHandler) GetBulk(oids []string, nonRepeaters, maxRepetitions uint8) (*x.SnmpPacket, error) {
	ret := m.ctrl.Call(m, "GetBulk", oids, nonRepeaters, maxRepetitions)
	ret0, _ := ret[0].(*x.SnmpPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBulk indicates an expected call of GetBulk
func (mr *MockHandlerMockRecorder) GetBulk(oids, nonRepeaters, maxRepetitions interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBulk", reflect.TypeOf((*MockHandler)(nil).GetBulk), oids, nonRepeaters, maxRepetitions)
}

// GetNext mocks base method
func (m *MockHandler) GetNext(oids []string) (*x.SnmpPacket, error) {
	ret := m.ctrl.Call(m, "GetNext", oids)
	ret0, _ := ret[0].(*x.SnmpPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNext indicates an expected call of GetNext
func (mr *MockHandlerMockRecorder) GetNext(oids interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNext", reflect.TypeOf((*MockHandler)(nil).GetNext), oids)
}

// Walk mocks base method
func (m *MockHandler) Walk(rootOid string, walkFn x.WalkFunc) error {
	ret := m.ctrl.Call(m, "Walk", rootOid, walkFn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk
func (mr *MockHandlerMockRecorder) Walk(rootOid, walkFn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockHandler)(nil).Walk), rootOid, walkFn)
}

// WalkAll mocks base method
func (m *MockHandler) WalkAll(rootOid string) ([]x.SnmpPDU, error) {
	ret := m.ctrl.Call(m, "WalkAll", rootOid)
	ret0, _ := ret[0].([]x.SnmpPDU)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalkAll indicates an expected call of WalkAll
func (mr *MockHandlerMockRecorder) WalkAll(rootOid interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkAll", reflect.TypeOf((*MockHandler)(nil).WalkAll), rootOid)
}

// BulkWalk mocks base method
func (m *MockHandler) BulkWalk(rootOid string, walkFn x.WalkFunc) error {
	ret := m.ctrl.Call(m, "BulkWalk", rootOid, walkFn)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkWalk indicates an expected call of BulkWalk
func (mr *MockHandlerMockRecorder) BulkWalk(rootOid, walkFn interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkWalk", reflect.TypeOf((*MockHandler)(nil).BulkWalk), rootOid, walkFn)
}

// BulkWalkAll mocks base method
func (m *MockHandler) BulkWalkAll(rootOid string) ([]x.SnmpPDU, error) {
	ret := m.ctrl.Call(m, "BulkWalkAll", rootOid)
	ret0, _ := ret[0].([]x.SnmpPDU)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkWalkAll indicates an expected call of BulkWalkAll
func (mr *MockHandlerMockRecorder) BulkWalkAll(rootOid interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkWalkAll", reflect.TypeOf((*MockHandler)(nil).BulkWalkAll), rootOid)
}

// SendTrap mocks base method
func (m *MockHandler) SendTrap(trap x.SnmpTrap) (*x.SnmpPacket, error) {
	ret := m.ctrl.Call(m, "SendTrap", trap)
	ret0, _ := ret[0].(*x.SnmpPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTrap indicates an expected call of SendTrap
func (mr *MockHandlerMockRecorder) SendTrap(trap interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTrap", reflect.TypeOf((*MockHandler)(nil).SendTrap), trap)
}

// UnmarshalTrap mocks base method
func (m *MockHandler) UnmarshalTrap(trap []byte) *x.SnmpPacket {
	ret := m.ctrl.Call(m, "UnmarshalTrap", trap)
	ret0, _ := ret[0].(*x.SnmpPacket)
	return ret0
}

// UnmarshalTrap indicates an expected call of UnmarshalTrap
func (mr *MockHandlerMockRecorder) UnmarshalTrap(trap interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalTrap", reflect.TypeOf((*MockHandler)(nil).UnmarshalTrap), trap)
}

// Set mocks base method
func (m *MockHandler) Set(pdus []x.SnmpPDU) (*x.SnmpPacket, error) {
	ret := m.ctrl.Call(m, "Set", pdus)
	ret0, _ := ret[0].(*x.SnmpPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set
func (mr *MockHandlerMockRecorder) Set(pdus interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockHandler)(nil).Set), pdus)
}

// Check mocks base method
func (m *MockHandler) Check(err error) {
	m.ctrl.Call(m, "Check", err)
}

// Check indicates an expected call of Check
func (mr *MockHandlerMockRecorder) Check(err interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockHandler)(nil).Check), err)
}

// Close mocks base method
func (m *MockHandler) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockHandlerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHandler)(nil).Close))
}

// Target mocks base method
func (m *MockHandler) Target() string {
	ret := m.ctrl.Call(m, "Target")
	ret0, _ := ret[0].(string)
	return ret0
}

// Target indicates an expected call of Target
func (mr *MockHandlerMockRecorder) Target() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Target", reflect.TypeOf((*MockHandler)(nil).Target))
}

// SetTarget mocks base method
func (m *MockHandler) SetTarget(target string) {
	m.ctrl.Call(m, "SetTarget", target)
}

// SetTarget indicates an expected call of SetTarget
func (mr *MockHandlerMockRecorder) SetTarget(target interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTarget", reflect.TypeOf((*MockHandler)(nil).SetTarget), target)
}

// Port mocks base method
func (m *MockHandler) Port() uint16 {
	ret := m.ctrl.Call(m, "Port")
	ret0, _ := ret[0].(uint16)
	return ret0
}

// Port indicates an expected call of Port
func (mr *MockHandlerMockRecorder) Port() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Port", reflect.TypeOf((*MockHandler)(nil).Port))
}

// SetPort mocks base method
func (m *MockHandler) SetPort(port uint16) {
	m.ctrl.Call(m, "SetPort", port)
}

// SetPort indicates an expected call of SetPort
func (mr *MockHandlerMockRecorder) SetPort(port interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPort", reflect.TypeOf((*MockHandler)(nil).SetPort), port)
}

// Community mocks base method
func (m *MockHandler) Community() string {
	ret := m.ctrl.Call(m, "Community")
	ret0, _ := ret[0].(string)
	return ret0
}

// Community indicates an expected call of Community
func (mr *MockHandlerMockRecorder) Community() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Community", reflect.TypeOf((*MockHandler)(nil).Community))
}

// SetCommunity mocks base method
func (m *MockHandler) SetCommunity(community string) {
	m.ctrl.Call(m, "SetCommunity", community)
}

// SetCommunity indicates an expected call of SetCommunity
func (mr *MockHandlerMockRecorder) SetCommunity(community interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCommunity", reflect.TypeOf((*MockHandler)(nil).SetCommunity), community)
}

// Version mocks base method
func (m *MockHandler) Version() x.SnmpVersion {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(x.SnmpVersion)
	return ret0
}

// Version indicates an expected call of Version
func (mr *MockHandlerMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockHandler)(nil).Version))
}

// SetVersion mocks base method
func (m *MockHandler) SetVersion(version x.SnmpVersion) {
	m.ctrl.Call(m, "SetVersion", version)
}

// SetVersion indicates an expected call of SetVersion
func (mr *MockHandlerMockRecorder) SetVersion(version interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVersion", reflect.TypeOf((*MockHandler)(nil).SetVersion), version)
}

// Timeout mocks base method
func (m *MockHandler) Timeout() time.Duration {
	ret := m.ctrl.Call(m, "Timeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// Timeout indicates an expected call of Timeout
func (mr *MockHandlerMockRecorder) Timeout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timeout", reflect.TypeOf((*MockHandler)(nil).Timeout))
}

// SetTimeout mocks base method
func (m *MockHandler) SetTimeout(timeout time.Duration) {
	m.ctrl.Call(m, "SetTimeout", timeout)
}

// SetTimeout indicates an expected call of SetTimeout
func (mr *MockHandlerMockRecorder) SetTimeout(timeout interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockHandler)(nil).SetTimeout), timeout)
}

// Retries mocks base method
func (m *MockHandler) Retries() int {
	ret := m.ctrl.Call(m, "Retries")
	ret0, _ := ret[0].(int)
	return ret0
}

// Retries indicates an expected call of Retries
func (mr *MockHandlerMockRecorder) Retries() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Retries", reflect.TypeOf((*MockHandler)(nil).Retries))
}

// SetRetries mocks base method
func (m *MockHandler) SetRetries(retries int) {
	m.ctrl.Call(m, "SetRetries", retries)
}

// SetRetries indicates an expected call of SetRetries
func (mr *MockHandlerMockRecorder) SetRetries(retries interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRetries", reflect.TypeOf((*MockHandler)(nil).SetRetries), retries)
}

// GetExponentialTimeout mocks base method
func (m *MockHandler) GetExponentialTimeout() bool {
	ret := m.ctrl.Call(m, "GetExponentialTimeout")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetExponentialTimeout indicates an expected call of GetExponentialTimeout
func (mr *MockHandlerMockRecorder) GetExponentialTimeout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExponentialTimeout", reflect.TypeOf((*MockHandler)(nil).GetExponentialTimeout))
}

// SetExponentialTimeout mocks base method
func (m *MockHandler) SetExponentialTimeout(value bool) {
	m.ctrl.Call(m, "SetExponentialTimeout", value)
}

// SetExponentialTimeout indicates an expected call of SetExponentialTimeout
func (mr *MockHandlerMockRecorder) SetExponentialTimeout(value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExponentialTimeout", reflect.TypeOf((*MockHandler)(nil).SetExponentialTimeout), value)
}

// Logger mocks base method
func (m *MockHandler) Logger() x.Logger {
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(x.Logger)
	return ret0
}

// Logger indicates an expected call of Logger
func (mr *MockHandlerMockRecorder) Logger() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockHandler)(nil).Logger))
}

// SetLogger mocks base method
func (m *MockHandler) SetLogger(logger x.Logger) {
	m.ctrl.Call(m, "SetLogger", logger)
}

// SetLogger indicates an expected call of SetLogger
func (mr *MockHandlerMockRecorder) SetLogger(logger interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockHandler)(nil).SetLogger), logger)
}

// MaxOids mocks base method
func (m *MockHandler) MaxOids() int {
	ret := m.ctrl.Call(m, "MaxOids")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxOids indicates an expected call of MaxOids
func (mr *MockHandlerMockRecorder) MaxOids() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxOids", reflect.TypeOf((*MockHandler)(nil).MaxOids))
}

// SetMaxOids mocks base method
func (m *MockHandler) SetMaxOids(maxOids int) {
	m.ctrl.Call(m, "SetMaxOids", maxOids)
}

// SetMaxOids indicates an expected call of SetMaxOids
func (mr *MockHandlerMockRecorder) SetMaxOids(maxOids interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxOids", reflect.TypeOf((*MockHandler)(nil).SetMaxOids), maxOids)
}

// MaxRepetitions mocks base method
func (m *MockHandler) MaxRepetitions() uint8 {
	ret := m.ctrl.Call(m, "MaxRepetitions")
	ret0, _ := ret[0].(uint8)
	return ret0
}

// MaxRepetitions indicates an expected call of MaxRepetitions
func (mr *MockHandlerMockRecorder) MaxRepetitions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxRepetitions", reflect.TypeOf((*MockHandler)(nil).MaxRepetitions))
}

// SetMaxRepetitions mocks base method
func (m *MockHandler) SetMaxRepetitions(maxRepetitions uint8) {
	m.ctrl.Call(m, "SetMaxRepetitions", maxRepetitions)
}

// SetMaxRepetitions indicates an expected call of SetMaxRepetitions
func (mr *MockHandlerMockRecorder) SetMaxRepetitions(maxRepetitions interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxRepetitions", reflect.TypeOf((*MockHandler)(nil).SetMaxRepetitions), maxRepetitions)
}

// NonRepeaters mocks base method
func (m *MockHandler) NonRepeaters() int {
	ret := m.ctrl.Call(m, "NonRepeaters")
	ret0, _ := ret[0].(int)
	return ret0
}

// NonRepeaters indicates an expected call of NonRepeaters
func (mr *MockHandlerMockRecorder) NonRepeaters() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NonRepeaters", reflect.TypeOf((*MockHandler)(nil).NonRepeaters))
}

// SetNonRepeaters mocks base method
func (m *MockHandler) SetNonRepeaters(nonRepeaters int) {
	m.ctrl.Call(m, "SetNonRepeaters", nonRepeaters)
}

// SetNonRepeaters indicates an expected call of SetNonRepeaters
func (mr *MockHandlerMockRecorder) SetNonRepeaters(nonRepeaters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNonRepeaters", reflect.TypeOf((*MockHandler)(nil).SetNonRepeaters), nonRepeaters)
}

// MsgFlags mocks base method
func (m *MockHandler) MsgFlags() x.SnmpV3MsgFlags {
	ret := m.ctrl.Call(m, "MsgFlags")
	ret0, _ := ret[0].(x.SnmpV3MsgFlags)
	return ret0
}

// MsgFlags indicates an expected call of MsgFlags
func (mr *MockHandlerMockRecorder) MsgFlags() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MsgFlags", reflect.TypeOf((*MockHandler)(nil).MsgFlags))
}

// SetMsgFlags mocks base method
func (m *MockHandler) SetMsgFlags(msgFlags x.SnmpV3MsgFlags) {
	m.ctrl.Call(m, "SetMsgFlags", msgFlags)
}

// SetMsgFlags indicates an expected call of SetMsgFlags
func (mr *MockHandlerMockRecorder) SetMsgFlags(msgFlags interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMsgFlags", reflect.TypeOf((*MockHandler)(nil).SetMsgFlags), msgFlags)
}

// SecurityModel mocks base method
func (m *MockHandler) SecurityModel() x.SnmpV3SecurityModel {
	ret := m.ctrl.Call(m, "SecurityModel")
	ret0, _ := ret[0].(x.SnmpV3SecurityModel)
	return ret0
}

// SecurityModel indicates an expected call of SecurityModel
func (mr *MockHandlerMockRecorder) SecurityModel() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityModel", reflect.TypeOf((*MockHandler)(nil).SecurityModel))
}

// SetSecurityModel mocks base method
func (m *MockHandler) SetSecurityModel(securityModel x.SnmpV3SecurityModel) {
	m.ctrl.Call(m, "SetSecurityModel", securityModel)
}

// SetSecurityModel indicates an expected call of SetSecurityModel
func (mr *MockHandlerMockRecorder) SetSecurityModel(securityModel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecurityModel", reflect.TypeOf((*MockHandler)(nil).SetSecurityModel), securityModel)
}

// SecurityParameters mocks base method
func (m *MockHandler) SecurityParameters() x.SnmpV3SecurityParameters {
	ret := m.ctrl.Call(m, "SecurityParameters")
	ret0, _ := ret[0].(x.SnmpV3SecurityParameters)
	return ret0
}

// SecurityParameters indicates an expected call of SecurityParameters
func (mr *MockHandlerMockRecorder) SecurityParameters() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityParameters", reflect.TypeOf((*MockHandler)(nil).SecurityParameters))
}

// SetSecurityParameters mocks base method
func (m *MockHandler) SetSecurityParameters(securityParameters x.SnmpV3SecurityParameters) {
	m.ctrl.Call(m, "SetSecurityParameters", securityParameters)
}

// SetSecurityParameters indicates an expected call of SetSecurityParameters
func (mr *MockHandlerMockRecorder) SetSecurityParameters(securityParameters interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecurityParameters", reflect.TypeOf((*MockHandler)(nil).SetSecurityParameters), securityParameters)
}

// ContextEngineID mocks base method
func (m *MockHandler) ContextEngineID() string {
	ret := m.ctrl.Call(m, "ContextEngineID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContextEngineID indicates an expected call of ContextEngineID
func (mr *MockHandlerMockRecorder) ContextEngineID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContextEngineID", reflect.TypeOf((*MockHandler)(nil).ContextEngineID))
}

// SetContextEngineID mocks base method
func (m *MockHandler) SetContextEngineID(contextEngineID string) {
	m.ctrl.Call(m, "SetContextEngineID", contextEngineID)
}

// SetContextEngineID indicates an expected call of SetContextEngineID
func (mr *MockHandlerMockRecorder) SetContextEngineID(contextEngineID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContextEngineID", reflect.TypeOf((*MockHandler)(nil).SetContextEngineID), contextEngineID)
}

// ContextName mocks base method
func (m *MockHandler) ContextName() string {
	ret := m.ctrl.Call(m, "ContextName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ContextName indicates an expected call of ContextName
func (mr *MockHandlerMockRecorder) ContextName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContextName", reflect.TypeOf((*MockHandler)(nil).ContextName))
}

// SetContextName mocks base method
func (m *MockHandler) SetContextName(contextName string) {
	m.ctrl.Call(m, "SetContextName", contextName)
}

// SetContextName indicates an expected call of SetContextName
func (mr *MockHandlerMockRecorder) SetContextName(contextName interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContextName", reflect.TypeOf((*MockHandler)(nil).SetContextName), contextName)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kaiachain/kaia/kaiax/gov (interfaces: GovModule)

// Package mock_gov is a generated GoMock package.
package mock_gov

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	state "github.com/kaiachain/kaia/blockchain/state"
	types "github.com/kaiachain/kaia/blockchain/types"
	common "github.com/kaiachain/kaia/common"
	gov "github.com/kaiachain/kaia/kaiax/gov"
	rpc "github.com/kaiachain/kaia/networks/rpc"
)

// MockGovModule is a mock of GovModule interface.
type MockGovModule struct {
	ctrl     *gomock.Controller
	recorder *MockGovModuleMockRecorder
}

// MockGovModuleMockRecorder is the mock recorder for MockGovModule.
type MockGovModuleMockRecorder struct {
	mock *MockGovModule
}

// NewMockGovModule creates a new mock instance.
func NewMockGovModule(ctrl *gomock.Controller) *MockGovModule {
	mock := &MockGovModule{ctrl: ctrl}
	mock.recorder = &MockGovModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGovModule) EXPECT() *MockGovModuleMockRecorder {
	return m.recorder
}

// APIs mocks base method.
func (m *MockGovModule) APIs() []rpc.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIs")
	ret0, _ := ret[0].([]rpc.API)
	return ret0
}

// APIs indicates an expected call of APIs.
func (mr *MockGovModuleMockRecorder) APIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIs", reflect.TypeOf((*MockGovModule)(nil).APIs))
}

// EffectiveParamSet mocks base method.
func (m *MockGovModule) EffectiveParamSet(arg0 uint64) gov.ParamSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EffectiveParamSet", arg0)
	ret0, _ := ret[0].(gov.ParamSet)
	return ret0
}

// EffectiveParamSet indicates an expected call of EffectiveParamSet.
func (mr *MockGovModuleMockRecorder) EffectiveParamSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EffectiveParamSet", reflect.TypeOf((*MockGovModule)(nil).EffectiveParamSet), arg0)
}

// FinalizeHeader mocks base method.
func (m *MockGovModule) FinalizeHeader(arg0 *types.Header, arg1 *state.StateDB, arg2 []*types.Transaction, arg3 []*types.Receipt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeHeader", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// FinalizeHeader indicates an expected call of FinalizeHeader.
func (mr *MockGovModuleMockRecorder) FinalizeHeader(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeHeader", reflect.TypeOf((*MockGovModule)(nil).FinalizeHeader), arg0, arg1, arg2, arg3)
}

// PostInsertBlock mocks base method.
func (m *MockGovModule) PostInsertBlock(arg0 *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostInsertBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostInsertBlock indicates an expected call of PostInsertBlock.
func (mr *MockGovModuleMockRecorder) PostInsertBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostInsertBlock", reflect.TypeOf((*MockGovModule)(nil).PostInsertBlock), arg0)
}

// PrepareHeader mocks base method.
func (m *MockGovModule) PrepareHeader(arg0 *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareHeader indicates an expected call of PrepareHeader.
func (mr *MockGovModuleMockRecorder) PrepareHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareHeader", reflect.TypeOf((*MockGovModule)(nil).PrepareHeader), arg0)
}

// RewindDelete mocks base method.
func (m *MockGovModule) RewindDelete(arg0 common.Hash, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RewindDelete", arg0, arg1)
}

// RewindDelete indicates an expected call of RewindDelete.
func (mr *MockGovModuleMockRecorder) RewindDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewindDelete", reflect.TypeOf((*MockGovModule)(nil).RewindDelete), arg0, arg1)
}

// RewindTo mocks base method.
func (m *MockGovModule) RewindTo(arg0 *types.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RewindTo", arg0)
}

// RewindTo indicates an expected call of RewindTo.
func (mr *MockGovModuleMockRecorder) RewindTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewindTo", reflect.TypeOf((*MockGovModule)(nil).RewindTo), arg0)
}

// Start mocks base method.
func (m *MockGovModule) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockGovModuleMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockGovModule)(nil).Start))
}

// Stop mocks base method.
func (m *MockGovModule) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockGovModuleMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockGovModule)(nil).Stop))
}

// VerifyHeader mocks base method.
func (m *MockGovModule) VerifyHeader(arg0 *types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHeader indicates an expected call of VerifyHeader.
func (mr *MockGovModuleMockRecorder) VerifyHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeader", reflect.TypeOf((*MockGovModule)(nil).VerifyHeader), arg0)
}
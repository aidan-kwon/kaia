// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kaiachain/kaia/kaiax/supply (interfaces: SupplyModule)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/kaiachain/kaia/blockchain/types"
	common "github.com/kaiachain/kaia/common"
	supply "github.com/kaiachain/kaia/kaiax/supply"
	rpc "github.com/kaiachain/kaia/networks/rpc"
)

// MockSupplyModule is a mock of SupplyModule interface.
type MockSupplyModule struct {
	ctrl     *gomock.Controller
	recorder *MockSupplyModuleMockRecorder
}

// MockSupplyModuleMockRecorder is the mock recorder for MockSupplyModule.
type MockSupplyModuleMockRecorder struct {
	mock *MockSupplyModule
}

// NewMockSupplyModule creates a new mock instance.
func NewMockSupplyModule(ctrl *gomock.Controller) *MockSupplyModule {
	mock := &MockSupplyModule{ctrl: ctrl}
	mock.recorder = &MockSupplyModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSupplyModule) EXPECT() *MockSupplyModuleMockRecorder {
	return m.recorder
}

// APIs mocks base method.
func (m *MockSupplyModule) APIs() []rpc.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIs")
	ret0, _ := ret[0].([]rpc.API)
	return ret0
}

// APIs indicates an expected call of APIs.
func (mr *MockSupplyModuleMockRecorder) APIs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIs", reflect.TypeOf((*MockSupplyModule)(nil).APIs))
}

// GetTotalSupply mocks base method.
func (m *MockSupplyModule) GetTotalSupply(arg0 uint64) (*supply.TotalSupply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalSupply", arg0)
	ret0, _ := ret[0].(*supply.TotalSupply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalSupply indicates an expected call of GetTotalSupply.
func (mr *MockSupplyModuleMockRecorder) GetTotalSupply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalSupply", reflect.TypeOf((*MockSupplyModule)(nil).GetTotalSupply), arg0)
}

// PostInsertBlock mocks base method.
func (m *MockSupplyModule) PostInsertBlock(arg0 *types.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostInsertBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostInsertBlock indicates an expected call of PostInsertBlock.
func (mr *MockSupplyModuleMockRecorder) PostInsertBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostInsertBlock", reflect.TypeOf((*MockSupplyModule)(nil).PostInsertBlock), arg0)
}

// RewindDelete mocks base method.
func (m *MockSupplyModule) RewindDelete(arg0 common.Hash, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RewindDelete", arg0, arg1)
}

// RewindDelete indicates an expected call of RewindDelete.
func (mr *MockSupplyModuleMockRecorder) RewindDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewindDelete", reflect.TypeOf((*MockSupplyModule)(nil).RewindDelete), arg0, arg1)
}

// RewindTo mocks base method.
func (m *MockSupplyModule) RewindTo(arg0 *types.Block) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RewindTo", arg0)
}

// RewindTo indicates an expected call of RewindTo.
func (mr *MockSupplyModuleMockRecorder) RewindTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RewindTo", reflect.TypeOf((*MockSupplyModule)(nil).RewindTo), arg0)
}

// Start mocks base method.
func (m *MockSupplyModule) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockSupplyModuleMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSupplyModule)(nil).Start))
}

// Stop mocks base method.
func (m *MockSupplyModule) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockSupplyModuleMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSupplyModule)(nil).Stop))
}
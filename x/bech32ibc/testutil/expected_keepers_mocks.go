// Code generated by MockGen. DO NOT EDIT.
// Source: x/bech32ibc/types/expected_keepers.go
//
// Generated by this command:
//
//	mockgen -source=x/bech32ibc/types/expected_keepers.go -package testutil -destination x/bech32ibc/testutil/expected_keepers_mocks.go
//
// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	gomock "go.uber.org/mock/gomock"
)

// MockChannelKeeper is a mock of ChannelKeeper interface.
type MockChannelKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockChannelKeeperMockRecorder
}

// MockChannelKeeperMockRecorder is the mock recorder for MockChannelKeeper.
type MockChannelKeeperMockRecorder struct {
	mock *MockChannelKeeper
}

// NewMockChannelKeeper creates a new mock instance.
func NewMockChannelKeeper(ctrl *gomock.Controller) *MockChannelKeeper {
	mock := &MockChannelKeeper{ctrl: ctrl}
	mock.recorder = &MockChannelKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChannelKeeper) EXPECT() *MockChannelKeeperMockRecorder {
	return m.recorder
}

// GetChannel mocks base method.
func (m *MockChannelKeeper) GetChannel(ctx types.Context, srcPort, srcChan string) (types0.Channel, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannel", ctx, srcPort, srcChan)
	ret0, _ := ret[0].(types0.Channel)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetChannel indicates an expected call of GetChannel.
func (mr *MockChannelKeeperMockRecorder) GetChannel(ctx, srcPort, srcChan any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannel", reflect.TypeOf((*MockChannelKeeper)(nil).GetChannel), ctx, srcPort, srcChan)
}

// MockTransferKeeper is a mock of TransferKeeper interface.
type MockTransferKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockTransferKeeperMockRecorder
}

// MockTransferKeeperMockRecorder is the mock recorder for MockTransferKeeper.
type MockTransferKeeperMockRecorder struct {
	mock *MockTransferKeeper
}

// NewMockTransferKeeper creates a new mock instance.
func NewMockTransferKeeper(ctrl *gomock.Controller) *MockTransferKeeper {
	mock := &MockTransferKeeper{ctrl: ctrl}
	mock.recorder = &MockTransferKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransferKeeper) EXPECT() *MockTransferKeeperMockRecorder {
	return m.recorder
}

// GetPort mocks base method.
func (m *MockTransferKeeper) GetPort(ctx types.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPort indicates an expected call of GetPort.
func (mr *MockTransferKeeperMockRecorder) GetPort(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockTransferKeeper)(nil).GetPort), ctx)
}

// MockNftTransferKeeper is a mock of NftTransferKeeper interface.
type MockNftTransferKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockNftTransferKeeperMockRecorder
}

// MockNftTransferKeeperMockRecorder is the mock recorder for MockNftTransferKeeper.
type MockNftTransferKeeperMockRecorder struct {
	mock *MockNftTransferKeeper
}

// NewMockNftTransferKeeper creates a new mock instance.
func NewMockNftTransferKeeper(ctrl *gomock.Controller) *MockNftTransferKeeper {
	mock := &MockNftTransferKeeper{ctrl: ctrl}
	mock.recorder = &MockNftTransferKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNftTransferKeeper) EXPECT() *MockNftTransferKeeperMockRecorder {
	return m.recorder
}

// GetPort mocks base method.
func (m *MockNftTransferKeeper) GetPort(ctx types.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPort", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPort indicates an expected call of GetPort.
func (mr *MockNftTransferKeeperMockRecorder) GetPort(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPort", reflect.TypeOf((*MockNftTransferKeeper)(nil).GetPort), ctx)
}
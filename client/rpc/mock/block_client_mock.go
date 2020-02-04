// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tendermint/tendermint@v0.32.8/rpc/client/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/tendermint/tendermint/libs/common"
	log "github.com/tendermint/tendermint/libs/log"
	client "github.com/tendermint/tendermint/rpc/client"
	types "github.com/tendermint/tendermint/rpc/core/types"
	types0 "github.com/tendermint/tendermint/types"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockClient) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockClient)(nil).Start))
}

// OnStart mocks base method
func (m *MockClient) OnStart() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnStart")
	ret0, _ := ret[0].(error)
	return ret0
}

// OnStart indicates an expected call of OnStart
func (mr *MockClientMockRecorder) OnStart() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnStart", reflect.TypeOf((*MockClient)(nil).OnStart))
}

// Stop mocks base method
func (m *MockClient) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop))
}

// OnStop mocks base method
func (m *MockClient) OnStop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnStop")
}

// OnStop indicates an expected call of OnStop
func (mr *MockClientMockRecorder) OnStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnStop", reflect.TypeOf((*MockClient)(nil).OnStop))
}

// Reset mocks base method
func (m *MockClient) Reset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset
func (mr *MockClientMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockClient)(nil).Reset))
}

// OnReset mocks base method
func (m *MockClient) OnReset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnReset")
	ret0, _ := ret[0].(error)
	return ret0
}

// OnReset indicates an expected call of OnReset
func (mr *MockClientMockRecorder) OnReset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnReset", reflect.TypeOf((*MockClient)(nil).OnReset))
}

// IsRunning mocks base method
func (m *MockClient) IsRunning() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRunning")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRunning indicates an expected call of IsRunning
func (mr *MockClientMockRecorder) IsRunning() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRunning", reflect.TypeOf((*MockClient)(nil).IsRunning))
}

// Quit mocks base method
func (m *MockClient) Quit() <-chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit")
	ret0, _ := ret[0].(<-chan struct{})
	return ret0
}

// Quit indicates an expected call of Quit
func (mr *MockClientMockRecorder) Quit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockClient)(nil).Quit))
}

// String mocks base method
func (m *MockClient) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String
func (mr *MockClientMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockClient)(nil).String))
}

// SetLogger mocks base method
func (m *MockClient) SetLogger(arg0 log.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLogger", arg0)
}

// SetLogger indicates an expected call of SetLogger
func (mr *MockClientMockRecorder) SetLogger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLogger", reflect.TypeOf((*MockClient)(nil).SetLogger), arg0)
}

// ABCIInfo mocks base method
func (m *MockClient) ABCIInfo() (*types.ResultABCIInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIInfo")
	ret0, _ := ret[0].(*types.ResultABCIInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIInfo indicates an expected call of ABCIInfo
func (mr *MockClientMockRecorder) ABCIInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIInfo", reflect.TypeOf((*MockClient)(nil).ABCIInfo))
}

// ABCIQuery mocks base method
func (m *MockClient) ABCIQuery(path string, data common.HexBytes) (*types.ResultABCIQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIQuery", path, data)
	ret0, _ := ret[0].(*types.ResultABCIQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIQuery indicates an expected call of ABCIQuery
func (mr *MockClientMockRecorder) ABCIQuery(path, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIQuery", reflect.TypeOf((*MockClient)(nil).ABCIQuery), path, data)
}

// ABCIQueryWithOptions mocks base method
func (m *MockClient) ABCIQueryWithOptions(path string, data common.HexBytes, opts client.ABCIQueryOptions) (*types.ResultABCIQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIQueryWithOptions", path, data, opts)
	ret0, _ := ret[0].(*types.ResultABCIQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIQueryWithOptions indicates an expected call of ABCIQueryWithOptions
func (mr *MockClientMockRecorder) ABCIQueryWithOptions(path, data, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIQueryWithOptions", reflect.TypeOf((*MockClient)(nil).ABCIQueryWithOptions), path, data, opts)
}

// BroadcastTxCommit mocks base method
func (m *MockClient) BroadcastTxCommit(tx types0.Tx) (*types.ResultBroadcastTxCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxCommit", tx)
	ret0, _ := ret[0].(*types.ResultBroadcastTxCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxCommit indicates an expected call of BroadcastTxCommit
func (mr *MockClientMockRecorder) BroadcastTxCommit(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxCommit", reflect.TypeOf((*MockClient)(nil).BroadcastTxCommit), tx)
}

// BroadcastTxAsync mocks base method
func (m *MockClient) BroadcastTxAsync(tx types0.Tx) (*types.ResultBroadcastTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxAsync", tx)
	ret0, _ := ret[0].(*types.ResultBroadcastTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxAsync indicates an expected call of BroadcastTxAsync
func (mr *MockClientMockRecorder) BroadcastTxAsync(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxAsync", reflect.TypeOf((*MockClient)(nil).BroadcastTxAsync), tx)
}

// BroadcastTxSync mocks base method
func (m *MockClient) BroadcastTxSync(tx types0.Tx) (*types.ResultBroadcastTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxSync", tx)
	ret0, _ := ret[0].(*types.ResultBroadcastTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxSync indicates an expected call of BroadcastTxSync
func (mr *MockClientMockRecorder) BroadcastTxSync(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxSync", reflect.TypeOf((*MockClient)(nil).BroadcastTxSync), tx)
}

// Subscribe mocks base method
func (m *MockClient) Subscribe(ctx context.Context, subscriber, query string, outCapacity ...int) (<-chan types.ResultEvent, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subscriber, query}
	for _, a := range outCapacity {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(<-chan types.ResultEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockClientMockRecorder) Subscribe(ctx, subscriber, query interface{}, outCapacity ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subscriber, query}, outCapacity...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockClient)(nil).Subscribe), varargs...)
}

// Unsubscribe mocks base method
func (m *MockClient) Unsubscribe(ctx context.Context, subscriber, query string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ctx, subscriber, query)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockClientMockRecorder) Unsubscribe(ctx, subscriber, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockClient)(nil).Unsubscribe), ctx, subscriber, query)
}

// UnsubscribeAll mocks base method
func (m *MockClient) UnsubscribeAll(ctx context.Context, subscriber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeAll", ctx, subscriber)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeAll indicates an expected call of UnsubscribeAll
func (mr *MockClientMockRecorder) UnsubscribeAll(ctx, subscriber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeAll", reflect.TypeOf((*MockClient)(nil).UnsubscribeAll), ctx, subscriber)
}

// Genesis mocks base method
func (m *MockClient) Genesis() (*types.ResultGenesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(*types.ResultGenesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Genesis indicates an expected call of Genesis
func (mr *MockClientMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockClient)(nil).Genesis))
}

// BlockchainInfo mocks base method
func (m *MockClient) BlockchainInfo(minHeight, maxHeight int64) (*types.ResultBlockchainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockchainInfo", minHeight, maxHeight)
	ret0, _ := ret[0].(*types.ResultBlockchainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockchainInfo indicates an expected call of BlockchainInfo
func (mr *MockClientMockRecorder) BlockchainInfo(minHeight, maxHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockchainInfo", reflect.TypeOf((*MockClient)(nil).BlockchainInfo), minHeight, maxHeight)
}

// NetInfo mocks base method
func (m *MockClient) NetInfo() (*types.ResultNetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetInfo")
	ret0, _ := ret[0].(*types.ResultNetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetInfo indicates an expected call of NetInfo
func (mr *MockClientMockRecorder) NetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetInfo", reflect.TypeOf((*MockClient)(nil).NetInfo))
}

// DumpConsensusState mocks base method
func (m *MockClient) DumpConsensusState() (*types.ResultDumpConsensusState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpConsensusState")
	ret0, _ := ret[0].(*types.ResultDumpConsensusState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpConsensusState indicates an expected call of DumpConsensusState
func (mr *MockClientMockRecorder) DumpConsensusState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpConsensusState", reflect.TypeOf((*MockClient)(nil).DumpConsensusState))
}

// ConsensusState mocks base method
func (m *MockClient) ConsensusState() (*types.ResultConsensusState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsensusState")
	ret0, _ := ret[0].(*types.ResultConsensusState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsensusState indicates an expected call of ConsensusState
func (mr *MockClientMockRecorder) ConsensusState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsensusState", reflect.TypeOf((*MockClient)(nil).ConsensusState))
}

// Health mocks base method
func (m *MockClient) Health() (*types.ResultHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(*types.ResultHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockClientMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockClient)(nil).Health))
}

// Block mocks base method
func (m *MockClient) Block(height *int64) (*types.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", height)
	ret0, _ := ret[0].(*types.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Block indicates an expected call of Block
func (mr *MockClientMockRecorder) Block(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockClient)(nil).Block), height)
}

// BlockResults mocks base method
func (m *MockClient) BlockResults(height *int64) (*types.ResultBlockResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockResults", height)
	ret0, _ := ret[0].(*types.ResultBlockResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockResults indicates an expected call of BlockResults
func (mr *MockClientMockRecorder) BlockResults(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockResults", reflect.TypeOf((*MockClient)(nil).BlockResults), height)
}

// Commit mocks base method
func (m *MockClient) Commit(height *int64) (*types.ResultCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", height)
	ret0, _ := ret[0].(*types.ResultCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit
func (mr *MockClientMockRecorder) Commit(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockClient)(nil).Commit), height)
}

// Validators mocks base method
func (m *MockClient) Validators(height *int64) (*types.ResultValidators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validators", height)
	ret0, _ := ret[0].(*types.ResultValidators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validators indicates an expected call of Validators
func (mr *MockClientMockRecorder) Validators(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validators", reflect.TypeOf((*MockClient)(nil).Validators), height)
}

// Tx mocks base method
func (m *MockClient) Tx(hash []byte, prove bool) (*types.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", hash, prove)
	ret0, _ := ret[0].(*types.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx
func (mr *MockClientMockRecorder) Tx(hash, prove interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockClient)(nil).Tx), hash, prove)
}

// TxSearch mocks base method
func (m *MockClient) TxSearch(query string, prove bool, page, perPage int) (*types.ResultTxSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxSearch", query, prove, page, perPage)
	ret0, _ := ret[0].(*types.ResultTxSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxSearch indicates an expected call of TxSearch
func (mr *MockClientMockRecorder) TxSearch(query, prove, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxSearch", reflect.TypeOf((*MockClient)(nil).TxSearch), query, prove, page, perPage)
}

// Status mocks base method
func (m *MockClient) Status() (*types.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(*types.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockClientMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockClient)(nil).Status))
}

// BroadcastEvidence mocks base method
func (m *MockClient) BroadcastEvidence(ev types0.Evidence) (*types.ResultBroadcastEvidence, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastEvidence", ev)
	ret0, _ := ret[0].(*types.ResultBroadcastEvidence)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastEvidence indicates an expected call of BroadcastEvidence
func (mr *MockClientMockRecorder) BroadcastEvidence(ev interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastEvidence", reflect.TypeOf((*MockClient)(nil).BroadcastEvidence), ev)
}

// UnconfirmedTxs mocks base method
func (m *MockClient) UnconfirmedTxs(limit int) (*types.ResultUnconfirmedTxs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnconfirmedTxs", limit)
	ret0, _ := ret[0].(*types.ResultUnconfirmedTxs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnconfirmedTxs indicates an expected call of UnconfirmedTxs
func (mr *MockClientMockRecorder) UnconfirmedTxs(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnconfirmedTxs", reflect.TypeOf((*MockClient)(nil).UnconfirmedTxs), limit)
}

// NumUnconfirmedTxs mocks base method
func (m *MockClient) NumUnconfirmedTxs() (*types.ResultUnconfirmedTxs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumUnconfirmedTxs")
	ret0, _ := ret[0].(*types.ResultUnconfirmedTxs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NumUnconfirmedTxs indicates an expected call of NumUnconfirmedTxs
func (mr *MockClientMockRecorder) NumUnconfirmedTxs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumUnconfirmedTxs", reflect.TypeOf((*MockClient)(nil).NumUnconfirmedTxs))
}

// MockABCIClient is a mock of ABCIClient interface
type MockABCIClient struct {
	ctrl     *gomock.Controller
	recorder *MockABCIClientMockRecorder
}

// MockABCIClientMockRecorder is the mock recorder for MockABCIClient
type MockABCIClientMockRecorder struct {
	mock *MockABCIClient
}

// NewMockABCIClient creates a new mock instance
func NewMockABCIClient(ctrl *gomock.Controller) *MockABCIClient {
	mock := &MockABCIClient{ctrl: ctrl}
	mock.recorder = &MockABCIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockABCIClient) EXPECT() *MockABCIClientMockRecorder {
	return m.recorder
}

// ABCIInfo mocks base method
func (m *MockABCIClient) ABCIInfo() (*types.ResultABCIInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIInfo")
	ret0, _ := ret[0].(*types.ResultABCIInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIInfo indicates an expected call of ABCIInfo
func (mr *MockABCIClientMockRecorder) ABCIInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIInfo", reflect.TypeOf((*MockABCIClient)(nil).ABCIInfo))
}

// ABCIQuery mocks base method
func (m *MockABCIClient) ABCIQuery(path string, data common.HexBytes) (*types.ResultABCIQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIQuery", path, data)
	ret0, _ := ret[0].(*types.ResultABCIQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIQuery indicates an expected call of ABCIQuery
func (mr *MockABCIClientMockRecorder) ABCIQuery(path, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIQuery", reflect.TypeOf((*MockABCIClient)(nil).ABCIQuery), path, data)
}

// ABCIQueryWithOptions mocks base method
func (m *MockABCIClient) ABCIQueryWithOptions(path string, data common.HexBytes, opts client.ABCIQueryOptions) (*types.ResultABCIQuery, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ABCIQueryWithOptions", path, data, opts)
	ret0, _ := ret[0].(*types.ResultABCIQuery)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ABCIQueryWithOptions indicates an expected call of ABCIQueryWithOptions
func (mr *MockABCIClientMockRecorder) ABCIQueryWithOptions(path, data, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ABCIQueryWithOptions", reflect.TypeOf((*MockABCIClient)(nil).ABCIQueryWithOptions), path, data, opts)
}

// BroadcastTxCommit mocks base method
func (m *MockABCIClient) BroadcastTxCommit(tx types0.Tx) (*types.ResultBroadcastTxCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxCommit", tx)
	ret0, _ := ret[0].(*types.ResultBroadcastTxCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxCommit indicates an expected call of BroadcastTxCommit
func (mr *MockABCIClientMockRecorder) BroadcastTxCommit(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxCommit", reflect.TypeOf((*MockABCIClient)(nil).BroadcastTxCommit), tx)
}

// BroadcastTxAsync mocks base method
func (m *MockABCIClient) BroadcastTxAsync(tx types0.Tx) (*types.ResultBroadcastTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxAsync", tx)
	ret0, _ := ret[0].(*types.ResultBroadcastTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxAsync indicates an expected call of BroadcastTxAsync
func (mr *MockABCIClientMockRecorder) BroadcastTxAsync(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxAsync", reflect.TypeOf((*MockABCIClient)(nil).BroadcastTxAsync), tx)
}

// BroadcastTxSync mocks base method
func (m *MockABCIClient) BroadcastTxSync(tx types0.Tx) (*types.ResultBroadcastTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastTxSync", tx)
	ret0, _ := ret[0].(*types.ResultBroadcastTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastTxSync indicates an expected call of BroadcastTxSync
func (mr *MockABCIClientMockRecorder) BroadcastTxSync(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastTxSync", reflect.TypeOf((*MockABCIClient)(nil).BroadcastTxSync), tx)
}

// MockSignClient is a mock of SignClient interface
type MockSignClient struct {
	ctrl     *gomock.Controller
	recorder *MockSignClientMockRecorder
}

// MockSignClientMockRecorder is the mock recorder for MockSignClient
type MockSignClientMockRecorder struct {
	mock *MockSignClient
}

// NewMockSignClient creates a new mock instance
func NewMockSignClient(ctrl *gomock.Controller) *MockSignClient {
	mock := &MockSignClient{ctrl: ctrl}
	mock.recorder = &MockSignClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSignClient) EXPECT() *MockSignClientMockRecorder {
	return m.recorder
}

// Block mocks base method
func (m *MockSignClient) Block(height *int64) (*types.ResultBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Block", height)
	ret0, _ := ret[0].(*types.ResultBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Block indicates an expected call of Block
func (mr *MockSignClientMockRecorder) Block(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Block", reflect.TypeOf((*MockSignClient)(nil).Block), height)
}

// BlockResults mocks base method
func (m *MockSignClient) BlockResults(height *int64) (*types.ResultBlockResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockResults", height)
	ret0, _ := ret[0].(*types.ResultBlockResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockResults indicates an expected call of BlockResults
func (mr *MockSignClientMockRecorder) BlockResults(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockResults", reflect.TypeOf((*MockSignClient)(nil).BlockResults), height)
}

// Commit mocks base method
func (m *MockSignClient) Commit(height *int64) (*types.ResultCommit, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", height)
	ret0, _ := ret[0].(*types.ResultCommit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit
func (mr *MockSignClientMockRecorder) Commit(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockSignClient)(nil).Commit), height)
}

// Validators mocks base method
func (m *MockSignClient) Validators(height *int64) (*types.ResultValidators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validators", height)
	ret0, _ := ret[0].(*types.ResultValidators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validators indicates an expected call of Validators
func (mr *MockSignClientMockRecorder) Validators(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validators", reflect.TypeOf((*MockSignClient)(nil).Validators), height)
}

// Tx mocks base method
func (m *MockSignClient) Tx(hash []byte, prove bool) (*types.ResultTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tx", hash, prove)
	ret0, _ := ret[0].(*types.ResultTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tx indicates an expected call of Tx
func (mr *MockSignClientMockRecorder) Tx(hash, prove interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tx", reflect.TypeOf((*MockSignClient)(nil).Tx), hash, prove)
}

// TxSearch mocks base method
func (m *MockSignClient) TxSearch(query string, prove bool, page, perPage int) (*types.ResultTxSearch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxSearch", query, prove, page, perPage)
	ret0, _ := ret[0].(*types.ResultTxSearch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TxSearch indicates an expected call of TxSearch
func (mr *MockSignClientMockRecorder) TxSearch(query, prove, page, perPage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxSearch", reflect.TypeOf((*MockSignClient)(nil).TxSearch), query, prove, page, perPage)
}

// MockHistoryClient is a mock of HistoryClient interface
type MockHistoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockHistoryClientMockRecorder
}

// MockHistoryClientMockRecorder is the mock recorder for MockHistoryClient
type MockHistoryClientMockRecorder struct {
	mock *MockHistoryClient
}

// NewMockHistoryClient creates a new mock instance
func NewMockHistoryClient(ctrl *gomock.Controller) *MockHistoryClient {
	mock := &MockHistoryClient{ctrl: ctrl}
	mock.recorder = &MockHistoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHistoryClient) EXPECT() *MockHistoryClientMockRecorder {
	return m.recorder
}

// Genesis mocks base method
func (m *MockHistoryClient) Genesis() (*types.ResultGenesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Genesis")
	ret0, _ := ret[0].(*types.ResultGenesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Genesis indicates an expected call of Genesis
func (mr *MockHistoryClientMockRecorder) Genesis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Genesis", reflect.TypeOf((*MockHistoryClient)(nil).Genesis))
}

// BlockchainInfo mocks base method
func (m *MockHistoryClient) BlockchainInfo(minHeight, maxHeight int64) (*types.ResultBlockchainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockchainInfo", minHeight, maxHeight)
	ret0, _ := ret[0].(*types.ResultBlockchainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockchainInfo indicates an expected call of BlockchainInfo
func (mr *MockHistoryClientMockRecorder) BlockchainInfo(minHeight, maxHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockchainInfo", reflect.TypeOf((*MockHistoryClient)(nil).BlockchainInfo), minHeight, maxHeight)
}

// MockStatusClient is a mock of StatusClient interface
type MockStatusClient struct {
	ctrl     *gomock.Controller
	recorder *MockStatusClientMockRecorder
}

// MockStatusClientMockRecorder is the mock recorder for MockStatusClient
type MockStatusClientMockRecorder struct {
	mock *MockStatusClient
}

// NewMockStatusClient creates a new mock instance
func NewMockStatusClient(ctrl *gomock.Controller) *MockStatusClient {
	mock := &MockStatusClient{ctrl: ctrl}
	mock.recorder = &MockStatusClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatusClient) EXPECT() *MockStatusClientMockRecorder {
	return m.recorder
}

// Status mocks base method
func (m *MockStatusClient) Status() (*types.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(*types.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockStatusClientMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockStatusClient)(nil).Status))
}

// MockNetworkClient is a mock of NetworkClient interface
type MockNetworkClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkClientMockRecorder
}

// MockNetworkClientMockRecorder is the mock recorder for MockNetworkClient
type MockNetworkClientMockRecorder struct {
	mock *MockNetworkClient
}

// NewMockNetworkClient creates a new mock instance
func NewMockNetworkClient(ctrl *gomock.Controller) *MockNetworkClient {
	mock := &MockNetworkClient{ctrl: ctrl}
	mock.recorder = &MockNetworkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkClient) EXPECT() *MockNetworkClientMockRecorder {
	return m.recorder
}

// NetInfo mocks base method
func (m *MockNetworkClient) NetInfo() (*types.ResultNetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetInfo")
	ret0, _ := ret[0].(*types.ResultNetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetInfo indicates an expected call of NetInfo
func (mr *MockNetworkClientMockRecorder) NetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetInfo", reflect.TypeOf((*MockNetworkClient)(nil).NetInfo))
}

// DumpConsensusState mocks base method
func (m *MockNetworkClient) DumpConsensusState() (*types.ResultDumpConsensusState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpConsensusState")
	ret0, _ := ret[0].(*types.ResultDumpConsensusState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DumpConsensusState indicates an expected call of DumpConsensusState
func (mr *MockNetworkClientMockRecorder) DumpConsensusState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpConsensusState", reflect.TypeOf((*MockNetworkClient)(nil).DumpConsensusState))
}

// ConsensusState mocks base method
func (m *MockNetworkClient) ConsensusState() (*types.ResultConsensusState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsensusState")
	ret0, _ := ret[0].(*types.ResultConsensusState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsensusState indicates an expected call of ConsensusState
func (mr *MockNetworkClientMockRecorder) ConsensusState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsensusState", reflect.TypeOf((*MockNetworkClient)(nil).ConsensusState))
}

// Health mocks base method
func (m *MockNetworkClient) Health() (*types.ResultHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(*types.ResultHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockNetworkClientMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockNetworkClient)(nil).Health))
}

// MockEventsClient is a mock of EventsClient interface
type MockEventsClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventsClientMockRecorder
}

// MockEventsClientMockRecorder is the mock recorder for MockEventsClient
type MockEventsClientMockRecorder struct {
	mock *MockEventsClient
}

// NewMockEventsClient creates a new mock instance
func NewMockEventsClient(ctrl *gomock.Controller) *MockEventsClient {
	mock := &MockEventsClient{ctrl: ctrl}
	mock.recorder = &MockEventsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventsClient) EXPECT() *MockEventsClientMockRecorder {
	return m.recorder
}

// Subscribe mocks base method
func (m *MockEventsClient) Subscribe(ctx context.Context, subscriber, query string, outCapacity ...int) (<-chan types.ResultEvent, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, subscriber, query}
	for _, a := range outCapacity {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(<-chan types.ResultEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockEventsClientMockRecorder) Subscribe(ctx, subscriber, query interface{}, outCapacity ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, subscriber, query}, outCapacity...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockEventsClient)(nil).Subscribe), varargs...)
}

// Unsubscribe mocks base method
func (m *MockEventsClient) Unsubscribe(ctx context.Context, subscriber, query string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", ctx, subscriber, query)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockEventsClientMockRecorder) Unsubscribe(ctx, subscriber, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockEventsClient)(nil).Unsubscribe), ctx, subscriber, query)
}

// UnsubscribeAll mocks base method
func (m *MockEventsClient) UnsubscribeAll(ctx context.Context, subscriber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnsubscribeAll", ctx, subscriber)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnsubscribeAll indicates an expected call of UnsubscribeAll
func (mr *MockEventsClientMockRecorder) UnsubscribeAll(ctx, subscriber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeAll", reflect.TypeOf((*MockEventsClient)(nil).UnsubscribeAll), ctx, subscriber)
}

// MockMempoolClient is a mock of MempoolClient interface
type MockMempoolClient struct {
	ctrl     *gomock.Controller
	recorder *MockMempoolClientMockRecorder
}

// MockMempoolClientMockRecorder is the mock recorder for MockMempoolClient
type MockMempoolClientMockRecorder struct {
	mock *MockMempoolClient
}

// NewMockMempoolClient creates a new mock instance
func NewMockMempoolClient(ctrl *gomock.Controller) *MockMempoolClient {
	mock := &MockMempoolClient{ctrl: ctrl}
	mock.recorder = &MockMempoolClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMempoolClient) EXPECT() *MockMempoolClientMockRecorder {
	return m.recorder
}

// UnconfirmedTxs mocks base method
func (m *MockMempoolClient) UnconfirmedTxs(limit int) (*types.ResultUnconfirmedTxs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnconfirmedTxs", limit)
	ret0, _ := ret[0].(*types.ResultUnconfirmedTxs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnconfirmedTxs indicates an expected call of UnconfirmedTxs
func (mr *MockMempoolClientMockRecorder) UnconfirmedTxs(limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnconfirmedTxs", reflect.TypeOf((*MockMempoolClient)(nil).UnconfirmedTxs), limit)
}

// NumUnconfirmedTxs mocks base method
func (m *MockMempoolClient) NumUnconfirmedTxs() (*types.ResultUnconfirmedTxs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumUnconfirmedTxs")
	ret0, _ := ret[0].(*types.ResultUnconfirmedTxs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NumUnconfirmedTxs indicates an expected call of NumUnconfirmedTxs
func (mr *MockMempoolClientMockRecorder) NumUnconfirmedTxs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumUnconfirmedTxs", reflect.TypeOf((*MockMempoolClient)(nil).NumUnconfirmedTxs))
}

// MockEvidenceClient is a mock of EvidenceClient interface
type MockEvidenceClient struct {
	ctrl     *gomock.Controller
	recorder *MockEvidenceClientMockRecorder
}

// MockEvidenceClientMockRecorder is the mock recorder for MockEvidenceClient
type MockEvidenceClientMockRecorder struct {
	mock *MockEvidenceClient
}

// NewMockEvidenceClient creates a new mock instance
func NewMockEvidenceClient(ctrl *gomock.Controller) *MockEvidenceClient {
	mock := &MockEvidenceClient{ctrl: ctrl}
	mock.recorder = &MockEvidenceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEvidenceClient) EXPECT() *MockEvidenceClientMockRecorder {
	return m.recorder
}

// BroadcastEvidence mocks base method
func (m *MockEvidenceClient) BroadcastEvidence(ev types0.Evidence) (*types.ResultBroadcastEvidence, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadcastEvidence", ev)
	ret0, _ := ret[0].(*types.ResultBroadcastEvidence)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BroadcastEvidence indicates an expected call of BroadcastEvidence
func (mr *MockEvidenceClientMockRecorder) BroadcastEvidence(ev interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastEvidence", reflect.TypeOf((*MockEvidenceClient)(nil).BroadcastEvidence), ev)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ../client.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/puppetlabs/pe-sdk-go/app/puppet-code/api/client"
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

// GetClient mocks base method
func (m *MockClient) GetClient() (*client.PuppetCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(*client.PuppetCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient
func (mr *MockClientMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockClient)(nil).GetClient))
}

// EnableCN mocks base method
func (m *MockClient) EnableCN() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EnableCN")
}

// EnableCN indicates an expected call of EnableCN
func (mr *MockClientMockRecorder) EnableCN() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCN", reflect.TypeOf((*MockClient)(nil).EnableCN))
}

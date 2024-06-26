// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/treeverse/lakefs/pkg/authentication/apiclient (interfaces: ClientWithResponsesInterface)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apiclient "github.com/treeverse/lakefs/pkg/authentication/apiclient"
)

// MockClientWithResponsesInterface is a mock of ClientWithResponsesInterface interface.
type MockClientWithResponsesInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientWithResponsesInterfaceMockRecorder
}

// MockClientWithResponsesInterfaceMockRecorder is the mock recorder for MockClientWithResponsesInterface.
type MockClientWithResponsesInterfaceMockRecorder struct {
	mock *MockClientWithResponsesInterface
}

// NewMockClientWithResponsesInterface creates a new mock instance.
func NewMockClientWithResponsesInterface(ctrl *gomock.Controller) *MockClientWithResponsesInterface {
	mock := &MockClientWithResponsesInterface{ctrl: ctrl}
	mock.recorder = &MockClientWithResponsesInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientWithResponsesInterface) EXPECT() *MockClientWithResponsesInterfaceMockRecorder {
	return m.recorder
}

// ExternalPrincipalLoginWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) ExternalPrincipalLoginWithBodyWithResponse(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 ...apiclient.RequestEditorFn) (*apiclient.ExternalPrincipalLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExternalPrincipalLoginWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.ExternalPrincipalLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExternalPrincipalLoginWithBodyWithResponse indicates an expected call of ExternalPrincipalLoginWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) ExternalPrincipalLoginWithBodyWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalPrincipalLoginWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).ExternalPrincipalLoginWithBodyWithResponse), varargs...)
}

// ExternalPrincipalLoginWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) ExternalPrincipalLoginWithResponse(arg0 context.Context, arg1 apiclient.ExternalPrincipalLoginJSONRequestBody, arg2 ...apiclient.RequestEditorFn) (*apiclient.ExternalPrincipalLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExternalPrincipalLoginWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.ExternalPrincipalLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExternalPrincipalLoginWithResponse indicates an expected call of ExternalPrincipalLoginWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) ExternalPrincipalLoginWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalPrincipalLoginWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).ExternalPrincipalLoginWithResponse), varargs...)
}

// HealthCheckWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) HealthCheckWithResponse(arg0 context.Context, arg1 ...apiclient.RequestEditorFn) (*apiclient.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HealthCheckWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheckWithResponse indicates an expected call of HealthCheckWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) HealthCheckWithResponse(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheckWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).HealthCheckWithResponse), varargs...)
}

// LDAPLoginWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) LDAPLoginWithBodyWithResponse(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 ...apiclient.RequestEditorFn) (*apiclient.LDAPLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LDAPLoginWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.LDAPLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LDAPLoginWithBodyWithResponse indicates an expected call of LDAPLoginWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) LDAPLoginWithBodyWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LDAPLoginWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).LDAPLoginWithBodyWithResponse), varargs...)
}

// LDAPLoginWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) LDAPLoginWithResponse(arg0 context.Context, arg1 apiclient.LDAPLoginJSONRequestBody, arg2 ...apiclient.RequestEditorFn) (*apiclient.LDAPLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "LDAPLoginWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.LDAPLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LDAPLoginWithResponse indicates an expected call of LDAPLoginWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) LDAPLoginWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LDAPLoginWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).LDAPLoginWithResponse), varargs...)
}

// OauthCallbackWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) OauthCallbackWithResponse(arg0 context.Context, arg1 ...apiclient.RequestEditorFn) (*apiclient.OauthCallbackResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OauthCallbackWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.OauthCallbackResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OauthCallbackWithResponse indicates an expected call of OauthCallbackWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) OauthCallbackWithResponse(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OauthCallbackWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).OauthCallbackWithResponse), varargs...)
}

// STSLoginWithBodyWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) STSLoginWithBodyWithResponse(arg0 context.Context, arg1 string, arg2 io.Reader, arg3 ...apiclient.RequestEditorFn) (*apiclient.STSLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "STSLoginWithBodyWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.STSLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// STSLoginWithBodyWithResponse indicates an expected call of STSLoginWithBodyWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) STSLoginWithBodyWithResponse(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "STSLoginWithBodyWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).STSLoginWithBodyWithResponse), varargs...)
}

// STSLoginWithResponse mocks base method.
func (m *MockClientWithResponsesInterface) STSLoginWithResponse(arg0 context.Context, arg1 apiclient.STSLoginJSONRequestBody, arg2 ...apiclient.RequestEditorFn) (*apiclient.STSLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "STSLoginWithResponse", varargs...)
	ret0, _ := ret[0].(*apiclient.STSLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// STSLoginWithResponse indicates an expected call of STSLoginWithResponse.
func (mr *MockClientWithResponsesInterfaceMockRecorder) STSLoginWithResponse(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "STSLoginWithResponse", reflect.TypeOf((*MockClientWithResponsesInterface)(nil).STSLoginWithResponse), varargs...)
}

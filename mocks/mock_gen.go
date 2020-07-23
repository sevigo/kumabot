// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sevigo/kumabot/pkg/core (interfaces: KumaBots,KumaBot,LineBot)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	linebot "github.com/line/line-bot-sdk-go/linebot"
	core "github.com/sevigo/kumabot/pkg/core"
	http "net/http"
	reflect "reflect"
)

// MockKumaBots is a mock of KumaBots interface
type MockKumaBots struct {
	ctrl     *gomock.Controller
	recorder *MockKumaBotsMockRecorder
}

// MockKumaBotsMockRecorder is the mock recorder for MockKumaBots
type MockKumaBotsMockRecorder struct {
	mock *MockKumaBots
}

// NewMockKumaBots creates a new mock instance
func NewMockKumaBots(ctrl *gomock.Controller) *MockKumaBots {
	mock := &MockKumaBots{ctrl: ctrl}
	mock.recorder = &MockKumaBotsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKumaBots) EXPECT() *MockKumaBotsMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockKumaBots) All() []core.KumaBot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]core.KumaBot)
	return ret0
}

// All indicates an expected call of All
func (mr *MockKumaBotsMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockKumaBots)(nil).All))
}

// MockKumaBot is a mock of KumaBot interface
type MockKumaBot struct {
	ctrl     *gomock.Controller
	recorder *MockKumaBotMockRecorder
}

// MockKumaBotMockRecorder is the mock recorder for MockKumaBot
type MockKumaBotMockRecorder struct {
	mock *MockKumaBot
}

// NewMockKumaBot creates a new mock instance
func NewMockKumaBot(ctrl *gomock.Controller) *MockKumaBot {
	mock := &MockKumaBot{ctrl: ctrl}
	mock.recorder = &MockKumaBotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKumaBot) EXPECT() *MockKumaBotMockRecorder {
	return m.recorder
}

// Match mocks base method
func (m *MockKumaBot) Match(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Match", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Match indicates an expected call of Match
func (mr *MockKumaBotMockRecorder) Match(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Match", reflect.TypeOf((*MockKumaBot)(nil).Match), arg0)
}

// Name mocks base method
func (m *MockKumaBot) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockKumaBotMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockKumaBot)(nil).Name))
}

// Run mocks base method
func (m *MockKumaBot) Run(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockKumaBotMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockKumaBot)(nil).Run), arg0)
}

// MockLineBot is a mock of LineBot interface
type MockLineBot struct {
	ctrl     *gomock.Controller
	recorder *MockLineBotMockRecorder
}

// MockLineBotMockRecorder is the mock recorder for MockLineBot
type MockLineBotMockRecorder struct {
	mock *MockLineBot
}

// NewMockLineBot creates a new mock instance
func NewMockLineBot(ctrl *gomock.Controller) *MockLineBot {
	mock := &MockLineBot{ctrl: ctrl}
	mock.recorder = &MockLineBotMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLineBot) EXPECT() *MockLineBotMockRecorder {
	return m.recorder
}

// ParseRequest mocks base method
func (m *MockLineBot) ParseRequest(arg0 *http.Request) ([]*linebot.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseRequest", arg0)
	ret0, _ := ret[0].([]*linebot.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseRequest indicates an expected call of ParseRequest
func (mr *MockLineBotMockRecorder) ParseRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRequest", reflect.TypeOf((*MockLineBot)(nil).ParseRequest), arg0)
}

// ReplyMessage mocks base method
func (m *MockLineBot) ReplyMessage(arg0 string, arg1 ...linebot.SendingMessage) (*linebot.BasicResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReplyMessage", varargs...)
	ret0, _ := ret[0].(*linebot.BasicResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplyMessage indicates an expected call of ReplyMessage
func (mr *MockLineBotMockRecorder) ReplyMessage(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplyMessage", reflect.TypeOf((*MockLineBot)(nil).ReplyMessage), varargs...)
}

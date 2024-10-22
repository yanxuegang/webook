// Code generated by MockGen. DO NOT EDIT.
// Source: ./type.go
//
// Generated by this command:
//
//	mockgen -source=./type.go -destination=./mocks/handler.mock.go -package=hdlmocks -typed=true Handler
//
// Package hdlmocks is a generated GoMock package.
package hdlmocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/ecodeclub/webook/internal/ai/internal/domain"
	handler "github.com/ecodeclub/webook/internal/ai/internal/service/llm/handler"
	gomock "go.uber.org/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// Handle mocks base method.
func (m *MockHandler) Handle(ctx context.Context, req domain.LLMRequest) (domain.LLMResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handle", ctx, req)
	ret0, _ := ret[0].(domain.LLMResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Handle indicates an expected call of Handle.
func (mr *MockHandlerMockRecorder) Handle(ctx, req any) *HandlerHandleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockHandler)(nil).Handle), ctx, req)
	return &HandlerHandleCall{Call: call}
}

// HandlerHandleCall wrap *gomock.Call
type HandlerHandleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *HandlerHandleCall) Return(arg0 domain.LLMResponse, arg1 error) *HandlerHandleCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *HandlerHandleCall) Do(f func(context.Context, domain.LLMRequest) (domain.LLMResponse, error)) *HandlerHandleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *HandlerHandleCall) DoAndReturn(f func(context.Context, domain.LLMRequest) (domain.LLMResponse, error)) *HandlerHandleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockBuilder is a mock of Builder interface.
type MockBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockBuilderMockRecorder
}

// MockBuilderMockRecorder is the mock recorder for MockBuilder.
type MockBuilderMockRecorder struct {
	mock *MockBuilder
}

// NewMockBuilder creates a new mock instance.
func NewMockBuilder(ctrl *gomock.Controller) *MockBuilder {
	mock := &MockBuilder{ctrl: ctrl}
	mock.recorder = &MockBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuilder) EXPECT() *MockBuilderMockRecorder {
	return m.recorder
}

// Next mocks base method.
func (m *MockBuilder) Next(next handler.Handler) handler.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", next)
	ret0, _ := ret[0].(handler.Handler)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockBuilderMockRecorder) Next(next any) *BuilderNextCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockBuilder)(nil).Next), next)
	return &BuilderNextCall{Call: call}
}

// BuilderNextCall wrap *gomock.Call
type BuilderNextCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *BuilderNextCall) Return(arg0 handler.Handler) *BuilderNextCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *BuilderNextCall) Do(f func(handler.Handler) handler.Handler) *BuilderNextCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *BuilderNextCall) DoAndReturn(f func(handler.Handler) handler.Handler) *BuilderNextCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

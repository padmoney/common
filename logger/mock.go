package logger

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

type MockLoggerMockRecorder struct {
	mock *MockLogger
}

func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

func (ml *MockLogger) Critical(message string) error {
	ret := ml.ctrl.Call(ml, "Critical", message)
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (mr *MockLoggerMockRecorder) Critical(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Critical", reflect.TypeOf((*MockLogger)(nil).Critical), message)
}

func (ml *MockLogger) Debug(message string) error {
	ret := ml.ctrl.Call(ml, "Debug", message)
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (mr *MockLoggerMockRecorder) Debug(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), message)
}

func (ml *MockLogger) Error(message string) error {
	ret := ml.ctrl.Call(ml, "Error", message)
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (mr *MockLoggerMockRecorder) Error(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), message)
}

func (ml *MockLogger) Info(message string) error {
	ret := ml.ctrl.Call(ml, "Info", message)
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (mr *MockLoggerMockRecorder) Info(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), message)
}

func (ml *MockLogger) Log(title, message, channel string) error {
	ret := ml.ctrl.Call(ml, "Log", title, message, channel)
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (mr *MockLoggerMockRecorder) Log(title, message, channel interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockLogger)(nil).Log), title, message, channel)
}

func (ml *MockLogger) Warning(message string) error {
	ret := ml.ctrl.Call(ml, "Warning", message)
	if ret[0] != nil {
		return ret[0].(error)
	}
	return nil
}

func (mr *MockLoggerMockRecorder) Warning(message interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockLogger)(nil).Warning), message)
}

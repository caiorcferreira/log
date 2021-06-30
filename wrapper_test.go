package log

import (
	"context"
	"io"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WrapperSuite struct {
	suite.Suite
}

func TestWrapperSuite(t *testing.T) {
	suite.Run(t, new(WrapperSuite))
}

func (s *WrapperSuite) TestWrapperWrappedMethods() {

	tt := []struct {
		name   string
		mock   func(*LoggerMock)
		method func()
	}{
		{
			name:   "Printf",
			method: func() { Printf("Blah") },
		},
		{
			name:   "Trace",
			method: func() { Trace("Blah") },
		},
		{
			name:   "Tracef",
			method: func() { Tracef("Blah") },
		},
		{
			name:   "Debug",
			method: func() { Debug("Blah") },
		},
		{
			name:   "Debugf",
			method: func() { Debugf("Blah") },
		},
		{
			name:   "Info",
			method: func() { Info("Blah") },
		},
		{
			name:   "Infof",
			method: func() { Infof("Blah") },
		},
		{
			name:   "Warn",
			method: func() { Warn("Blah") },
		},
		{
			name:   "Warnf",
			method: func() { Warnf("Blah") },
		},
		{
			name:   "Error",
			method: func() { Error("Blah") },
		},
		{
			name:   "Errorf",
			method: func() { Errorf("Blah") },
		},
		{
			name:   "Panic",
			method: func() { Panic("Blah") },
		},
		{
			name:   "Panicf",
			method: func() { Panicf("Blah") },
		},
		{
			name:   "Fatal",
			method: func() { Fatal("Blah") },
		},
		{
			name:   "Fatalf",
			method: func() { Fatalf("Blah") },
		},
		{
			name: "WithField",
			mock: func(l *LoggerMock) {
				l.On("WithField", mock.Anything, mock.Anything).Times(1).Return(l)
			},
			method: func() { WithField("key", "value") },
		},
		{
			name: "WithFields",
			mock: func(l *LoggerMock) {
				l.On("WithFields", mock.Anything).Times(1).Return(l)
			},
			method: func() {
				WithFields(Fields{
					"key":  "value",
					"key2": "value2",
				})
			},
		},
		{
			name: "WithTypeOf",
			mock: func(l *LoggerMock) {
				l.On("WithTypeOf", mock.Anything).Times(1).Return(l)
			},
			method: func() {
				WithTypeOf("obj")
			},
		},
	}

	for _, t := range tt {
		s.Run(t.name, func() {
			l := new(LoggerMock)
			if t.mock == nil {
				l.On(t.name, mock.Anything).Times(1)
			} else {
				t.mock(l)
			}
			NewLogger(l)
			t.method()
			l.AssertExpectations(s.T())
		})
	}
}

func (s *WrapperSuite) TestWrapperGetLogger() {

	tt := []struct {
		name string
		want Logger
	}{
		{
			name: "Success",
			want: new(LoggerMock),
		},
	}

	for _, t := range tt {
		s.Run(t.name, func() {
			NewLogger(t.want)
			got := GetLogger()
			s.Assert().True(reflect.DeepEqual(got, t.want), "got  %v\nwant %v", got, t.want)
		})
	}
}

//MOCK ------------------------------------------------------------
// LoggerMock is an autogenerated mock type for the LoggerMock type
type LoggerMock struct {
	mock.Mock
}

// Debug provides a mock function with given fields: args
func (_m *LoggerMock) Debug(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Debugf provides a mock function with given fields: format, args
func (_m *LoggerMock) Debugf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: args
func (_m *LoggerMock) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: format, args
func (_m *LoggerMock) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: args
func (_m *LoggerMock) Fatal(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: format, args
func (_m *LoggerMock) Fatalf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fields provides a mock function with given fields:
func (_m *LoggerMock) Fields() Fields {
	ret := _m.Called()

	var r0 Fields
	if rf, ok := ret.Get(0).(func() Fields); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Fields)
		}
	}

	return r0
}

// FromContext provides a mock function with given fields: ctx
func (_m *LoggerMock) FromContext(ctx context.Context) Logger {
	ret := _m.Called(ctx)

	var r0 Logger
	if rf, ok := ret.Get(0).(func(context.Context) Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

// Info provides a mock function with given fields: args
func (_m *LoggerMock) Info(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Infof provides a mock function with given fields: format, args
func (_m *LoggerMock) Infof(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Output provides a mock function with given fields:
func (_m *LoggerMock) Output() io.Writer {
	ret := _m.Called()

	var r0 io.Writer
	if rf, ok := ret.Get(0).(func() io.Writer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Writer)
		}
	}

	return r0
}

// Panic provides a mock function with given fields: args
func (_m *LoggerMock) Panic(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Panicf provides a mock function with given fields: format, args
func (_m *LoggerMock) Panicf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Printf provides a mock function with given fields: format, args
func (_m *LoggerMock) Printf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// ToContext provides a mock function with given fields: ctx
func (_m *LoggerMock) ToContext(ctx context.Context) context.Context {
	ret := _m.Called(ctx)

	var r0 context.Context
	if rf, ok := ret.Get(0).(func(context.Context) context.Context); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Trace provides a mock function with given fields: args
func (_m *LoggerMock) Trace(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Tracef provides a mock function with given fields: format, args
func (_m *LoggerMock) Tracef(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: args
func (_m *LoggerMock) Warn(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warnf provides a mock function with given fields: format, args
func (_m *LoggerMock) Warnf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// WithError provides a mock function with given fields: err
func (_m *LoggerMock) WithError(err error) Logger {
	ret := _m.Called(err)

	var r0 Logger
	if rf, ok := ret.Get(0).(func(error) Logger); ok {
		r0 = rf(err)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

// WithField provides a mock function with given fields: key, value
func (_m *LoggerMock) WithField(key string, value interface{}) Logger {
	ret := _m.Called(key, value)

	var r0 Logger
	if rf, ok := ret.Get(0).(func(string, interface{}) Logger); ok {
		r0 = rf(key, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

// WithFields provides a mock function with given fields: keyValues
func (_m *LoggerMock) WithFields(keyValues Fields) Logger {
	ret := _m.Called(keyValues)

	var r0 Logger
	if rf, ok := ret.Get(0).(func(Fields) Logger); ok {
		r0 = rf(keyValues)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

// WithTypeOf provides a mock function with given fields: obj
func (_m *LoggerMock) WithTypeOf(obj interface{}) Logger {
	ret := _m.Called(obj)

	var r0 Logger
	if rf, ok := ret.Get(0).(func(interface{}) Logger); ok {
		r0 = rf(obj)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Logger)
		}
	}

	return r0
}

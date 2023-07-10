// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	fixtureshttp "github.com/vektra/mockery/v2/pkg/fixtures/http"

	mock "gitlab.com/incubus8/gotest/mock"
)

// Example is an autogenerated mock type for the Example type
type Example struct {
	mock.Mock
}

type Example_Expecter struct {
	mock *mock.Mock
}

func (_m *Example) EXPECT() *Example_Expecter {
	return &Example_Expecter{mock: &_m.Mock}
}

// A provides a mock function with given fields:
func (_m *Example) A() http.Flusher {
	ret := _m.Called()

	var r0 http.Flusher
	if rf, ok := ret.Get(0).(func() http.Flusher); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Flusher)
		}
	}

	return r0
}

// Example_A_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'A'
type Example_A_Call struct {
	*mock.Call
}

// A is a helper method to define mock.On call
func (_e *Example_Expecter) A() *Example_A_Call {
	return &Example_A_Call{Call: _e.mock.On("A")}
}

func (_c *Example_A_Call) Run(run func()) *Example_A_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Example_A_Call) Return(_a0 http.Flusher) *Example_A_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Example_A_Call) RunAndReturn(run func() http.Flusher) *Example_A_Call {
	_c.Call.Return(run)
	return _c
}

// B provides a mock function with given fields: _a0
func (_m *Example) B(_a0 string) fixtureshttp.MyStruct {
	ret := _m.Called(_a0)

	var r0 fixtureshttp.MyStruct
	if rf, ok := ret.Get(0).(func(string) fixtureshttp.MyStruct); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(fixtureshttp.MyStruct)
	}

	return r0
}

// Example_B_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'B'
type Example_B_Call struct {
	*mock.Call
}

// B is a helper method to define mock.On call
//   - _a0 string
func (_e *Example_Expecter) B(_a0 interface{}) *Example_B_Call {
	return &Example_B_Call{Call: _e.mock.On("B", _a0)}
}

func (_c *Example_B_Call) Run(run func(_a0 string)) *Example_B_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Example_B_Call) Return(_a0 fixtureshttp.MyStruct) *Example_B_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Example_B_Call) RunAndReturn(run func(string) fixtureshttp.MyStruct) *Example_B_Call {
	_c.Call.Return(run)
	return _c
}

// NewExample creates a new instance of Example. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExample(t interface {
	mock.TestingT
	Cleanup(func())
}) *Example {
	mock := &Example{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

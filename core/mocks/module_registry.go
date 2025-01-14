// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	module "github.com/goto/entropy/core/module"
	mock "github.com/stretchr/testify/mock"
)

// ModuleRegistry is an autogenerated mock type for the Registry type
type ModuleRegistry struct {
	mock.Mock
}

type ModuleRegistry_Expecter struct {
	mock *mock.Mock
}

func (_m *ModuleRegistry) EXPECT() *ModuleRegistry_Expecter {
	return &ModuleRegistry_Expecter{mock: &_m.Mock}
}

// GetDriver provides a mock function with given fields: ctx, mod
func (_m *ModuleRegistry) GetDriver(ctx context.Context, mod module.Module) (module.Driver, module.Descriptor, error) {
	ret := _m.Called(ctx, mod)

	var r0 module.Driver
	var r1 module.Descriptor
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, module.Module) (module.Driver, module.Descriptor, error)); ok {
		return rf(ctx, mod)
	}
	if rf, ok := ret.Get(0).(func(context.Context, module.Module) module.Driver); ok {
		r0 = rf(ctx, mod)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(module.Driver)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, module.Module) module.Descriptor); ok {
		r1 = rf(ctx, mod)
	} else {
		r1 = ret.Get(1).(module.Descriptor)
	}

	if rf, ok := ret.Get(2).(func(context.Context, module.Module) error); ok {
		r2 = rf(ctx, mod)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ModuleRegistry_GetDriver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDriver'
type ModuleRegistry_GetDriver_Call struct {
	*mock.Call
}

// GetDriver is a helper method to define mock.On call
//   - ctx context.Context
//   - mod module.Module
func (_e *ModuleRegistry_Expecter) GetDriver(ctx interface{}, mod interface{}) *ModuleRegistry_GetDriver_Call {
	return &ModuleRegistry_GetDriver_Call{Call: _e.mock.On("GetDriver", ctx, mod)}
}

func (_c *ModuleRegistry_GetDriver_Call) Run(run func(ctx context.Context, mod module.Module)) *ModuleRegistry_GetDriver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(module.Module))
	})
	return _c
}

func (_c *ModuleRegistry_GetDriver_Call) Return(_a0 module.Driver, _a1 module.Descriptor, _a2 error) *ModuleRegistry_GetDriver_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *ModuleRegistry_GetDriver_Call) RunAndReturn(run func(context.Context, module.Module) (module.Driver, module.Descriptor, error)) *ModuleRegistry_GetDriver_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewModuleRegistry interface {
	mock.TestingT
	Cleanup(func())
}

// NewModuleRegistry creates a new instance of ModuleRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewModuleRegistry(t mockConstructorTestingTNewModuleRegistry) *ModuleRegistry {
	mock := &ModuleRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

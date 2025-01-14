// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	json "encoding/json"

	mock "github.com/stretchr/testify/mock"

	module "github.com/goto/entropy/core/module"
)

// ModuleService is an autogenerated mock type for the ModuleService type
type ModuleService struct {
	mock.Mock
}

type ModuleService_Expecter struct {
	mock *mock.Mock
}

func (_m *ModuleService) EXPECT() *ModuleService_Expecter {
	return &ModuleService_Expecter{mock: &_m.Mock}
}

// CreateModule provides a mock function with given fields: ctx, mod
func (_m *ModuleService) CreateModule(ctx context.Context, mod module.Module) (*module.Module, error) {
	ret := _m.Called(ctx, mod)

	var r0 *module.Module
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, module.Module) (*module.Module, error)); ok {
		return rf(ctx, mod)
	}
	if rf, ok := ret.Get(0).(func(context.Context, module.Module) *module.Module); ok {
		r0 = rf(ctx, mod)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*module.Module)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, module.Module) error); ok {
		r1 = rf(ctx, mod)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleService_CreateModule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateModule'
type ModuleService_CreateModule_Call struct {
	*mock.Call
}

// CreateModule is a helper method to define mock.On call
//   - ctx context.Context
//   - mod module.Module
func (_e *ModuleService_Expecter) CreateModule(ctx interface{}, mod interface{}) *ModuleService_CreateModule_Call {
	return &ModuleService_CreateModule_Call{Call: _e.mock.On("CreateModule", ctx, mod)}
}

func (_c *ModuleService_CreateModule_Call) Run(run func(ctx context.Context, mod module.Module)) *ModuleService_CreateModule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(module.Module))
	})
	return _c
}

func (_c *ModuleService_CreateModule_Call) Return(_a0 *module.Module, _a1 error) *ModuleService_CreateModule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ModuleService_CreateModule_Call) RunAndReturn(run func(context.Context, module.Module) (*module.Module, error)) *ModuleService_CreateModule_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteModule provides a mock function with given fields: ctx, urn
func (_m *ModuleService) DeleteModule(ctx context.Context, urn string) error {
	ret := _m.Called(ctx, urn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, urn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModuleService_DeleteModule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteModule'
type ModuleService_DeleteModule_Call struct {
	*mock.Call
}

// DeleteModule is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
func (_e *ModuleService_Expecter) DeleteModule(ctx interface{}, urn interface{}) *ModuleService_DeleteModule_Call {
	return &ModuleService_DeleteModule_Call{Call: _e.mock.On("DeleteModule", ctx, urn)}
}

func (_c *ModuleService_DeleteModule_Call) Run(run func(ctx context.Context, urn string)) *ModuleService_DeleteModule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ModuleService_DeleteModule_Call) Return(_a0 error) *ModuleService_DeleteModule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ModuleService_DeleteModule_Call) RunAndReturn(run func(context.Context, string) error) *ModuleService_DeleteModule_Call {
	_c.Call.Return(run)
	return _c
}

// GetModule provides a mock function with given fields: ctx, urn
func (_m *ModuleService) GetModule(ctx context.Context, urn string) (*module.Module, error) {
	ret := _m.Called(ctx, urn)

	var r0 *module.Module
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*module.Module, error)); ok {
		return rf(ctx, urn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *module.Module); ok {
		r0 = rf(ctx, urn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*module.Module)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, urn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleService_GetModule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetModule'
type ModuleService_GetModule_Call struct {
	*mock.Call
}

// GetModule is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
func (_e *ModuleService_Expecter) GetModule(ctx interface{}, urn interface{}) *ModuleService_GetModule_Call {
	return &ModuleService_GetModule_Call{Call: _e.mock.On("GetModule", ctx, urn)}
}

func (_c *ModuleService_GetModule_Call) Run(run func(ctx context.Context, urn string)) *ModuleService_GetModule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ModuleService_GetModule_Call) Return(_a0 *module.Module, _a1 error) *ModuleService_GetModule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ModuleService_GetModule_Call) RunAndReturn(run func(context.Context, string) (*module.Module, error)) *ModuleService_GetModule_Call {
	_c.Call.Return(run)
	return _c
}

// ListModules provides a mock function with given fields: ctx, project
func (_m *ModuleService) ListModules(ctx context.Context, project string) ([]module.Module, error) {
	ret := _m.Called(ctx, project)

	var r0 []module.Module
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]module.Module, error)); ok {
		return rf(ctx, project)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []module.Module); ok {
		r0 = rf(ctx, project)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]module.Module)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleService_ListModules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListModules'
type ModuleService_ListModules_Call struct {
	*mock.Call
}

// ListModules is a helper method to define mock.On call
//   - ctx context.Context
//   - project string
func (_e *ModuleService_Expecter) ListModules(ctx interface{}, project interface{}) *ModuleService_ListModules_Call {
	return &ModuleService_ListModules_Call{Call: _e.mock.On("ListModules", ctx, project)}
}

func (_c *ModuleService_ListModules_Call) Run(run func(ctx context.Context, project string)) *ModuleService_ListModules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ModuleService_ListModules_Call) Return(_a0 []module.Module, _a1 error) *ModuleService_ListModules_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ModuleService_ListModules_Call) RunAndReturn(run func(context.Context, string) ([]module.Module, error)) *ModuleService_ListModules_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateModule provides a mock function with given fields: ctx, urn, newConfigs
func (_m *ModuleService) UpdateModule(ctx context.Context, urn string, newConfigs json.RawMessage) (*module.Module, error) {
	ret := _m.Called(ctx, urn, newConfigs)

	var r0 *module.Module
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, json.RawMessage) (*module.Module, error)); ok {
		return rf(ctx, urn, newConfigs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, json.RawMessage) *module.Module); ok {
		r0 = rf(ctx, urn, newConfigs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*module.Module)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, json.RawMessage) error); ok {
		r1 = rf(ctx, urn, newConfigs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleService_UpdateModule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateModule'
type ModuleService_UpdateModule_Call struct {
	*mock.Call
}

// UpdateModule is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
//   - newConfigs json.RawMessage
func (_e *ModuleService_Expecter) UpdateModule(ctx interface{}, urn interface{}, newConfigs interface{}) *ModuleService_UpdateModule_Call {
	return &ModuleService_UpdateModule_Call{Call: _e.mock.On("UpdateModule", ctx, urn, newConfigs)}
}

func (_c *ModuleService_UpdateModule_Call) Run(run func(ctx context.Context, urn string, newConfigs json.RawMessage)) *ModuleService_UpdateModule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(json.RawMessage))
	})
	return _c
}

func (_c *ModuleService_UpdateModule_Call) Return(_a0 *module.Module, _a1 error) *ModuleService_UpdateModule_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ModuleService_UpdateModule_Call) RunAndReturn(run func(context.Context, string, json.RawMessage) (*module.Module, error)) *ModuleService_UpdateModule_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewModuleService interface {
	mock.TestingT
	Cleanup(func())
}

// NewModuleService creates a new instance of ModuleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewModuleService(t mockConstructorTestingTNewModuleService) *ModuleService {
	mock := &ModuleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

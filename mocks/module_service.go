// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/odpf/entropy/domain"
	mock "github.com/stretchr/testify/mock"
)

// ModuleService is an autogenerated mock type for the ServiceInterface type
type ModuleService struct {
	mock.Mock
}

type ModuleService_Expecter struct {
	mock *mock.Mock
}

func (_m *ModuleService) EXPECT() *ModuleService_Expecter {
	return &ModuleService_Expecter{mock: &_m.Mock}
}

// Act provides a mock function with given fields: ctx, r, action, params
func (_m *ModuleService) Act(ctx context.Context, r *domain.Resource, action string, params map[string]interface{}) (map[string]interface{}, error) {
	ret := _m.Called(ctx, r, action, params)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Resource, string, map[string]interface{}) map[string]interface{}); ok {
		r0 = rf(ctx, r, action, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Resource, string, map[string]interface{}) error); ok {
		r1 = rf(ctx, r, action, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleService_Act_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Act'
type ModuleService_Act_Call struct {
	*mock.Call
}

// Act is a helper method to define mock.On call
//  - ctx context.Context
//  - r *domain.Resource
//  - action string
//  - params map[string]interface{}
func (_e *ModuleService_Expecter) Act(ctx interface{}, r interface{}, action interface{}, params interface{}) *ModuleService_Act_Call {
	return &ModuleService_Act_Call{Call: _e.mock.On("Act", ctx, r, action, params)}
}

func (_c *ModuleService_Act_Call) Run(run func(ctx context.Context, r *domain.Resource, action string, params map[string]interface{})) *ModuleService_Act_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Resource), args[2].(string), args[3].(map[string]interface{}))
	})
	return _c
}

func (_c *ModuleService_Act_Call) Return(_a0 map[string]interface{}, _a1 error) *ModuleService_Act_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Sync provides a mock function with given fields: ctx, r
func (_m *ModuleService) Sync(ctx context.Context, r *domain.Resource) (*domain.Resource, error) {
	ret := _m.Called(ctx, r)

	var r0 *domain.Resource
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Resource) *domain.Resource); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Resource) error); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModuleService_Sync_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Sync'
type ModuleService_Sync_Call struct {
	*mock.Call
}

// Sync is a helper method to define mock.On call
//  - ctx context.Context
//  - r *domain.Resource
func (_e *ModuleService_Expecter) Sync(ctx interface{}, r interface{}) *ModuleService_Sync_Call {
	return &ModuleService_Sync_Call{Call: _e.mock.On("Sync", ctx, r)}
}

func (_c *ModuleService_Sync_Call) Run(run func(ctx context.Context, r *domain.Resource)) *ModuleService_Sync_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Resource))
	})
	return _c
}

func (_c *ModuleService_Sync_Call) Return(_a0 *domain.Resource, _a1 error) *ModuleService_Sync_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Validate provides a mock function with given fields: ctx, res
func (_m *ModuleService) Validate(ctx context.Context, res *domain.Resource) error {
	ret := _m.Called(ctx, res)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Resource) error); ok {
		r0 = rf(ctx, res)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ModuleService_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type ModuleService_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//  - ctx context.Context
//  - res *domain.Resource
func (_e *ModuleService_Expecter) Validate(ctx interface{}, res interface{}) *ModuleService_Validate_Call {
	return &ModuleService_Validate_Call{Call: _e.mock.On("Validate", ctx, res)}
}

func (_c *ModuleService_Validate_Call) Run(run func(ctx context.Context, res *domain.Resource)) *ModuleService_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.Resource))
	})
	return _c
}

func (_c *ModuleService_Validate_Call) Return(_a0 error) *ModuleService_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

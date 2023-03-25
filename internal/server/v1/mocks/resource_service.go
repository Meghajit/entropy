// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"

	module "github.com/goto/entropy/core/module"
	mock "github.com/stretchr/testify/mock"

	resource "github.com/goto/entropy/core/resource"
)

// ResourceService is an autogenerated mock type for the ResourceService type
type ResourceService struct {
	mock.Mock
}

type ResourceService_Expecter struct {
	mock *mock.Mock
}

func (_m *ResourceService) EXPECT() *ResourceService_Expecter {
	return &ResourceService_Expecter{mock: &_m.Mock}
}

// ApplyAction provides a mock function with given fields: ctx, urn, action
func (_m *ResourceService) ApplyAction(ctx context.Context, urn string, action module.ActionRequest) (*resource.Resource, error) {
	ret := _m.Called(ctx, urn, action)

	var r0 *resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, module.ActionRequest) (*resource.Resource, error)); ok {
		return rf(ctx, urn, action)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, module.ActionRequest) *resource.Resource); ok {
		r0 = rf(ctx, urn, action)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, module.ActionRequest) error); ok {
		r1 = rf(ctx, urn, action)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_ApplyAction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyAction'
type ResourceService_ApplyAction_Call struct {
	*mock.Call
}

// ApplyAction is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
//   - action module.ActionRequest
func (_e *ResourceService_Expecter) ApplyAction(ctx interface{}, urn interface{}, action interface{}) *ResourceService_ApplyAction_Call {
	return &ResourceService_ApplyAction_Call{Call: _e.mock.On("ApplyAction", ctx, urn, action)}
}

func (_c *ResourceService_ApplyAction_Call) Run(run func(ctx context.Context, urn string, action module.ActionRequest)) *ResourceService_ApplyAction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(module.ActionRequest))
	})
	return _c
}

func (_c *ResourceService_ApplyAction_Call) Return(_a0 *resource.Resource, _a1 error) *ResourceService_ApplyAction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_ApplyAction_Call) RunAndReturn(run func(context.Context, string, module.ActionRequest) (*resource.Resource, error)) *ResourceService_ApplyAction_Call {
	_c.Call.Return(run)
	return _c
}

// CreateResource provides a mock function with given fields: ctx, res
func (_m *ResourceService) CreateResource(ctx context.Context, res resource.Resource) (*resource.Resource, error) {
	ret := _m.Called(ctx, res)

	var r0 *resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource) (*resource.Resource, error)); ok {
		return rf(ctx, res)
	}
	if rf, ok := ret.Get(0).(func(context.Context, resource.Resource) *resource.Resource); ok {
		r0 = rf(ctx, res)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, resource.Resource) error); ok {
		r1 = rf(ctx, res)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_CreateResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateResource'
type ResourceService_CreateResource_Call struct {
	*mock.Call
}

// CreateResource is a helper method to define mock.On call
//   - ctx context.Context
//   - res resource.Resource
func (_e *ResourceService_Expecter) CreateResource(ctx interface{}, res interface{}) *ResourceService_CreateResource_Call {
	return &ResourceService_CreateResource_Call{Call: _e.mock.On("CreateResource", ctx, res)}
}

func (_c *ResourceService_CreateResource_Call) Run(run func(ctx context.Context, res resource.Resource)) *ResourceService_CreateResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.Resource))
	})
	return _c
}

func (_c *ResourceService_CreateResource_Call) Return(_a0 *resource.Resource, _a1 error) *ResourceService_CreateResource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_CreateResource_Call) RunAndReturn(run func(context.Context, resource.Resource) (*resource.Resource, error)) *ResourceService_CreateResource_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteResource provides a mock function with given fields: ctx, urn
func (_m *ResourceService) DeleteResource(ctx context.Context, urn string) error {
	ret := _m.Called(ctx, urn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, urn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResourceService_DeleteResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteResource'
type ResourceService_DeleteResource_Call struct {
	*mock.Call
}

// DeleteResource is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
func (_e *ResourceService_Expecter) DeleteResource(ctx interface{}, urn interface{}) *ResourceService_DeleteResource_Call {
	return &ResourceService_DeleteResource_Call{Call: _e.mock.On("DeleteResource", ctx, urn)}
}

func (_c *ResourceService_DeleteResource_Call) Run(run func(ctx context.Context, urn string)) *ResourceService_DeleteResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ResourceService_DeleteResource_Call) Return(_a0 error) *ResourceService_DeleteResource_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ResourceService_DeleteResource_Call) RunAndReturn(run func(context.Context, string) error) *ResourceService_DeleteResource_Call {
	_c.Call.Return(run)
	return _c
}

// GetLog provides a mock function with given fields: ctx, urn, filter
func (_m *ResourceService) GetLog(ctx context.Context, urn string, filter map[string]string) (<-chan module.LogChunk, error) {
	ret := _m.Called(ctx, urn, filter)

	var r0 <-chan module.LogChunk
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) (<-chan module.LogChunk, error)); ok {
		return rf(ctx, urn, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) <-chan module.LogChunk); ok {
		r0 = rf(ctx, urn, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan module.LogChunk)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, urn, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_GetLog_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLog'
type ResourceService_GetLog_Call struct {
	*mock.Call
}

// GetLog is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
//   - filter map[string]string
func (_e *ResourceService_Expecter) GetLog(ctx interface{}, urn interface{}, filter interface{}) *ResourceService_GetLog_Call {
	return &ResourceService_GetLog_Call{Call: _e.mock.On("GetLog", ctx, urn, filter)}
}

func (_c *ResourceService_GetLog_Call) Run(run func(ctx context.Context, urn string, filter map[string]string)) *ResourceService_GetLog_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(map[string]string))
	})
	return _c
}

func (_c *ResourceService_GetLog_Call) Return(_a0 <-chan module.LogChunk, _a1 error) *ResourceService_GetLog_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_GetLog_Call) RunAndReturn(run func(context.Context, string, map[string]string) (<-chan module.LogChunk, error)) *ResourceService_GetLog_Call {
	_c.Call.Return(run)
	return _c
}

// GetResource provides a mock function with given fields: ctx, urn
func (_m *ResourceService) GetResource(ctx context.Context, urn string) (*resource.Resource, error) {
	ret := _m.Called(ctx, urn)

	var r0 *resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*resource.Resource, error)); ok {
		return rf(ctx, urn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *resource.Resource); ok {
		r0 = rf(ctx, urn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, urn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_GetResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetResource'
type ResourceService_GetResource_Call struct {
	*mock.Call
}

// GetResource is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
func (_e *ResourceService_Expecter) GetResource(ctx interface{}, urn interface{}) *ResourceService_GetResource_Call {
	return &ResourceService_GetResource_Call{Call: _e.mock.On("GetResource", ctx, urn)}
}

func (_c *ResourceService_GetResource_Call) Run(run func(ctx context.Context, urn string)) *ResourceService_GetResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *ResourceService_GetResource_Call) Return(_a0 *resource.Resource, _a1 error) *ResourceService_GetResource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_GetResource_Call) RunAndReturn(run func(context.Context, string) (*resource.Resource, error)) *ResourceService_GetResource_Call {
	_c.Call.Return(run)
	return _c
}

// GetRevisions provides a mock function with given fields: ctx, selector
func (_m *ResourceService) GetRevisions(ctx context.Context, selector resource.RevisionsSelector) ([]resource.Revision, error) {
	ret := _m.Called(ctx, selector)

	var r0 []resource.Revision
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.RevisionsSelector) ([]resource.Revision, error)); ok {
		return rf(ctx, selector)
	}
	if rf, ok := ret.Get(0).(func(context.Context, resource.RevisionsSelector) []resource.Revision); ok {
		r0 = rf(ctx, selector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]resource.Revision)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, resource.RevisionsSelector) error); ok {
		r1 = rf(ctx, selector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_GetRevisions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRevisions'
type ResourceService_GetRevisions_Call struct {
	*mock.Call
}

// GetRevisions is a helper method to define mock.On call
//   - ctx context.Context
//   - selector resource.RevisionsSelector
func (_e *ResourceService_Expecter) GetRevisions(ctx interface{}, selector interface{}) *ResourceService_GetRevisions_Call {
	return &ResourceService_GetRevisions_Call{Call: _e.mock.On("GetRevisions", ctx, selector)}
}

func (_c *ResourceService_GetRevisions_Call) Run(run func(ctx context.Context, selector resource.RevisionsSelector)) *ResourceService_GetRevisions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.RevisionsSelector))
	})
	return _c
}

func (_c *ResourceService_GetRevisions_Call) Return(_a0 []resource.Revision, _a1 error) *ResourceService_GetRevisions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_GetRevisions_Call) RunAndReturn(run func(context.Context, resource.RevisionsSelector) ([]resource.Revision, error)) *ResourceService_GetRevisions_Call {
	_c.Call.Return(run)
	return _c
}

// ListResources provides a mock function with given fields: ctx, filter
func (_m *ResourceService) ListResources(ctx context.Context, filter resource.Filter) ([]resource.Resource, error) {
	ret := _m.Called(ctx, filter)

	var r0 []resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, resource.Filter) ([]resource.Resource, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, resource.Filter) []resource.Resource); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]resource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, resource.Filter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_ListResources_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListResources'
type ResourceService_ListResources_Call struct {
	*mock.Call
}

// ListResources is a helper method to define mock.On call
//   - ctx context.Context
//   - filter resource.Filter
func (_e *ResourceService_Expecter) ListResources(ctx interface{}, filter interface{}) *ResourceService_ListResources_Call {
	return &ResourceService_ListResources_Call{Call: _e.mock.On("ListResources", ctx, filter)}
}

func (_c *ResourceService_ListResources_Call) Run(run func(ctx context.Context, filter resource.Filter)) *ResourceService_ListResources_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(resource.Filter))
	})
	return _c
}

func (_c *ResourceService_ListResources_Call) Return(_a0 []resource.Resource, _a1 error) *ResourceService_ListResources_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_ListResources_Call) RunAndReturn(run func(context.Context, resource.Filter) ([]resource.Resource, error)) *ResourceService_ListResources_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateResource provides a mock function with given fields: ctx, urn, req
func (_m *ResourceService) UpdateResource(ctx context.Context, urn string, req resource.UpdateRequest) (*resource.Resource, error) {
	ret := _m.Called(ctx, urn, req)

	var r0 *resource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, resource.UpdateRequest) (*resource.Resource, error)); ok {
		return rf(ctx, urn, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, resource.UpdateRequest) *resource.Resource); ok {
		r0 = rf(ctx, urn, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, resource.UpdateRequest) error); ok {
		r1 = rf(ctx, urn, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceService_UpdateResource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateResource'
type ResourceService_UpdateResource_Call struct {
	*mock.Call
}

// UpdateResource is a helper method to define mock.On call
//   - ctx context.Context
//   - urn string
//   - req resource.UpdateRequest
func (_e *ResourceService_Expecter) UpdateResource(ctx interface{}, urn interface{}, req interface{}) *ResourceService_UpdateResource_Call {
	return &ResourceService_UpdateResource_Call{Call: _e.mock.On("UpdateResource", ctx, urn, req)}
}

func (_c *ResourceService_UpdateResource_Call) Run(run func(ctx context.Context, urn string, req resource.UpdateRequest)) *ResourceService_UpdateResource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(resource.UpdateRequest))
	})
	return _c
}

func (_c *ResourceService_UpdateResource_Call) Return(_a0 *resource.Resource, _a1 error) *ResourceService_UpdateResource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ResourceService_UpdateResource_Call) RunAndReturn(run func(context.Context, string, resource.UpdateRequest) (*resource.Resource, error)) *ResourceService_UpdateResource_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewResourceService interface {
	mock.TestingT
	Cleanup(func())
}

// NewResourceService creates a new instance of ResourceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewResourceService(t mockConstructorTestingTNewResourceService) *ResourceService {
	mock := &ResourceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	agents "go-drop-logistik/modules/agents"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Usecase) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, start, last
func (_m *Usecase) Fetch(ctx context.Context, start int, last int) ([]agents.Domain, int, error) {
	ret := _m.Called(ctx, start, last)

	var r0 []agents.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []agents.Domain); ok {
		r0 = rf(ctx, start, last)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]agents.Domain)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, int, int) int); ok {
		r1 = rf(ctx, start, last)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int, int) error); ok {
		r2 = rf(ctx, start, last)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) GetByID(ctx context.Context, id int) (agents.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 agents.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) agents.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(agents.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, email, password, sso
func (_m *Usecase) Login(ctx context.Context, email string, password string, sso bool) (string, string, error) {
	ret := _m.Called(ctx, email, password, sso)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) string); ok {
		r0 = rf(ctx, email, password, sso)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool) string); ok {
		r1 = rf(ctx, email, password, sso)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, bool) error); ok {
		r2 = rf(ctx, email, password, sso)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Register provides a mock function with given fields: ctx, data, sso
func (_m *Usecase) Register(ctx context.Context, data *agents.Domain, sso bool) error {
	ret := _m.Called(ctx, data, sso)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *agents.Domain, bool) error); ok {
		r0 = rf(ctx, data, sso)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, data, id
func (_m *Usecase) Update(ctx context.Context, data *agents.Domain, id int) error {
	ret := _m.Called(ctx, data, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *agents.Domain, int) error); ok {
		r0 = rf(ctx, data, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

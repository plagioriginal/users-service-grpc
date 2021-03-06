// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	protos "github.com/plagioriginal/users-service-grpc/users"
)

// UsersClient is an autogenerated mock type for the UsersClient type
type UsersClient struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: ctx, in, opts
func (_m *UsersClient) AddUser(ctx context.Context, in *protos.NewUserRequest, opts ...grpc.CallOption) (*protos.UserResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protos.NewUserRequest, ...grpc.CallOption) *protos.UserResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.NewUserRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, in, opts
func (_m *UsersClient) Login(ctx context.Context, in *protos.LoginRequest, opts ...grpc.CallOption) (*protos.TokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.TokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protos.LoginRequest, ...grpc.CallOption) *protos.TokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.TokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.LoginRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logout provides a mock function with given fields: ctx, in, opts
func (_m *UsersClient) Logout(ctx context.Context, in *protos.RefreshRequest, opts ...grpc.CallOption) (*protos.TokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.TokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protos.RefreshRequest, ...grpc.CallOption) *protos.TokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.TokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.RefreshRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: ctx, in, opts
func (_m *UsersClient) Refresh(ctx context.Context, in *protos.RefreshRequest, opts ...grpc.CallOption) (*protos.TokenResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *protos.TokenResponse
	if rf, ok := ret.Get(0).(func(context.Context, *protos.RefreshRequest, ...grpc.CallOption) *protos.TokenResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.TokenResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.RefreshRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

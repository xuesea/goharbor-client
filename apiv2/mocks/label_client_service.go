// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	label "github.com/xuesea/goharbor-client/v5/apiv2/internal/api/client/label"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// MockLabelClientService is an autogenerated mock type for the ClientService type
type MockLabelClientService struct {
	mock.Mock
}

// CreateLabel provides a mock function with given fields: params, authInfo
func (_m *MockLabelClientService) CreateLabel(params *label.CreateLabelParams, authInfo runtime.ClientAuthInfoWriter) (*label.CreateLabelCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *label.CreateLabelCreated
	if rf, ok := ret.Get(0).(func(*label.CreateLabelParams, runtime.ClientAuthInfoWriter) *label.CreateLabelCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*label.CreateLabelCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*label.CreateLabelParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLabel provides a mock function with given fields: params, authInfo
func (_m *MockLabelClientService) DeleteLabel(params *label.DeleteLabelParams, authInfo runtime.ClientAuthInfoWriter) (*label.DeleteLabelOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *label.DeleteLabelOK
	if rf, ok := ret.Get(0).(func(*label.DeleteLabelParams, runtime.ClientAuthInfoWriter) *label.DeleteLabelOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*label.DeleteLabelOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*label.DeleteLabelParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLabelByID provides a mock function with given fields: params, authInfo
func (_m *MockLabelClientService) GetLabelByID(params *label.GetLabelByIDParams, authInfo runtime.ClientAuthInfoWriter) (*label.GetLabelByIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *label.GetLabelByIDOK
	if rf, ok := ret.Get(0).(func(*label.GetLabelByIDParams, runtime.ClientAuthInfoWriter) *label.GetLabelByIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*label.GetLabelByIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*label.GetLabelByIDParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListLabels provides a mock function with given fields: params, authInfo
func (_m *MockLabelClientService) ListLabels(params *label.ListLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*label.ListLabelsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *label.ListLabelsOK
	if rf, ok := ret.Get(0).(func(*label.ListLabelsParams, runtime.ClientAuthInfoWriter) *label.ListLabelsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*label.ListLabelsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*label.ListLabelsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockLabelClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateLabel provides a mock function with given fields: params, authInfo
func (_m *MockLabelClientService) UpdateLabel(params *label.UpdateLabelParams, authInfo runtime.ClientAuthInfoWriter) (*label.UpdateLabelOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *label.UpdateLabelOK
	if rf, ok := ret.Get(0).(func(*label.UpdateLabelParams, runtime.ClientAuthInfoWriter) *label.UpdateLabelOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*label.UpdateLabelOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*label.UpdateLabelParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockLabelClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockLabelClientService creates a new instance of MockLabelClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockLabelClientService(t mockConstructorTestingTNewMockLabelClientService) *MockLabelClientService {
	mock := &MockLabelClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

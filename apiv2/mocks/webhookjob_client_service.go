// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	webhookjob "github.com/xuesea/goharbor-client/v5/apiv2/internal/api/client/webhookjob"
)

// MockWebhookjobClientService is an autogenerated mock type for the ClientService type
type MockWebhookjobClientService struct {
	mock.Mock
}

// ListWebhookJobs provides a mock function with given fields: params, authInfo
func (_m *MockWebhookjobClientService) ListWebhookJobs(params *webhookjob.ListWebhookJobsParams, authInfo runtime.ClientAuthInfoWriter) (*webhookjob.ListWebhookJobsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *webhookjob.ListWebhookJobsOK
	if rf, ok := ret.Get(0).(func(*webhookjob.ListWebhookJobsParams, runtime.ClientAuthInfoWriter) *webhookjob.ListWebhookJobsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*webhookjob.ListWebhookJobsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*webhookjob.ListWebhookJobsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockWebhookjobClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewMockWebhookjobClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWebhookjobClientService creates a new instance of MockWebhookjobClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockWebhookjobClientService(t mockConstructorTestingTNewMockWebhookjobClientService) *MockWebhookjobClientService {
	mock := &MockWebhookjobClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

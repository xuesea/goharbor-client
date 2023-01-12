// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	scan_all "github.com/xuesea/goharbor-client/v5/apiv2/internal/api/client/scan_all"
)

// MockScan_allClientService is an autogenerated mock type for the ClientService type
type MockScan_allClientService struct {
	mock.Mock
}

// CreateScanAllSchedule provides a mock function with given fields: params, authInfo
func (_m *MockScan_allClientService) CreateScanAllSchedule(params *scan_all.CreateScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*scan_all.CreateScanAllScheduleCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_all.CreateScanAllScheduleCreated
	if rf, ok := ret.Get(0).(func(*scan_all.CreateScanAllScheduleParams, runtime.ClientAuthInfoWriter) *scan_all.CreateScanAllScheduleCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_all.CreateScanAllScheduleCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_all.CreateScanAllScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestScanAllMetrics provides a mock function with given fields: params, authInfo
func (_m *MockScan_allClientService) GetLatestScanAllMetrics(params *scan_all.GetLatestScanAllMetricsParams, authInfo runtime.ClientAuthInfoWriter) (*scan_all.GetLatestScanAllMetricsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_all.GetLatestScanAllMetricsOK
	if rf, ok := ret.Get(0).(func(*scan_all.GetLatestScanAllMetricsParams, runtime.ClientAuthInfoWriter) *scan_all.GetLatestScanAllMetricsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_all.GetLatestScanAllMetricsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_all.GetLatestScanAllMetricsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLatestScheduledScanAllMetrics provides a mock function with given fields: params, authInfo
func (_m *MockScan_allClientService) GetLatestScheduledScanAllMetrics(params *scan_all.GetLatestScheduledScanAllMetricsParams, authInfo runtime.ClientAuthInfoWriter) (*scan_all.GetLatestScheduledScanAllMetricsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_all.GetLatestScheduledScanAllMetricsOK
	if rf, ok := ret.Get(0).(func(*scan_all.GetLatestScheduledScanAllMetricsParams, runtime.ClientAuthInfoWriter) *scan_all.GetLatestScheduledScanAllMetricsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_all.GetLatestScheduledScanAllMetricsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_all.GetLatestScheduledScanAllMetricsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScanAllSchedule provides a mock function with given fields: params, authInfo
func (_m *MockScan_allClientService) GetScanAllSchedule(params *scan_all.GetScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*scan_all.GetScanAllScheduleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_all.GetScanAllScheduleOK
	if rf, ok := ret.Get(0).(func(*scan_all.GetScanAllScheduleParams, runtime.ClientAuthInfoWriter) *scan_all.GetScanAllScheduleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_all.GetScanAllScheduleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_all.GetScanAllScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockScan_allClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// StopScanAll provides a mock function with given fields: params, authInfo
func (_m *MockScan_allClientService) StopScanAll(params *scan_all.StopScanAllParams, authInfo runtime.ClientAuthInfoWriter) (*scan_all.StopScanAllAccepted, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_all.StopScanAllAccepted
	if rf, ok := ret.Get(0).(func(*scan_all.StopScanAllParams, runtime.ClientAuthInfoWriter) *scan_all.StopScanAllAccepted); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_all.StopScanAllAccepted)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_all.StopScanAllParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateScanAllSchedule provides a mock function with given fields: params, authInfo
func (_m *MockScan_allClientService) UpdateScanAllSchedule(params *scan_all.UpdateScanAllScheduleParams, authInfo runtime.ClientAuthInfoWriter) (*scan_all.UpdateScanAllScheduleOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *scan_all.UpdateScanAllScheduleOK
	if rf, ok := ret.Get(0).(func(*scan_all.UpdateScanAllScheduleParams, runtime.ClientAuthInfoWriter) *scan_all.UpdateScanAllScheduleOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*scan_all.UpdateScanAllScheduleOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*scan_all.UpdateScanAllScheduleParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockScan_allClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockScan_allClientService creates a new instance of MockScan_allClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockScan_allClientService(t mockConstructorTestingTNewMockScan_allClientService) *MockScan_allClientService {
	mock := &MockScan_allClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

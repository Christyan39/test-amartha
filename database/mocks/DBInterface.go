// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "test/amartha/usecase/model"

// DBInterface is an autogenerated mock type for the DBInterface type
type DBInterface struct {
	mock.Mock
}

// CountVisitingURL provides a mock function with given fields: code
func (_m *DBInterface) CountVisitingURL(code string) error {
	ret := _m.Called(code)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(code)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateShortenCode provides a mock function with given fields: shorten
func (_m *DBInterface) CreateShortenCode(shorten *model.ShortlnRequest) error {
	ret := _m.Called(shorten)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ShortlnRequest) error); ok {
		r0 = rf(shorten)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetShortenByCode provides a mock function with given fields: code
func (_m *DBInterface) GetShortenByCode(code string) *model.ShortlnRequest {
	ret := _m.Called(code)

	var r0 *model.ShortlnRequest
	if rf, ok := ret.Get(0).(func(string) *model.ShortlnRequest); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ShortlnRequest)
		}
	}

	return r0
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import httprouter "github.com/julienschmidt/httprouter"
import mock "github.com/stretchr/testify/mock"

// HttpRouterIO is an autogenerated mock type for the HttpRouterIO type
type HttpRouterIO struct {
	mock.Mock
}

// GET provides a mock function with given fields: path, handle
func (_m *HttpRouterIO) GET(path string, handle httprouter.Handle) {
	_m.Called(path, handle)
}

// POST provides a mock function with given fields: path, handle
func (_m *HttpRouterIO) POST(path string, handle httprouter.Handle) {
	_m.Called(path, handle)
}
// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	terrakey "github.com/smartcontractkit/chainlink/core/services/keystore/keys/terrakey"
	mock "github.com/stretchr/testify/mock"
)

// Terra is an autogenerated mock type for the Terra type
type Terra struct {
	mock.Mock
}

// Add provides a mock function with given fields: key
func (_m *Terra) Add(key terrakey.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(terrakey.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields:
func (_m *Terra) Create() (terrakey.Key, error) {
	ret := _m.Called()

	var r0 terrakey.Key
	if rf, ok := ret.Get(0).(func() terrakey.Key); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(terrakey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Terra) Delete(id string) (terrakey.Key, error) {
	ret := _m.Called(id)

	var r0 terrakey.Key
	if rf, ok := ret.Get(0).(func(string) terrakey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(terrakey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EnsureKey provides a mock function with given fields:
func (_m *Terra) EnsureKey() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Export provides a mock function with given fields: id, password
func (_m *Terra) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *Terra) Get(id string) (terrakey.Key, error) {
	ret := _m.Called(id)

	var r0 terrakey.Key
	if rf, ok := ret.Get(0).(func(string) terrakey.Key); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(terrakey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *Terra) GetAll() ([]terrakey.Key, error) {
	ret := _m.Called()

	var r0 []terrakey.Key
	if rf, ok := ret.Get(0).(func() []terrakey.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]terrakey.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Import provides a mock function with given fields: keyJSON, password
func (_m *Terra) Import(keyJSON []byte, password string) (terrakey.Key, error) {
	ret := _m.Called(keyJSON, password)

	var r0 terrakey.Key
	if rf, ok := ret.Get(0).(func([]byte, string) terrakey.Key); ok {
		r0 = rf(keyJSON, password)
	} else {
		r0 = ret.Get(0).(terrakey.Key)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
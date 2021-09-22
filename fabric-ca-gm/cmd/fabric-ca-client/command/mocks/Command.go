// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import lib "github.com/VoneChain-CS/fabric-ca-gm/lib"
import mock "github.com/stretchr/testify/mock"
import viper "github.com/spf13/viper"

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

// ConfigInit provides a mock function with given fields:
func (_m *Command) ConfigInit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCfgFileName provides a mock function with given fields:
func (_m *Command) GetCfgFileName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetClientCfg provides a mock function with given fields:
func (_m *Command) GetClientCfg() *lib.ClientConfig {
	ret := _m.Called()

	var r0 *lib.ClientConfig
	if rf, ok := ret.Get(0).(func() *lib.ClientConfig); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lib.ClientConfig)
		}
	}

	return r0
}

// GetHomeDirectory provides a mock function with given fields:
func (_m *Command) GetHomeDirectory() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetViper provides a mock function with given fields:
func (_m *Command) GetViper() *viper.Viper {
	ret := _m.Called()

	var r0 *viper.Viper
	if rf, ok := ret.Get(0).(func() *viper.Viper); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*viper.Viper)
		}
	}

	return r0
}

// LoadMyIdentity provides a mock function with given fields:
func (_m *Command) LoadMyIdentity() (*lib.Identity, error) {
	ret := _m.Called()

	var r0 *lib.Identity
	if rf, ok := ret.Get(0).(func() *lib.Identity); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lib.Identity)
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

// SetDefaultLogLevel provides a mock function with given fields: _a0
func (_m *Command) SetDefaultLogLevel(_a0 string) {
	_m.Called(_a0)
}

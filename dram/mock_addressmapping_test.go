// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/mem/v2/dram/internal/addressmapping (interfaces: AddressConverter,Mapper)

package dram

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	addressmapping "gitlab.com/akita/mem/v2/dram/internal/addressmapping"
)

// MockAddressConverter is a mock of AddressConverter interface.
type MockAddressConverter struct {
	ctrl     *gomock.Controller
	recorder *MockAddressConverterMockRecorder
}

// MockAddressConverterMockRecorder is the mock recorder for MockAddressConverter.
type MockAddressConverterMockRecorder struct {
	mock *MockAddressConverter
}

// NewMockAddressConverter creates a new mock instance.
func NewMockAddressConverter(ctrl *gomock.Controller) *MockAddressConverter {
	mock := &MockAddressConverter{ctrl: ctrl}
	mock.recorder = &MockAddressConverterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddressConverter) EXPECT() *MockAddressConverterMockRecorder {
	return m.recorder
}

// ConvertExternalToInternal mocks base method.
func (m *MockAddressConverter) ConvertExternalToInternal(arg0 uint64) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertExternalToInternal", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ConvertExternalToInternal indicates an expected call of ConvertExternalToInternal.
func (mr *MockAddressConverterMockRecorder) ConvertExternalToInternal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertExternalToInternal", reflect.TypeOf((*MockAddressConverter)(nil).ConvertExternalToInternal), arg0)
}

// ConvertInternalToExternal mocks base method.
func (m *MockAddressConverter) ConvertInternalToExternal(arg0 uint64) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertInternalToExternal", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// ConvertInternalToExternal indicates an expected call of ConvertInternalToExternal.
func (mr *MockAddressConverterMockRecorder) ConvertInternalToExternal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertInternalToExternal", reflect.TypeOf((*MockAddressConverter)(nil).ConvertInternalToExternal), arg0)
}

// MockMapper is a mock of Mapper interface.
type MockMapper struct {
	ctrl     *gomock.Controller
	recorder *MockMapperMockRecorder
}

// MockMapperMockRecorder is the mock recorder for MockMapper.
type MockMapperMockRecorder struct {
	mock *MockMapper
}

// NewMockMapper creates a new mock instance.
func NewMockMapper(ctrl *gomock.Controller) *MockMapper {
	mock := &MockMapper{ctrl: ctrl}
	mock.recorder = &MockMapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMapper) EXPECT() *MockMapperMockRecorder {
	return m.recorder
}

// Map mocks base method.
func (m *MockMapper) Map(arg0 uint64) addressmapping.Location {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Map", arg0)
	ret0, _ := ret[0].(addressmapping.Location)
	return ret0
}

// Map indicates an expected call of Map.
func (mr *MockMapperMockRecorder) Map(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Map", reflect.TypeOf((*MockMapper)(nil).Map), arg0)
}

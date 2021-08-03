// Code generated by mockery 2.9.0. DO NOT EDIT.

package mocks

import (
	domain "section9/domain"

	mock "github.com/stretchr/testify/mock"
)

// GarageRepository is an autogenerated mock type for the GarageRepository type
type GarageRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: gar
func (_m *GarageRepository) Create(gar *domain.Garage) (*domain.Garage, error) {
	ret := _m.Called(gar)

	var r0 *domain.Garage
	if rf, ok := ret.Get(0).(func(*domain.Garage) *domain.Garage); ok {
		r0 = rf(gar)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Garage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Garage) error); ok {
		r1 = rf(gar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: gar
func (_m *GarageRepository) Delete(gar *domain.Garage) (bool, error) {
	ret := _m.Called(gar)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.Garage) bool); ok {
		r0 = rf(gar)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Garage) error); ok {
		r1 = rf(gar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields:
func (_m *GarageRepository) FindAll() ([]domain.Garage, error) {
	ret := _m.Called()

	var r0 []domain.Garage
	if rf, ok := ret.Get(0).(func() []domain.Garage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Garage)
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

// FindByID provides a mock function with given fields: ID
func (_m *GarageRepository) FindByID(ID int) (*domain.Garage, error) {
	ret := _m.Called(ID)

	var r0 *domain.Garage
	if rf, ok := ret.Get(0).(func(int) *domain.Garage); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Garage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: gar
func (_m *GarageRepository) Update(gar *domain.Garage) (bool, error) {
	ret := _m.Called(gar)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*domain.Garage) bool); ok {
		r0 = rf(gar)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.Garage) error); ok {
		r1 = rf(gar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
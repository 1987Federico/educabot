// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import entity "github.com/fede/golang_api/internal/domain/entity"
import mock "github.com/stretchr/testify/mock"

// TripRepository is an autogenerated mock type for the TripRepository type
type TripRepository struct {
	mock.Mock
}

// AssignTripToDriver provides a mock function with given fields: trip
func (_m *TripRepository) AssignTripToDriver(trip entity.Trip) entity.Trip {
	ret := _m.Called(trip)

	var r0 entity.Trip
	if rf, ok := ret.Get(0).(func(entity.Trip) entity.Trip); ok {
		r0 = rf(trip)
	} else {
		r0 = ret.Get(0).(entity.Trip)
	}

	return r0
}

// CloseTrip provides a mock function with given fields: driverID
func (_m *TripRepository) CloseTrip(driverID uint64) error {
	ret := _m.Called(driverID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(driverID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsTripOpen provides a mock function with given fields: driverID, finished
func (_m *TripRepository) IsTripOpen(driverID uint64, finished bool) *entity.Trip {
	ret := _m.Called(driverID, finished)

	var r0 *entity.Trip
	if rf, ok := ret.Get(0).(func(uint64, bool) *entity.Trip); ok {
		r0 = rf(driverID, finished)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Trip)
		}
	}

	return r0
}

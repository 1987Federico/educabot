package service

import (
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/mocks"
	"github.com/fede/golang_api/internal/platform/storage/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var (
	DB *gorm.DB
)

func TestMain(m *testing.M) {
	DB = mocks.MockDB()
	os.Exit(m.Run())
}

func newTripService() *TripServices {
	r := repository.NewTripRepository(DB)
	s := NewTripService(r)
	return s
}

func TestTripToDriverWithAssignError(t *testing.T) {
	s := newTripService()
	assert.Panics(t, func() {
		s.IsTripOpen(1, false, "assign")
	})
}

func TestTripToDriverWithAssign(t *testing.T) {
	expected := &entity.Trip{
		ID:       1,
		DriverID: 1,
		Finished: false,
	}
	s := newTripService()
	r := s.IsTripOpen(1, false, "close")
	assert.IsType(t, expected, r)
	assert.Equal(t, expected.DriverID, r.DriverID)
	assert.Equal(t, expected.ID, r.ID)
}

func TestTripToDriverWithClose(t *testing.T) {
	expected := &entity.Trip{
		ID:       1,
		DriverID: 1,
		Finished: false,
	}
	s := newTripService()
	r := s.IsTripOpen(1, false, "close")
	assert.IsType(t, expected, r)
	assert.Equal(t, expected.DriverID, r.DriverID)
	assert.Equal(t, expected.ID, r.ID)
}

func TestAssignTripToDriverWithCloseError(t *testing.T) {
	s := newTripService()
	assert.Panics(t, func() {
		s.IsTripOpen(2, false, "close")
	})
}

func TestCloseOK(t *testing.T) {
	s := newTripService()
	assert.NotPanics(t, func() {
		s.CloseTrip(1)
	})
}

func TestAssignTripToDriver(t *testing.T) {
	expected := &entity.Trip{
		ID:        3,
		StartTime: time.Time{},
		EndTime:   time.Time{},
		DriverID:  1,
		Finished:  false,
	}
	s := newTripService()
	r := s.AssignTripToDriver(1)

	assert.IsType(t, expected, r)
	assert.Equal(t, expected.DriverID, r.DriverID)
	assert.Equal(t, expected.Finished, r.Finished)
	assert.Equal(t, expected.ID, r.ID)
}

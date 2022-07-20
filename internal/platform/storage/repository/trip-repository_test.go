package repository

import (
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var (
	DBE *gorm.DB
	DB  *gorm.DB
)

func TestMain(m *testing.M) {
	DB = mocks.MockDB()
	os.Exit(m.Run())
}

func TestFindTripOpenWithData(t *testing.T) {
	repo := NewTripRepository(DB)

	result := repo.IsTripOpen(1, false)
	assert.IsType(t, &entity.Trip{}, result)
	assert.NotNil(t, result)
}

func TestFindTripOpenWithOutData(t *testing.T) {
	repo := NewTripRepository(DB)
	result := repo.IsTripOpen(4, false)
	assert.IsType(t, &entity.Trip{}, result)
	assert.Nil(t, result)
}

func TestCloseTripOK(t *testing.T) {
	repo := NewTripRepository(DB)
	err := repo.CloseTrip(3)
	assert.Nil(t, err)
}

func TestAssignTripToDrive(t *testing.T) {
	trip := entity.Trip{
		StartTime: time.Now(),
		Finished:  false,
		DriverID:  3,
		Driver:    nil,
	}
	repo := NewTripRepository(DB)
	result := repo.AssignTripToDriver(trip)
	assert.NotNil(t, result)
	assert.IsType(t, trip, result)
}

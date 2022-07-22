package service

import (
	"context"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	"github.com/fede/golang_api/internal/platform/storage/repository"
	"time"
)

type TripService interface {
	AssignTripToDriver(driverID uint64, ctx context.Context) *entity.Trip
	IsTripOpen(driverID uint64, finished bool, toTrip string, ctx context.Context) *entity.Trip
	CloseTrip(driverID uint64, ctx context.Context)
}

//go:generate mockery --case=snake   --case=camel --outpkg=mocks --output=../mocks --name=TripService

type TripServices struct {
	tripRepository *repository.TripConnection
	//driverRepository *repository.DriverRepository
}

//NewTripService .....
func NewTripService(tripRepo *repository.TripConnection) *TripServices {
	return &TripServices{
		tripRepository: tripRepo,
	}
}

func (t *TripServices) AssignTripToDriver(driverID uint64, ctx context.Context) *entity.Trip {
	var tripEntity = entity.Trip{
		StartTime: time.Now(),
		Finished:  false,
		DriverID:  driverID,
	}
	resp := t.tripRepository.AssignTripToDriver(tripEntity, ctx)
	return &resp
}

func (t *TripServices) IsTripOpen(driverID uint64, finished bool, toTrip string, ctx context.Context) *entity.Trip {
	resp := t.tripRepository.IsTripOpen(driverID, finished, ctx)
	if resp != nil && toTrip == "assign" {
		panic(errorCustom.ConflictApiError("Failed to process request", "before opening a new trip you must close the previous one"))
	} else if resp == nil && toTrip == "close" {
		panic(errorCustom.ConflictApiError("Failed to process request", "Driver has no open trips to close"))
	}
	return resp
}

func (t *TripServices) CloseTrip(driverID uint64, ctx context.Context) {
	if err := t.tripRepository.CloseTrip(driverID, ctx); err != nil {
		panic(errorCustom.InternalServerApiError("Error updating resource", err.Error()))
	}
}

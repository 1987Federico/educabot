package repository

import (
	"context"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/storage"
	"gorm.io/gorm"
	"time"
)

type TripRepository interface {
	AssignTripToDriver(trip entity.Trip, ctx context.Context) entity.Trip
	IsTripOpen(driverID uint64, finished bool, ctx context.Context) *entity.Trip
	CloseTrip(driverID uint64, ctx context.Context) error
}

//go:generate mockery --case=snake   --case=camel --outpkg=mocks --output=../../mocks --name=TripRepository

type TripConnection struct {
	connection *gorm.DB
}

//NewTripRepository is creates a new instance of UserRepository
func NewTripRepository(db *gorm.DB) *TripConnection {
	return &TripConnection{
		connection: db,
	}
}

func (t *TripConnection) AssignTripToDriver(trip entity.Trip, ctx context.Context) entity.Trip {
	db := storage.FromContext(ctx)
	db.Save(&trip).Preload("Driver").Find(&trip)
	return trip
}

func (t *TripConnection) IsTripOpen(driverID uint64, finished bool, ctx context.Context) *entity.Trip {
	var trip entity.Trip
	db := storage.FromContext(ctx)

	err := db.Where("driver_id = ? AND finished = ?", driverID, finished).
		Preload("Driver").Take(&trip).Error
	if err != nil {
		return nil
	}
	return &trip
}

func (t *TripConnection) CloseTrip(driverID uint64, ctx context.Context) error {
	db := storage.FromContext(ctx)

	err := db.Model(&entity.Trip{}).
		Where("driver_id = ? AND finished = ?", driverID, false).
		Updates(entity.Trip{Finished: true, EndTime: time.Now()}).Error
	if err != nil {
		return err
	}
	return nil
}

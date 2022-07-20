package repository

import (
	"github.com/fede/golang_api/internal/domain/entity"
	"gorm.io/gorm"
	"time"
)

type TripRepository interface {
	AssignTripToDriver(trip entity.Trip) entity.Trip
	IsTripOpen(driverID uint64, finished bool) *entity.Trip
	CloseTrip(driverID uint64) error
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

func (db *TripConnection) AssignTripToDriver(trip entity.Trip) entity.Trip {
	db.connection.Save(&trip).Preload("Driver").Find(&trip)
	return trip
}

func (db *TripConnection) IsTripOpen(driverID uint64, finished bool) *entity.Trip {
	var trip entity.Trip

	err := db.connection.Where("driver_id = ? AND finished = ?", driverID, finished).
		Preload("Driver").Take(&trip).Error
	if err != nil {
		return nil
	}
	return &trip
}

func (db *TripConnection) CloseTrip(driverID uint64) error {
	err := db.connection.Model(&entity.Trip{}).
		Where("driver_id = ? AND finished = ?", driverID, false).
		Updates(entity.Trip{Finished: true, EndTime: time.Now()}).Error
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errors"
	"gorm.io/gorm"
)

//DriverRepository is a ....
type DriverRepository interface {
	AllDriver(offset, limit int) []entity.Driver
	FindDriverByID(bookID uint64) *entity.Driver
	FindDriverByFile(driverFile uint64) *entity.Driver
	DriversWithoutTripsProgress() *[]entity.Driver
}

type DriverConnection struct {
	connection *gorm.DB
}

//NewDriverRepository creates an instance driverRepository
func NewDriverRepository(dbConn *gorm.DB) *DriverConnection {
	return &DriverConnection{
		connection: dbConn,
	}
}

func (db *DriverConnection) FindDriverByID(DriverID uint64) *entity.Driver {
	var driver entity.Driver
	err := db.connection.Preload("User").Preload("Trip").Find(&driver, DriverID).Error
	if err != nil {
		return nil
	}
	return &driver
}

func (db *DriverConnection) AllDriver(offset, limit int) []entity.Driver {
	var driver []entity.Driver
	db.connection.Preload("User").
		Preload("Trip").Offset(offset).Limit(limit).Find(&driver)
	return driver
}

func (db *DriverConnection) FindDriverByFile(driverFile uint64) *entity.Driver {
	var driver entity.Driver
	err := db.connection.Where("driver_file = ?", driverFile).Take(&driver).Error
	if err != nil {
		return nil
	}
	return &driver
}

func (db *DriverConnection) DriversWithoutTripsProgress() *[]entity.Driver {
	var driver []entity.Driver
	if err := db.connection.Raw("SELECT distinct d.*" +
		"FROM trips t JOIN drivers d ON d.id = t.driver_id where t.finished = true AND " +
		"t.driver_id NOT IN (SELECT driver_id FROM trips aux WHERE aux.finished = false)").Scan(&driver).Error; err != nil {
		panic(errors.InternalServerApiError("error base date", err.Error()))
	}
	return &driver
}

package repository

import (
	"context"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	"github.com/fede/golang_api/internal/platform/storage"
	"gorm.io/gorm"
)

//DriverRepository is a ....
type DriverRepository interface {
	AllDriver(offset, limit int, ctx context.Context) []entity.Driver
	FindDriverByID(bookID uint64, ctx context.Context) *entity.Driver
	FindDriverByFile(driverFile uint64, ctx context.Context) *entity.Driver
	DriversWithoutTripsProgress(ctx context.Context) *[]entity.Driver
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

func (d *DriverConnection) FindDriverByID(DriverID uint64, ctx context.Context) *entity.Driver {
	var driver entity.Driver
	db := storage.FromContext(ctx)
	err := db.Preload("User").Preload("Trip").Find(&driver, DriverID).Error
	if err != nil {
		return nil
	}
	return &driver
}

func (d *DriverConnection) AllDriver(offset, limit int, ctx context.Context) []entity.Driver {
	var driver []entity.Driver
	db := storage.FromContext(ctx)
	err := db.Preload("User").
		Preload("Trip").Offset(offset).Limit(limit).Find(&driver).Error
	if err != nil {
		panic(err)
	}
	return driver
}

func (d *DriverConnection) FindDriverByFile(driverFile uint64, ctx context.Context) *entity.Driver {
	var driver entity.Driver
	db := storage.FromContext(ctx)
	err := db.Where("driver_file = ?", driverFile).Take(&driver).Error
	if err != nil {
		return nil
	}
	return &driver
}

func (d *DriverConnection) DriversWithoutTripsProgress(ctx context.Context) *[]entity.Driver {
	var driver []entity.Driver
	db := storage.FromContext(ctx)
	if err := db.Raw("SELECT distinct d.*" +
		"FROM trips t JOIN drivers d ON d.id = t.driver_id where t.finished = true AND " +
		"t.driver_id NOT IN (SELECT driver_id FROM trips aux WHERE aux.finished = false)").Scan(&driver).Error; err != nil {
		panic(errorCustom.InternalServerApiError("error base date", err.Error()))
	}
	return &driver
}

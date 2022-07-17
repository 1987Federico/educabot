package repository

import (
	"github.com/ydhnwb/golang_api/entity"
	"gorm.io/gorm"
)

//DriverRepository is a ....
type DriverRepository interface {
	InsertDriver(b entity.Driver) entity.Driver
	UpdateDriver(b entity.Driver) entity.Driver
	DeleteDriver(b entity.Driver)
	AllDriver() []entity.Driver
	FindDriverByID(bookID uint64) entity.Driver
	FindDriverByFile(driverFile uint64) *entity.Driver
}

type DriverConnection struct {
	connection *gorm.DB
}

//NewDriverRepository creates an instance DriverRepository
func NewDriverRepository(dbConn *gorm.DB) DriverRepository {
	return &DriverConnection{
		connection: dbConn,
	}
}

func (db *DriverConnection) InsertDriver(b entity.Driver) entity.Driver {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *DriverConnection) UpdateDriver(b entity.Driver) entity.Driver {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *DriverConnection) DeleteDriver(b entity.Driver) {
	db.connection.Delete(&b)
}

func (db *DriverConnection) FindDriverByID(DriverID uint64) entity.Driver {
	var book entity.Driver
	db.connection.Preload("User").Find(&book, DriverID)
	return book
}

func (db *DriverConnection) AllDriver() []entity.Driver {
	var driver []entity.Driver
	db.connection.Preload("User").Find(&driver)
	return driver
}

func (db *DriverConnection) FindDriverByFile(driverFile uint64) *entity.Driver {
	var driver entity.Driver
	db.connection.Where("driverFile = ?", driverFile).Take(&driver)
	return &driver
}

package service

import (
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errors"
	"github.com/fede/golang_api/internal/platform/storage/repository"
)

//DriverService is a ....
type DriverService interface {
	All(pagination dto.DriverSearch) []entity.Driver
	FindByID(bookID uint64) *entity.Driver
	DriverExist(driverFile uint64) *entity.Driver
	DriversWithoutTripsProgress() *[]entity.Driver
}

type DriverServices struct {
	driverRepository *repository.DriverConnection
}

//NewDriverService .....
func NewDriverService(driverRepo *repository.DriverConnection) *DriverServices {
	return &DriverServices{
		driverRepository: driverRepo,
	}
}

func (s *DriverServices) All(pagination dto.DriverSearch) []entity.Driver {
	return s.driverRepository.AllDriver(pagination.Offset, pagination.Limit)
}

func (s *DriverServices) FindByID(driverID uint64) *entity.Driver {
	resp := s.driverRepository.FindDriverByID(driverID)
	if resp == nil {
		panic(errors.NotFoundApiError("Data not found", "No data with given id"))
	}
	return resp
}

func (s *DriverServices) DriverExist(driverFile uint64) *entity.Driver {
	res := s.driverRepository.FindDriverByFile(driverFile)
	if res == nil {
		panic(errors.NotFoundApiError("The driver with the requested file does not exist", "not found"))
	}
	return res
}

func (s *DriverServices) DriversWithoutTripsProgress() *[]entity.Driver {
	trip := s.driverRepository.DriversWithoutTripsProgress()
	if len(*trip) == 0 {
		panic(errors.NotFoundApiError("not found driver", "No drivers available"))
	}
	return trip
}

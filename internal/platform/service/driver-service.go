package service

import (
	"context"
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	"github.com/fede/golang_api/internal/platform/storage/repository"
)

//DriverService is a ....
type DriverService interface {
	AllDriver(pagination dto.DriverSearch, ctx context.Context) []entity.Driver
	FindByID(bookID uint64, ctx context.Context) *entity.Driver
	DriverExist(driverFile uint64, ctx context.Context) *entity.Driver
	DriversWithoutTripsProgress(ctx context.Context) *[]entity.Driver
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

func (s *DriverServices) AllDriver(pagination dto.DriverSearch, ctx context.Context) []entity.Driver {
	return s.driverRepository.AllDriver(pagination.Offset, pagination.Limit, ctx)
}

func (s *DriverServices) FindByID(driverID uint64, ctx context.Context) *entity.Driver {
	resp := s.driverRepository.FindDriverByID(driverID, ctx)
	if resp == nil {
		panic(errorCustom.NotFoundApiError("Data not found", "No data with given id"))
	}
	return resp
}

func (s *DriverServices) DriverExist(driverFile uint64, ctx context.Context) *entity.Driver {
	res := s.driverRepository.FindDriverByFile(driverFile, ctx)
	if res == nil {
		panic(errorCustom.NotFoundApiError("The driver with the requested file does not exist", "not found"))
	}
	return res
}

func (s *DriverServices) DriversWithoutTripsProgress(ctx context.Context) *[]entity.Driver {
	trip := s.driverRepository.DriversWithoutTripsProgress(ctx)
	if len(*trip) == 0 {
		panic(errorCustom.NotFoundApiError("not found driver", "No drivers available"))
	}
	return trip
}

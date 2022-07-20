package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errors"
	"github.com/fede/golang_api/internal/platform/helper/response"
	service2 "github.com/fede/golang_api/internal/platform/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//DriverController is a ...
type DriverController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	DriversWithoutTripsProgress(context *gin.Context)
}

type DriverControllers struct {
	driverService *service2.DriverServices
}

//NewDriverController create a new instances of DriverController
func NewDriverController(driverServ *service2.DriverServices, jwtServ *service2.JwtServices) *DriverControllers {
	return &DriverControllers{
		driverService: driverServ,
	}
}

func (d *DriverControllers) All(ctx *gin.Context) {
	var (
		drivers    []entity.Driver
		pagination = dto.DriverSearch{
			Offset: 0,
			Limit:  10,
		}
	)
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" && claims["role"] != "driver" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}
	if err := ctx.BindQuery(&pagination); err != nil {
		panic(errors.BadRequestApiError("Failed to process request", err.Error()))
	}
	drivers = d.driverService.All(pagination)
	res := response.BuildResponse(true, "OK", drivers)
	ctx.JSON(http.StatusOK, res)
}

func (d *DriverControllers) FindByID(context *gin.Context) {
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		panic(errors.BadRequestApiError("No param id was found", err.Error()))
	}

	token, _ := context.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" && claims["role"] != "driver" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	driver := d.driverService.FindByID(uint64(id))
	res := response.BuildResponse(true, "OK", driver)
	context.JSON(http.StatusOK, res)
}

func (d *DriverControllers) DriversWithoutTripsProgress(ctx *gin.Context) {
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	driver := d.driverService.DriversWithoutTripsProgress()
	res := response.BuildResponse(true, "OK", driver)
	ctx.JSON(http.StatusOK, res)

}

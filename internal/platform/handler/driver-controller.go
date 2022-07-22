package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
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

// AllDriver godoc
// @Summary Drivers
// @Description returns drivers who do not have a trip in progress
// @Tags Drivers
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Driver
// @Failure 401 {object} errorCustom.ApiError
// @Failure 404 {object} errorCustom.ApiError
// @Router /api/driver [get]
func (d *DriverControllers) AllDriver(ctx *gin.Context) {
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
		panic(errorCustom.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}
	if err := ctx.BindQuery(&pagination); err != nil {
		panic(errorCustom.BadRequestApiError("Failed to process request", err.Error()))
	}
	drivers = d.driverService.AllDriver(pagination, ctx.Request.Context())
	res := response.BuildResponse(true, "OK", drivers)
	ctx.JSON(http.StatusOK, res)
}

// FindByID godoc
// @Summary returns a driver
// @Description returns a driver searched by his id
// @Tags Drivers
// @Accept  json
// @Produce  json
// @Param id query string true "driver id"
// @Failure 400 {object} errorCustom.ApiError
// @Failure 401 {object} errorCustom.ApiError
// @Failure 404 {object} errorCustom.ApiError
// @Router /api/driver/{id} [get]
func (d *DriverControllers) FindByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		panic(errorCustom.BadRequestApiError("No param id was found", err.Error()))
	}

	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" && claims["role"] != "driver" {
		panic(errorCustom.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	driver := d.driverService.FindByID(uint64(id), ctx.Request.Context())
	res := response.BuildResponse(true, "OK", driver)
	ctx.JSON(http.StatusOK, res)
}

// DriversWithoutTripsProgress godoc
// @Summary Drivers Without Trips Progress
// @Description returns drivers who do not have a trip in progress
// @Tags Drivers
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Driver
// @Failure 401 {object} errorCustom.ApiError
// @Failure 404 {object} errorCustom.ApiError
// @Router /api/trip/driver/without/progress [get]
func (d *DriverControllers) DriversWithoutTripsProgress(ctx *gin.Context) {
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errorCustom.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	driver := d.driverService.DriversWithoutTripsProgress(ctx.Request.Context())
	res := response.BuildResponse(true, "OK", driver)
	ctx.JSON(http.StatusOK, res)

}

package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	"github.com/fede/golang_api/internal/platform/helper/response"
	service2 "github.com/fede/golang_api/internal/platform/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TripController interface {
	AssignTripToDriver(context *gin.Context)
	CloseTripToDriver(context *gin.Context)
}

type TripControllers struct {
	tripService   *service2.TripServices
	driverService *service2.DriverServices
}

func NewTripController(tripService *service2.TripServices, driverService *service2.DriverServices) *TripControllers {
	return &TripControllers{
		tripService:   tripService,
		driverService: driverService,
	}
}

// AssignTripToDriver godoc
// @Summary assign a ride to a driver
// @Description assign a trip to a driver who does not have a trip in progress
// @Tags Trips
// @Accept  json
// @Produce json
// @Param driver body dto.Trip true "driver's file to open a trip"
// @Success 201 {object} *entity.Trip
// @Failure 400 {object} errorCustom.ApiError
// @Failure 401 {object} errorCustom.ApiError
// @Failure 404 {object} errorCustom.ApiError
// @Failure 409 {object} errorCustom.ApiError
// @Failure 500 {object} errorCustom.ApiError
// @Router /api/trip/assign/driver[post]
func (t TripControllers) AssignTripToDriver(ctx *gin.Context) {
	var trip dto.Trip
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errorCustom.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	errDTO := ctx.ShouldBind(&trip)
	if errDTO != nil {
		panic(errorCustom.BadRequestApiError("Failed to process request", errDTO.Error()))
	}

	drive := t.driverService.DriverExist(trip.DriverFile, ctx.Request.Context())
	tripEntity := t.tripService.IsTripOpen(drive.ID, false, "assign", ctx.Request.Context())
	tripEntity = t.tripService.AssignTripToDriver(drive.ID, ctx.Request.Context())

	res := response.BuildResponse(true, "OK", tripEntity)
	ctx.JSON(http.StatusOK, res)
}

// CloseTripToDriver godoc
// @Summary assign a ride to a driver
// @Description assign a trip to a driver who does not have a trip in progress
// @Tags Trips
// @Accept  json
// @Produce json
// @Param driver body dto.Trip true "driver's file to open a trip"
// @Success 201 {object} *entity.Trip
// @Failure 400 {object} errorCustom.ApiError
// @Failure 401 {object} errorCustom.ApiError
// @Failure 404 {object} errorCustom.ApiError
// @Failure 409 {object} errorCustom.ApiError
// @Failure 500 {object} errorCustom.ApiError
// @Router /api/trip/close/driver[post]
func (t TripControllers) CloseTripToDriver(ctx *gin.Context) {
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errorCustom.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		panic(errorCustom.BadRequestApiError("No param id was found", err.Error()))
	}
	drive := t.driverService.FindByID(uint64(id), ctx.Request.Context())

	tripEntity := t.tripService.IsTripOpen(drive.ID, false, "close", ctx.Request.Context())
	t.tripService.CloseTrip(drive.ID, ctx.Request.Context())
	res := response.BuildResponse(true, "OK", tripEntity)
	ctx.JSON(http.StatusNoContent, res)

}

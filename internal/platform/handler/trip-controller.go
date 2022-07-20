package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/platform/helper/errors"
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

func (t TripControllers) AssignTripToDriver(ctx *gin.Context) {
	var trip dto.Trip
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	errDTO := ctx.ShouldBind(&trip)
	if errDTO != nil {
		panic(errors.BadRequestApiError("Failed to process request", errDTO.Error()))
	}

	drive := t.driverService.DriverExist(trip.DriverFile)
	tripEntity := t.tripService.IsTripOpen(drive.ID, false, "assign")
	tripEntity = t.tripService.AssignTripToDriver(drive.ID)

	res := response.BuildResponse(true, "OK", tripEntity)
	ctx.JSON(http.StatusOK, res)
}

func (t TripControllers) CloseTripToDriver(ctx *gin.Context) {
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		panic(errors.BadRequestApiError("No param id was found", err.Error()))
	}
	drive := t.driverService.FindByID(uint64(id))

	tripEntity := t.tripService.IsTripOpen(drive.ID, false, "close")
	t.tripService.CloseTrip(drive.ID)
	res := response.BuildResponse(true, "OK", tripEntity)
	ctx.JSON(http.StatusNoContent, res)

}

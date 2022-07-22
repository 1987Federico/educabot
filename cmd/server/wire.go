package server

import (
	"github.com/fede/golang_api/internal/platform/handler"
	"github.com/fede/golang_api/internal/platform/middleware"
	"github.com/fede/golang_api/internal/platform/service"
	"github.com/fede/golang_api/internal/platform/storage"
	"github.com/fede/golang_api/internal/platform/storage/repository"
	"go.uber.org/zap"
)

type WebHandler struct {
	authHandler   *handler.AuthControllers
	tripHandler   *handler.TripControllers
	driverHandler *handler.DriverControllers
	userHandler   *handler.UserControllers
	error         *middleware.Error
	authJwt       *middleware.AuthJWT
	transMid      *middleware.DBTransaction
}

func newWire() *WebHandler {
	l := newLogger()
	db := storage.SetupDatabaseConnection()
	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	driverRepository := repository.NewDriverRepository(db)
	tripRepository := repository.NewTripRepository(db)

	//SERVICES
	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository)
	driverService := service.NewDriverService(driverRepository)
	authService := service.NewAuthService(userRepository, roleRepository, driverRepository)
	tripService := service.NewTripService(tripRepository)

	//HANDLERS
	authHandler := handler.NewAuthController(authService, jwtService)
	userHandler := handler.NewUserController(userService)
	driverHandler := handler.NewDriverController(driverService, jwtService)
	tripHandler := handler.NewTripController(tripService, driverService)

	errMid := middleware.NewError(l)
	jwtMid := middleware.NewAuthorizeJWT(jwtService)
	transactionMid := middleware.NewDBTransaction(db)

	return &WebHandler{
		authHandler:   authHandler,
		tripHandler:   tripHandler,
		driverHandler: driverHandler,
		userHandler:   userHandler,
		error:         errMid,
		authJwt:       jwtMid,
		transMid:      transactionMid,
	}
}

func newLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return sugarLogger
}

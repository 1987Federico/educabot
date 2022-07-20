package server

import (
	handler2 "github.com/fede/golang_api/internal/platform/handler"
	middleware2 "github.com/fede/golang_api/internal/platform/middleware"
	service2 "github.com/fede/golang_api/internal/platform/service"
	"github.com/fede/golang_api/internal/platform/storage"
	repository2 "github.com/fede/golang_api/internal/platform/storage/repository"
	"go.uber.org/zap"
)

type WebHandler struct {
	authHandler   *handler2.AuthControllers
	tripHandler   *handler2.TripControllers
	driverHandler *handler2.DriverControllers
	userHandler   *handler2.UserControllers
	error         *middleware2.Error
	authJwt       *middleware2.AuthJWT
	//logger          *middleware.Logger
	//dataDog         *middleware.Telemetry
}

func newWire() *WebHandler {
	l := newLogger()
	db := storage.SetupDatabaseConnection()
	userRepository := repository2.NewUserRepository(db)
	roleRepository := repository2.NewRoleRepository(db)
	driverRepository := repository2.NewDriverRepository(db)
	tripRepository := repository2.NewTripRepository(db)

	//SERVICES
	jwtService := service2.NewJWTService()
	userService := service2.NewUserService(userRepository)
	driverService := service2.NewDriverService(driverRepository)
	authService := service2.NewAuthService(userRepository, roleRepository, driverRepository)
	tripService := service2.NewTripService(tripRepository)

	//HANDLERS
	authHandler := handler2.NewAuthController(authService, jwtService)
	userHandler := handler2.NewUserController(userService)
	driverHandler := handler2.NewDriverController(driverService, jwtService)
	tripHandler := handler2.NewTripController(tripService, driverService)

	//logMid := middleware.NewLogger(l)
	errMid := middleware2.NewError(l)
	jwtMid := middleware2.NewAuthorizeJWT(jwtService)
	//dataDog := middleware.NewTelemetry()

	return &WebHandler{
		authHandler:   authHandler,
		tripHandler:   tripHandler,
		driverHandler: driverHandler,
		userHandler:   userHandler,
		error:         errMid,
		authJwt:       jwtMid,
		//logger:          logMid,
		//dataDog:         dataDog,
	}
}

func newLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	sugarLogger := logger.Sugar()
	return sugarLogger
}

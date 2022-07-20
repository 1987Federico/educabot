package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

const (
	Namespace = "/challenge/educabot"
)

func NewServer() *Server {
	engine := gin.Default()
	return &Server{engine: engine}
}

func (s *Server) Run() error {
	wire := newWire()
	s.registerRouter(wire)
	//binding.Validator = new(validations.DefaultValidator)
	return s.engine.Run(fmt.Sprintf(":%s", "8080"))
}

func (s *Server) registerRouter(c *WebHandler) {
	//s.engine.Use(c.logger.Handler)
	s.engine.Use(c.error.Handler)
	s.engine.Use(c.authJwt.AuthorizeJWT)
	//s.engine.Use(c.dataDog.Handler)

	mainGroup := s.engine.Group(Namespace)

	auth := mainGroup.Group("/api/auth")
	auth.POST("/login", c.authHandler.Login)

	driver := mainGroup.Group("api/driver")
	driver.GET("/", c.driverHandler.All)
	driver.GET("/:id", c.driverHandler.FindByID)
	driver.POST("/register/driver", c.authHandler.RegisterDrive)
	//userRoutes.POST("/register/admin", authController.RegisterAdmin)

	trip := mainGroup.Group("api/trip")
	trip.GET("/driver/without/progress", c.driverHandler.DriversWithoutTripsProgress)
	trip.POST("/assign/driver", c.tripHandler.AssignTripToDriver)
	trip.PUT("/close/driver", c.tripHandler.CloseTripToDriver)
}

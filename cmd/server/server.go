package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	return s.engine.Run(fmt.Sprintf(":%s", "8080"))
}

func (s *Server) registerRouter(c *WebHandler) {

	s.engine.Use(c.error.Handler)
	s.engine.Use(c.transMid.Handler)

	mainGroup := s.engine.Group(Namespace)

	auth := mainGroup.Group("/api/auth")
	auth.POST("/login", c.authHandler.Login)

	driver := mainGroup.Group("api/driver", c.authJwt.AuthorizeJWT)
	driver.GET("/", c.driverHandler.AllDriver)
	driver.GET("/:id", c.driverHandler.FindByID)
	driver.POST("/register/driver", c.authHandler.RegisterDrive)

	trip := mainGroup.Group("api/trip", c.authJwt.AuthorizeJWT)
	trip.GET("/driver/without/progress", c.driverHandler.DriversWithoutTripsProgress)
	trip.POST("/assign/driver", c.tripHandler.AssignTripToDriver)
	trip.PUT("/close/driver", c.tripHandler.CloseTripToDriver)

	mainGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}

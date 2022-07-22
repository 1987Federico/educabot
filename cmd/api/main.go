package main

import (
	"fmt"
	"github.com/fede/golang_api/cmd/server"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title           Educabot
// @version         1.0
// @description     This is a challenge to space guru.
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  federicomatias.celeste@gmail.com
// @host      localhost:8080
// @BasePath  /challenge/educabot
// @securityDefinitions.basic  BasicAuth

func main() {
	docs.SwaggerInfo.Title = "Swagger Educabot"
	docs.SwaggerInfo.Description = "This is a challenge to space guru."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/challenge/educabot"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	srv := server.NewServer()
	err := srv.Run()
	if err != nil {
		panic(fmt.Errorf("failed to start server %w", err))
	}
}

//go:generate swag init -g main.go -o ../../docs/specs/ --parseDependency true --parseInternal true

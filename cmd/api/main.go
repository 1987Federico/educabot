package main

import (
	"fmt"
	"github.com/fede/golang_api/cmd/server"
)

func main() {
	srv := server.NewServer()
	err := srv.Run()
	if err != nil {
		panic(fmt.Errorf("failed to start server %w", err))
	}
}

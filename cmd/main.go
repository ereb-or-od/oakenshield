package main

import (
	"github.com/ereb-or-od/kenobi/pkg/server"
	"github.com/ereb-or-od/oakenshield/pkg/api"
)

func main(){
	kenobiServer := server.New("id_service").
		WithDefaultLogger().
		UseHttp().
		WithLoggingMiddleware().
		WithRecoverMiddleware().
		WithRequestIDMiddleware().
		WithAllowAnyCORSMiddleware().
		WithGzipMiddleware().
		WithHealthCheckMiddleware("/ping", "pong!").
		WithController(api.NewIdController())
	kenobiServer.Start()
}

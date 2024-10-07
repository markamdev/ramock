package main

import (
	"log/slog"

	_ "github.com/markamdev/ramock/internal/logging"
	"github.com/markamdev/ramock/internal/server"
	"github.com/markamdev/ramock/internal/settings"
)

func main() {
	slog.Info("starting REST API mocking service")

	opts := settings.ReadConfiguration("RAMOCK")
	er := server.NewEndpointRegisterer()
	er.RegisterHealthCheck()

	// test code - to be removed
	er.RegisterEndpoint("POST /test1", 201)
	er.RegisterEndpoint("DELETE /test2", 200)
	// test code end

	er.StartServer(opts.ListenPort)

}

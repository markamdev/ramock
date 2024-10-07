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
	er.StartServer(opts.ListenPort)
}

package main

import (
	"log/slog"

	_ "github.com/markamdev/ramock/internal/logging"
	"github.com/markamdev/ramock/internal/server"
	"github.com/markamdev/ramock/internal/settings"
)

func main() {
	slog.Info("starting REST API mocking service")

	opts := settings.ReadConfigurationFromEnv()
	er := server.NewEndpointRegisterer()
	er.RegisterHealthCheck()

	err := er.ReadEndpointsFromFile(opts.EndpointsFile)
	if err != nil {
		slog.Error("failed to read endpoints from file", "error", err)
		return
	}

	er.StartServer(opts.ListenPort)

}

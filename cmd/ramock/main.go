package main

import (
	"fmt"
	"log/slog"

	_ "github.com/markamdev/ramock/internal/logging"
	"github.com/markamdev/ramock/internal/server"
	"github.com/markamdev/ramock/internal/settings"
)

func main() {
	slog.Info(fmt.Sprintf("starting REST API mocking service (version %s)", server.RamockVersion))

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

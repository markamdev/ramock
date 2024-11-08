package main

import (
	"log/slog"
	"net/http"

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
	er.RegisterEndpoint(server.EndpointConfig{
		Path: "POST /items",
		Response: server.ResponseData{
			Code:        http.StatusCreated,
			Body:        []byte("{\"status\":\"created\",\"id\":5}"),
			ContentType: "application/json",
		},
	})
	er.RegisterEndpoint(server.EndpointConfig{
		Path: "GET /items",
		Response: server.ResponseData{
			Code: http.StatusOK,
			Body: []byte("{\"items\" : [ \"1\", \"3\", \"4\" ]}"),
			ContentType: "application/json",
		},
	})
	er.RegisterEndpoint(server.EndpointConfig{
		Path: "DELETE /items/it1",
		Response: server.ResponseData{
			Code: http.StatusOK,
		},
	})
	er.RegisterEndpoint(server.EndpointConfig{
		Path: "DELETE /items/it2",
		Response: server.ResponseData{
			Code: http.StatusNotFound,
		},
	})
	// test code end

	er.StartServer(opts.ListenPort)

}

package server

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/markamdev/ramock/internal/logging"
)

var (
	loggerInit    sync.Once
	packageLogger *slog.Logger
)

type EnpointHandler interface {
	RegisterHealthCheck()
	RegisterEndpoint(EndpointConfig) error // endpoint with method (as for ServeMux.HandleFunc) and HTTP error code number
	StartServer(int) error
	ReadEndpointsFromFile(path string) error // read endpoints from file and register them
}

func NewEndpointRegisterer() EnpointHandler {
	// this function if an entry point for server usage so Logger can be created here
	loggerInit.Do(func() {
		packageLogger = logging.GetSubLogger("server")
	})

	return &endpointHandler{
		mux:       http.NewServeMux(),
		endpoints: map[string]EndpointConfig{},
	}
}

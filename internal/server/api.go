package server

import (
	"log/slog"
	"net/http"

	"github.com/markamdev/ramock/internal/logging"
)

var (
	packageLogger *slog.Logger
)

type EnpointHandler interface {
	RegisterHealthCheck()
	RegisterEndpoint(string, int) error // endpoint with method (as for ServeMux.HandleFunc) and HTTP error code number
	StartServer(int) error
}

func NewEndpointRegisterer() EnpointHandler {
	// this function if an entry point for server usage so Logger can be created here
	if packageLogger == nil {
		packageLogger = logging.GetSubLogger("server")
	}
	return &endpointHandler{
		mux:       http.NewServeMux(),
		endpoints: map[string]int{},
	}
}

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type endpointHandler struct {
	mux       *http.ServeMux
	endpoints map[string]bool // to be changed/extended later
}

func (er *endpointHandler) RegisterHealthCheck() {
	packageLogger.Info("registering /health endpoint")
	er.mux.HandleFunc("GET /health", healthHandler)
	er.endpoints["/health"] = true
}

func (er *endpointHandler) StartServer(port int) error {
	listenAddress := fmt.Sprintf(":%d", port)

	packageLogger.Info(fmt.Sprintf("starting server at %s", listenAddress))
	return http.ListenAndServe(listenAddress, er.mux)
}

func healthHandler(wr http.ResponseWriter, req *http.Request) {
	packageLogger.Info("/health called")
	buf, err := json.Marshal(struct {
		State string `json:"state"`
	}{
		State: "running",
	})
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		return
	}

	wr.Write(buf)
	wr.Header().Add("Content-Type", "application/json")
	//wr.WriteHeader(http.StatusOK)
}

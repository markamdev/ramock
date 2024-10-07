package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type endpointHandler struct {
	mux       *http.ServeMux
	endpoints map[string]int // to be changed/extended later
}

func (er *endpointHandler) RegisterHealthCheck() {
	packageLogger.Info("registering /health endpoint")
	er.mux.HandleFunc("GET /health", er.healthHandler)
	er.endpoints["/health"] = http.StatusOK
}

func (er *endpointHandler) RegisterEndpoint(endpoint string, code int) error {
	_, ok := er.endpoints[endpoint]
	if ok {
		return ErrAlreadyRegistered
	}

	er.endpoints[endpoint] = code
	er.mux.HandleFunc(endpoint, er.commonHandler)
	return nil
}

func (er *endpointHandler) StartServer(port int) error {
	listenAddress := fmt.Sprintf(":%d", port)

	packageLogger.Info(fmt.Sprintf("starting server at %s", listenAddress))
	return http.ListenAndServe(listenAddress, er.mux)
}

func (er *endpointHandler) healthHandler(wr http.ResponseWriter, req *http.Request) {
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
}

func (er *endpointHandler) commonHandler(wr http.ResponseWriter, req *http.Request) {
	path := req.RequestURI
	method := req.Method
	endpoint := fmt.Sprintf("%s %s", method, path)
	code, ok := er.endpoints[endpoint]
	if !ok {
		packageLogger.Warn("unregistered uri and/or method", "uri", endpoint)
		wr.WriteHeader(http.StatusNotFound)
		return
	}

	wr.WriteHeader(code)
}

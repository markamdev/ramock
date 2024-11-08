package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type endpointHandler struct {
	mux       *http.ServeMux
	endpoints map[string]EndpointConfig
}

func (er *endpointHandler) RegisterHealthCheck() {
	packageLogger.Info("registering /health endpoint")
	er.mux.HandleFunc("GET /health", er.commonHandler)
	er.endpoints["GET /health"] = EndpointConfig{
		Path: "GET /health",
		Response: ResponseData{
			Code:        http.StatusOK,
			Body:        []byte("{\"state\":\"running\"}"),
			ContentType: "application/json",
		},
	}
}

func (er *endpointHandler) RegisterEndpoint(ec EndpointConfig) error {
	_, ok := er.endpoints[ec.Path]
	if ok {
		packageLogger.Error("endpoint and method '%s' already registered", "endpoint", ec.Path)
		return ErrAlreadyRegistered
	}

	er.endpoints[ec.Path] = ec
	er.mux.HandleFunc(ec.Path, er.commonHandler)
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
	endpointConfig, ok := er.endpoints[endpoint]
	if !ok {
		packageLogger.Warn("unregistered method-path pair", "endpoint", endpoint)
		wr.WriteHeader(http.StatusNotFound)
		return
	}

	if len(endpointConfig.Response.Body) > 0 {
		wr.Write(endpointConfig.Response.Body)
	}
	if len(endpointConfig.Response.ContentType) > 0 {
		wr.Header().Set("Content-Type", endpointConfig.Response.ContentType)
	}
	wr.WriteHeader(endpointConfig.Response.Code)
}

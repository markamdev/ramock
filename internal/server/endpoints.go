package server

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type endpointHandler struct {
	mux       *http.ServeMux
	endpoints map[string]EndpointDescription
}

func (er *endpointHandler) RegisterHealthCheck() {
	packageLogger.Info("registering /health endpoint")
	er.mux.HandleFunc("GET /health", er.commonHandler)
	er.endpoints["GET /health"] = EndpointDescription{
		Path:        "/health",
		Method:      "GET",
		Code:        http.StatusOK,
		Response:    "{\"state\":\"running\"}",
		ContentType: "application/json",
	}
}

func (er *endpointHandler) RegisterEndpoint(ec EndpointDescription) error {
	_, ok := er.endpoints[ec.Path]
	if ok {
		packageLogger.Error("endpoint and method '%s' already registered", "endpoint", ec.Path)
		return ErrAlreadyRegistered
	}

	epName := fmt.Sprintf("%s %s", ec.Method, ec.Path)
	packageLogger.Info("registering endpoint", "endpoint", epName)
	er.endpoints[epName] = ec
	er.mux.HandleFunc(epName, er.commonHandler)
	return nil
}

func (er *endpointHandler) StartServer(port int) error {
	listenAddress := fmt.Sprintf(":%d", port)

	packageLogger.Info(fmt.Sprintf("starting server at %s", listenAddress))
	return http.ListenAndServe(listenAddress, er.mux)
}

func (er *endpointHandler) commonHandler(wr http.ResponseWriter, req *http.Request) {
	path := req.RequestURI
	method := req.Method
	epName := fmt.Sprintf("%s %s", method, path)
	endpointConfig, ok := er.endpoints[epName]
	if !ok {
		packageLogger.Warn("unregistered method-path pair", "endpoint", epName)
		wr.WriteHeader(http.StatusNotFound)
		return
	}

	wr.WriteHeader(endpointConfig.Code)
	if len(endpointConfig.Response) > 0 {
		wr.Write([]byte(endpointConfig.Response))
	}
	if len(endpointConfig.ContentType) > 0 {
		wr.Header().Set("Content-Type", endpointConfig.ContentType)
	}
}

func (er *endpointHandler) ReadEndpointsFromFile(path string) error {
	if path == "" {
		return ErrInvalidEndpointFile
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var config ConfigData
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal yaml: %w", err)
	}

	// TODO: check version compatibility

	for _, ec := range config.Endpoints {
		err = er.RegisterEndpoint(ec)
		if err != nil {
			return fmt.Errorf("failed to register endpoint: %w", err)
		}
	}

	return nil
}

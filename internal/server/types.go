package server

import "errors"

var (
	ErrAlreadyRegistered = errors.New("already registered")
)

// EndpointConfig basic decription of supported endpoint
// TODO expand with expected request body
type EndpointConfig struct {
	Path     string
	Response ResponseData
}

type ResponseData struct {
	Code        int
	Body        []byte
	ContentType string
}

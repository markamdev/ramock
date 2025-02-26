package server

import "errors"

var (
	ErrAlreadyRegistered   = errors.New("already registered")
	ErrInvalidEndpointFile = errors.New("invalid endpoint file")
)

// EndpointConfig basic decription of supported endpoint
// TODO expand with expected request body
type EndpointConfig struct {
	Path        string
	Method      string
	Code        int
	Response    string
	ContentType string
}

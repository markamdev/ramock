package server

import "errors"

var (
	ErrAlreadyRegistered   = errors.New("already registered")
	ErrInvalidEndpointFile = errors.New("invalid endpoint file")
)

// EndpointDescription basic decription of supported endpoint
// TODO expand with expected request body
type EndpointDescription struct {
	Path        string `yaml:"path"`
	Method      string `yaml:"method"`
	Code        int    `yaml:"code"`
	Response    string `yaml:"body"`
	ContentType string `yaml:"contentType"`
}

type ConfigData struct {
	RamockVersion string                `yaml:"ramockVersion"`
	Endpoints     []EndpointDescription `yaml:"endpoints"`
}
